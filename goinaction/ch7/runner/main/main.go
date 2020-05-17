package main

import (
	"go_learning/goinaction/ch7/runner"
	"log"
	"os"
	"time"
)

const timeout = 1 * time.Second

func main() {
	log.Println("Starting work.")
	// 指定超时时间
	r := runner.New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(2)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(3)
		}
		log.Println("Process ended.")
	}
}

// createTask 返回一个根据id 休眠指定秒数的示例任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor-Task # %d.", id+1)
		time.Sleep(time.Duration(id+1) * time.Second)
	}
}
