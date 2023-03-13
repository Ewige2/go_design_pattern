package bridge_pattern

import "fmt"

// 精确抽象

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("windows  print...")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
