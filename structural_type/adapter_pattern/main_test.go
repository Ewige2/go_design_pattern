package adapter_pattern

import "testing"

func TestClient(t *testing.T) {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}
	// 调用适配器中实现的系统接口， 然后在适配器中调用未知服务，的方法
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

}
