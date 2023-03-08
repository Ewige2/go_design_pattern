package adapter_pattern

// 一个未知的服务

import "fmt"

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("usb  接口已连入  windows  电脑")
}
