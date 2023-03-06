package prototyping_pattern

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nfloder2  的层次结构")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\n打印克隆文件夹的层次结构")
	cloneFolder.print("  ")
}
