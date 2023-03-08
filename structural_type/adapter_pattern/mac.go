package adapter_pattern

import "fmt"

// 我们需要提供的服务

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("充电器已经插入mac接口")
}
