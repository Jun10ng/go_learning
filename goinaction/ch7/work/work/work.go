package work

import (
	"sync"
)

/*
work包的目的是展示如何使用无缓冲的通道来创建一个goroutine池，这些goroutine执行并控制一组工作，让其并发执行。

这种无缓冲的通道的方法允许使用者知道什么时候goroutine池正在执行工作。
*/

type Worker interface {
	Task()
}

// Pool 提供一个goroutine池，这个池可以完成任何已提交的Worker任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个新工作池
func New(maxGoroutine int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutine)

	for i := 0; i < maxGoroutine; i++ {
		go func() {
			// 等任务提交
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有goroutine停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
