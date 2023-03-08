package singleton_pattern

import (
	"log"
	"sync"
	"testing"
	"time"
)

type Singleton2 struct {
	Name string
}

var (
	instance2 *Singleton2
	lock2     sync.RWMutex
)

func GetInstance2() *Singleton2 {
	//  这里为了  避免 数据竞争，  使用粒度更小的读写锁，  读锁和写锁是互斥的
	lock2.RLock()
	lock2.RUnlock()
	if instance2 == nil {
		lock.Lock() //  加写锁
		defer lock.Unlock()
		if instance2 == nil {
			instance2 = &Singleton2{Name: "zhangsan2"}
		}
	}
	return instance2
}

func TestClent2(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			resp := GetInstance2()
			log.Printf("the value is : %s", resp.Name)
		}()
	}
	time.Sleep(time.Second)
}
