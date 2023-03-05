package abstract_factory

// 具体的产品

type NikeShirt struct {
	Shirt
}

type NikeShoe struct {
	Shoe
}

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		// 返回 实现了  IShop接口的 Shop
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}
