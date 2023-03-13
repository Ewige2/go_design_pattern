package combinatorial_pattern

import "testing"

func Test(t *testing.T) {
	file2 := &File{name: "file2"}
	file1 := &File{name: "file1"}
	file3 := &File{name: "file3"}

	folder1 := &Folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &Folder{
		name: "Folder2",
	}

	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)
	folder2.search("hali")

}
