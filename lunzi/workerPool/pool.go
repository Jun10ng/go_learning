package workerPool

import (
	"sync"
	"sync/atomic"
)

const (
	// 默认worker队列大小
	defaultQSize = 128
	// 输出队列大小
	outputChanSize = 100
)

// GoWorkers worker集合

// All workers will be killed after Stop() is called if their respective job finishes.
type GoWorkers struct {
	numWorkers uint32
	maxWorkers uint32
	numJobs    uint32
	workerQ    chan func()
	bufferedQ  chan func()
	jobQ       chan func()
	stopping   int32
	done       chan struct{}

	// ErrChan 收集job产生的err，在Stop() return后被关闭。
	// 仅对SubmitCheckError()和SubmitCheckResult()有效
	// 在submit job之前监听ErrChan，以便不会漏过任何更新
	ErrChan chan error

	ResultChan chan interface{}
}

// Workers worker数目，如果没指定或者0，则按需求生成worker
// QSize job数
// Minimum value is 128.
type Options struct {
	Workers uint32
	QSize   uint32
}

func New(args ...Options) *GoWorkers {
	gw := &GoWorkers{
		workerQ: make(chan func()),
		// Do not remove jobQ. To stop receiving input once Stop() is called
		jobQ:       make(chan func()),
		ErrChan:    make(chan error, outputChanSize),
		ResultChan: make(chan interface{}, outputChanSize),
		done:       make(chan struct{}),
	}

	gw.bufferedQ = make(chan func(), defaultQSize)
	if len(args) == 1 {
		gw.maxWorkers = args[0].Workers
		if args[0].QSize > defaultQSize {
			gw.bufferedQ = make(chan func(), args[0].QSize)
		}
	}

	go gw.start()

	return gw
}

// JobNum returns number of active jobs
func (gw *GoWorkers) JobNum() uint32 {
	return atomic.LoadUint32(&gw.numJobs)
}

// WorkerNum returns number of active workers
func (gw *GoWorkers) WorkerNum() uint32 {
	return atomic.LoadUint32(&gw.numWorkers)
}

// 无返回error的job
func (gw *GoWorkers) Submit(job func()) {
	if atomic.LoadInt32(&gw.stopping) == 1 {
		return
	}
	atomic.AddUint32(&gw.numJobs, uint32(1))
	gw.jobQ <- func() { job() }
}

// 有返回error的job
func (gw *GoWorkers) SubmitCheckError(job func() error) {
	if atomic.LoadInt32(&gw.stopping) == 1 {
		return
	}
	atomic.AddUint32(&gw.numJobs, uint32(1))
	gw.jobQ <- func() {
		err := job()
		if err != nil {
			select {
			case gw.ErrChan <- err:
			default:
			}
		}
	}
}

// 需要结果与错误
func (gw *GoWorkers) SubmitCheckResult(job func() (interface{}, error)) {
	if atomic.LoadInt32(&gw.stopping) == 1 {
		return
	}
	atomic.AddUint32(&gw.numJobs, uint32(1))
	gw.jobQ <- func() {
		result, err := job()
		if err != nil {
			select {
			case gw.ErrChan <- err:
			default:
			}
		} else {
			select {
			case gw.ResultChan <- result:
			default:
			}
		}
	}
}

// Stop 相当于 Wait()
// wait = false 不会等待errChan 或者 resultChan 消费干净
// 所以当你想接受所有的result或者err时 请使用true
func (gw *GoWorkers) Stop(wait bool) {
	if !atomic.CompareAndSwapInt32(&gw.stopping, 0, 1) {
		return
	}
	if gw.JobNum() != 0 {
		<-gw.done
	}

	if wait {
		for len(gw.ResultChan)|len(gw.ErrChan) == 0 {
			break
		}
	}

	// close the input channel
	close(gw.jobQ)
}

var mx sync.Mutex

func (gw *GoWorkers) spawnWorker() {
	defer mx.Unlock()
	mx.Lock()
	if ((gw.maxWorkers == 0) || (gw.WorkerNum() < gw.maxWorkers)) && (gw.JobNum() > gw.WorkerNum()) {
		go gw.startWorker()
	}
}

func (gw *GoWorkers) start() {
	defer func() {
		close(gw.bufferedQ)
		close(gw.workerQ)
		close(gw.ErrChan)
		close(gw.ResultChan)
	}()

	// start a worker in advance
	go gw.startWorker()

	go func() {
		for {
			select {
			// keep processing the queued jobs
			case job, ok := <-gw.bufferedQ:
				if !ok {
					return
				}
				go func() {
					gw.spawnWorker()
					gw.workerQ <- job
				}()
			}
		}
	}()

	for {
		select {
		case job, ok := <-gw.jobQ:
			if !ok {
				return
			}
			select {
			// if possible, process the job without queueing
			case gw.workerQ <- job:
				go gw.spawnWorker()
			// queue it if no workers are available
			default:
				gw.bufferedQ <- job
			}
		}
	}
}

func (gw *GoWorkers) startWorker() {
	defer func() {
		atomic.AddUint32(&gw.numWorkers, ^uint32(0))
	}()

	atomic.AddUint32(&gw.numWorkers, 1)

	for job := range gw.workerQ {
		job()
		if (atomic.AddUint32(&gw.numJobs, ^uint32(0)) == 0) && (atomic.LoadInt32(&gw.stopping) == 1) {
			gw.done <- struct{}{}
		}
	}
}
