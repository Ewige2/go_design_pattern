package adapter_pattern

// 适配器，  对windows  需要的服务进行转换

import "fmt"

type WindowsAdapter struct {
	windowMachine *Windows
}

// 通过系统提供的flash 接口  执行 windows 需要的服务
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("windows  适配器啊转换， 将usb 转为充电")
	w.windowMachine.insertIntoUSBPort()
}
