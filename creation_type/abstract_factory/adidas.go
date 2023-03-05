package abstract_factory

//   adidas 的具体工厂，   可以 生产  该品牌下的各种商品

// 具体的产品

type AdidasShirt struct {
	Shirt
}

type AdidasShoe struct {
	Shoe
}

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
