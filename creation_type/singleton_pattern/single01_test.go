package singleton_pattern

import (
	"log"
	"sync"
	"testing"
	"time"
)

// 使用双重检查锁， 实现单例

/*
go run  --race   single01_test.go  检查是否存在数据竞争
这个简易版的存在数据竞争
*/

type Singleton struct {
	Name string
}

var (
	instance *Singleton
	lock     sync.Mutex
)

func GetInstance() *Singleton {
	if instance == nil { //  判断实例是否为空(会存在数据竞争，其他线程会进入读取instance的值,  进行判断)
		lock.Lock() //  加锁
		defer lock.Unlock()
		if instance == nil { //   初始化
			instance = &Singleton{Name: "zhangsan"}
		}
	}
	return instance
}

func TestClent(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			resp := GetInstance()
			log.Printf("the value is : %s", resp.Name)
		}()
	}
	time.Sleep(time.Second)
}
