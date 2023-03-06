package prototyping_pattern

import "fmt"

type Folder struct {
	children []Inode //  文件夹内包含多个文件
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolde := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	// 拷贝文件夹内的所有文件的Inode  加入到  tempChildern 中
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolde.children = tempChildren
	return cloneFolde
}
