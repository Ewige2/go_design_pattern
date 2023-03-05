package abstract_factory

import (
	"fmt"
	"testing"
)

// 测试文件必须是   xxx_test 格式

func TestGet(t *testing.T) {
	// 创建不同品牌的工厂
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	// 创建不同品牌下的不同生产线
	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())

}
