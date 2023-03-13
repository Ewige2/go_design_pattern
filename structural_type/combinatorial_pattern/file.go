package combinatorial_pattern

// 组件接口

import "fmt"

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf(" 搜索关键字:  %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}
