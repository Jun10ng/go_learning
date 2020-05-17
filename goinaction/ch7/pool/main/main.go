package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/goinaction/code/chapter7/patterns/pool"
)

const (
	// 要使用的goroutine的数量
	maxGoroutines = 25
	// 池中资源数量
	pooledResources = 2
)

// dbConnection 模拟要共享的资源
type dbConnection struct {
	ID int32
}

// Close 实现了io.Closer接口,以便dbConnection可以被池管理.Close用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close:Connection", dbConn.ID)
	return nil
}

// idCounter 用来给每一个连接分配一个独一无二的id
var idCounter int32

// createConnection是一个工厂函数
// 当需要一个新连接的时候，资源池会调用这个函数
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create:New Connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建用来连接池
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用池里的连接来完成查询
	for query := 0; query < maxGoroutines; query++ {
		// 每个goroutine需要自己复制一份要查询值的副本，不然所有的查询会共享一个查询变量
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	// 等待goroutin结束
	wg.Wait()
	// 关闭池
	log.Println("shutdown program")
	p.Close()
}

// performQueries 用来测试连接的资源池
func performQueries(query int, p *pool.Pool) {
	// 从池里请求一个连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 将该链接释放回池里
	defer p.Release(conn)

	// 用等待来模拟查询相应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)

}
