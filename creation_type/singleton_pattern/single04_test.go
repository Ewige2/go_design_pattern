package singleton_pattern

// go  内部 提供的  双重检查锁

import (
	"log"
	"sync"
	"testing"
	"time"
)

type Singleton4 struct {
	Name string
}

var (
	instance4 *Singleton4
	once      sync.Once
)

func Getinstance4() *Singleton4 {
	// 本质还是 调用  atomic  包
	once.Do(func() {
		instance4 = &Singleton4{Name: "zhangsan"}
	})
	return instance4
}

func TestClient4(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			res := Getinstance4()
			log.Printf("the value is : %s", res.Name)
		}()
	}
	time.Sleep(time.Second)
}
