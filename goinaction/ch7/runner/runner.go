package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

/*
runner 包用于展示如何使用通道来监视程序的执行时间，如果程序运行时间太长，也可以
用 runner 包来终止程序。当开发需要调度后台处理任务的程序的时候，这种模式会很有用。这
个程序可能会作为 cron 作业执行，或者在基于定时任务的云环境（如 iron.io）里执行。

* 程序可以在分配的时间内完成工作，正常终止
* 程序没有及时完成工作，“自杀”
* 接收到操作系统发送的中断事件，程序立即尝试清理状态并停止工作

*/

// Runner 在给定的超时时间内执行一组任务。并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt 通道报告从操作系统发送的信号
	interrupt chan os.Signal

	// complete 通道报告处理任务已经完成
	complete chan error

	// timeout 通道报告处理任务已经超时
	timeout <-chan time.Time

	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout 会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新的准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add 将一个任务附加到Runner上，这个任务是一个接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	// 我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务运行过程中时发出错误信号，立即返回
	case err := <-r.complete:
		return err
	// 当任务处理程序运行超时时发出信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run 执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已注册的任务
		task(id)
	}
	return nil
}

// goInterrupt 验证是否接收到中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被触发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
