package bridge_pattern

// 具体实施：  具体的打印机（爱普生）

import "fmt"

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
