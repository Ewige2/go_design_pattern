package generator_pattern

// 生成器接口

type IBuilder interface {
	setWindowType()  //  设置窗户类型
	setDoorType()    // 设置门的类型
	setNumFloor()    // 设置楼层
	getHouse() House // 获取房子
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
