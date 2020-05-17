package main

import (
	"go_learning/goinaction/ch7/work/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
}

// namePrinter 使用特点的方式打印名字
type namePrinter struct {
	name string
}

// Task 实现Worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	// 使用两个goroutine 
	// 创建工作池
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	for i := 0; i < 10; i++ {
		// 每个名字都会创建10个goroutine来提交任务
		for _, name := range names {
			// 创建一个namePrinter并提供指定的名字
			np := namePrinter{
				name: name,
			}
			go func() {
				// 将任务提交执行，当Run返回时就可以知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()

	// 让工作池停止工作，等待所有工作完成
	p.Shutdown()
}
