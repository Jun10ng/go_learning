package pool

import (
	"log"
	"errors"
	"io"
	"sync"
)

/*
使用channel实现资源池，来管理可以在任意数量的gorotine之间共享及独立使用的资源。
这种模式在需要共享以组静态资源的情况下（如共享数据库连接或者内存缓存区）非常有用。

如果goroutine需要从这些资源里的其中一个，可以从池里申请，使用完后归还到资源池。

也就是标准库里的 sync.pool
*/

// Pool 管理一组可以安全在多个goroutine间共享的资源，被管理的资源必须实现io.Closer接口
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed   bool
}

// ErrPoolClosed 表示请求（Acquire) 了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed.")

// New创建一个用来管理资源的池，这个池需要一个可以分配新资源的函数，并规定池的大小
func New(fn func()(io.Closer,error),size uint)(* Pool,error){
	if size <= 0{
		return nil,errors.New("Size value too small.")
	}
	return &Pool{
		factory: fn,
		resources:make(chan io.Closer, size)
	},nil
}

// Acquire 从池中获取一个资源
func (p *Pool)Acquire() (io.Closer,error) {
	select{
	case r,ok := <-p.resources:
		log.Println("Acquire:","Shared Resource")
		if !ok{
			return nil,ErrPoolClosed
		}
		return r,nil
		// 没有默认空闲资源可用，所以提供一个新资源
	default:
		log.Println("Acquire:","New Resource")
		return p.factory()
	}
}

// Release 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和Close操作的安全
	// 释放前要先加锁
	p.m.Lock()
	defer p.m.Unlock()
	// 如果池已经关闭，销毁这个资源
	if p.closed{
		r.Close()
		return 
	}

	/* 
	这个实现是通过 select/case 语句来检
	查有缓冲的通道里是否还有资源来完成的。如果通道里还有资源，就取出这个资源，并返回给调用者。如果该通道里没有资源可取，就会执行 default 分支。在这个示例中，在第 54 行执行用户提供的工厂函数，并且创建并返回一个新资源
	*/
	select{
		// 试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:","In Queue")

		// 如果队列满了则关闭这个资源
	default:
		log.Println("Release:","Closing")
		r.Close()
	}
}

// Close 会让资源池停止工作，并关闭所有现有资源
func (p *Pool) Close() {
	// 保证本操作与Release操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	// 如果pool已经被关闭了，什么也不做
	if p.closed{
		return
	}

	// 关闭池
	p.closed = true

	// 在清空通道里的资源之前，将通道关闭
	// 如果不这么做，会发生死锁
	close(p.resources)

	// 关闭资源
	for r:= range p.resources{
		r.Close()
	}
}