package bridge_pattern

// 抽象层，  表示计算机

type Computer interface {
	Print()             //  打印操作
	SetPrinter(Printer) //  设置 具体打印机类型
}
