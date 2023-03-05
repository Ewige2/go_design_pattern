package abstract_factory

import "fmt"

// 抽象工厂接口 （用于生产体育用品）

type ISportsFactory interface {
	// 不同的生产线， 市场不同的产品（都有共同的特点  logo，  尺码）
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
