package adapter_pattern

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("用户将充电器插入计算机接口")
	com.InsertIntoLightningPort()
}
