package main

import "fmt"

// 简单工厂模式

/*
基本思路
1. 定义一个基础数据库相关的结构体
2.需要实现接口，初始化连接，返回连接信息
3.创建一个工厂，这个工厂需要实现接口， 可以创建连接
4.后续工厂如果需要实现其他功能，直接在接口中定义让后实现即可， 不需要修改其他代码

应用场景
1. 你在编写代码的过程中，如果无法预知对象确切类别及其
依赖关系时，可使用工厂方法。
2. 工厂方法将创建产品的代码与实际使用产品的代码分离，从
而能在不影响其他代码的情况下扩展产品创建部分代码。
3. 例如，如果需要向应用中添加一种新产品，你只需要开发新
的创建者子类，然后重写其工厂方法即可。
4. 如果你希望用户能扩展你软件库或框架的内部组件，可使用工厂方法。

缺点
1. 简单工厂模式的工厂类单一，负责所有产品的创建，职责过重，一旦异常，整个系统将受影响。且工厂类代码会非常臃肿，违背高内聚原则.
2. 使用简单工厂模式会增加系统中类的个数（引入新的工厂类），增加系统的复杂度和理解难度.
3. 系统扩展困难，一旦增加新产品不得不修改工厂逻辑，在产品类型较多时，可能造成逻辑过于复杂.
4. 该接口必须对所有产品都有意义


*/

type BaseDB struct {
	address string
	port    int64
	account string
	passwd  string
}

type DBconnect interface {
	Init(address string, port int64, account string, passwd string)
	Connect() string
}

func (this *BaseDB) Init(address string, port int64, account string, passwd string) {
	this.address = address
	this.port = port
	this.account = account
	this.passwd = passwd
}

type MysqlDB struct {
	BaseDB
}

func (mdb *MysqlDB) Connect() string {
	return fmt.Sprintf("连接了mysql,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

type MongoDB struct {
	BaseDB
}

func (mdb *MongoDB) Connect() string {
	return fmt.Sprintf("mongo,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

// CreateDBFactory 创建数据库连接的工厂
type CreateDBFactory struct {
}

// CreateDBInter 定义工厂需要实现的接口
type CreateDBInter interface {
	CreateDBConnect(dbname string) DBconnect
}

func (cdb *CreateDBFactory) CreateDBConnect(dbname string) (bdb DBconnect, err error) {
	switch dbname {
	case "mysql":
		bdb = &MysqlDB{}
	case "mongodb":
		bdb = &MongoDB{}
	default:
		bdb = &MysqlDB{}
	}
	return
}

func main() {
	// 初始化工厂
	dbFactory := &CreateDBFactory{}
	// 常见Mysql连接
	mysqlObject, err := dbFactory.CreateDBConnect("mysql")
	if err != nil {
		panic("mysql  error")
	}
	// 初始化mysql 连接
	mysqlObject.Init("mysqladdress", 3306, "root", "root")
	mysqlConnect := mysqlObject.Connect()
	fmt.Println("mysql connect is " + mysqlConnect)

}
