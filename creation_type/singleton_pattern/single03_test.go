package singleton_pattern

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
在「简易版」中出现数据竞争的原因是:  instance 变量在同一时刻既有 goroutine 读，也有 goroutine 写，
本质原因是变量的修改非原子操作，而 golang 中提供了一个原子操作 package atomic ，
atomic 的逻辑是实现在硬件层面之上，其意味着即使有多个 goroutine 修改同一个 atomic 变量，
该变量也会正确更新且不会发生数据竞争。

*/

type Singleton3 struct {
	Name string
}

var (
	instance3 *Singleton3
	lock3     sync.Mutex
	flag3     uint32
)

func GetInstance3() *Singleton3 {
	if atomic.LoadUint32(&flag3) == 0 {
		lock3.Lock()
		defer lock3.Unlock()
		if atomic.LoadUint32(&flag3) == 0 {
			instance3 = &Singleton3{Name: "zhangsan3"}
			defer atomic.StoreUint32(&flag3, 1)
		}
	}
	return instance3
}

func TestClient3(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			resp := GetInstance3()
			log.Printf("the value is : %s", resp.Name)
		}()
	}
	time.Sleep(time.Second)
}
