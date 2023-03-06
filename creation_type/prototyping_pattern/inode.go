package prototyping_pattern

// 原型接口，  模拟操作系统 文件
// 每个文件和文件夹都可用一个 inode接口来表示。 inode接口中同样也有 clone克隆功能。

type Inode interface {
	print(string)
	clone() Inode
}
