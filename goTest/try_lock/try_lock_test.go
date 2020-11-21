package try_lock

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

/*
	Go 并发编程实战课
	https://time.geekbang.org/column/article/296793

	TryLock
	TryLock 方法请求锁的时候，如果这把锁没有被其他 goroutine 所持有，
    那么，这个 goroutine 就持有了这把锁，并返回 true；如果这把锁已经被其他 goroutine 所持有，
    或者是正在准备交给某个被唤醒的 goroutine，那么，这个请求锁的 goroutine 就直接返回 false，不会阻塞在方法调用上。
*/

const (
	mutexLocked      = 1 << iota //1
	mutexWoken                   //2
	mutexStarving                //4
	mutexWaiterShift = iota
)

type TryLockMutex struct {
	sync.Mutex
}

func (tl *TryLockMutex) TryLock() bool {
	// 如果能成功抢到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&tl.Mutex)), 0, mutexLocked) {
		return true
	}

	// 如果处于唤醒、加锁或者饥饿状态，这次请求就不参与竞争了，返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&tl.Mutex))) // 其实是获取mutex的state字段
	if old&(mutexLocked|mutexStarving|mutexStarving) != 0 {
		return false
	}

	// 尝试在竞争的状态下请求锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&tl.Mutex)), old, new)
}

func TestTryLock(t *testing.T) {
	var tl TryLockMutex
	go func() { // 先起一个协程持有锁
		tl.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		tl.Unlock()
	}()

	ok := tl.TryLock()
	if ok {
		fmt.Println("got the lock") // do something mu.Unlock() return
		tl.Unlock()
		return
	}
	fmt.Println("can not get the lock")
}
