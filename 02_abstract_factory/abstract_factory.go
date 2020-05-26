package abstract_factory

import "fmt"

// 定义小轮车工厂,产品
type Factory interface {
	CreateLunzi() Product
	CreateChetai() Product
	CreateChejia() Product
}
type Product interface {
	Creating()
}

// 实现上海小轮车
type ShanghaiFactoty struct {
	Name string
}

func (s *ShanghaiFactoty) CreateLunzi() Product {
	return &Lunzi{
		ProductName: s.Name,
		Name:        "轮子",
	}
}

func (s *ShanghaiFactoty) CreateChetai() Product {
	return &Luntai{
		ProductName: s.Name,
		Name:        "轮胎",
	}
}

func (s *ShanghaiFactoty) CreateChejia() Product {
	return &Chejia{
		ProductName: s.Name,
		Name:        "车架",
	}
}

// 实现北京小轮车
type BeijingFactoty struct {
	Name string
}

func (s *BeijingFactoty) CreateLunzi() Product {
	return &Lunzi{
		ProductName: s.Name,
		Name:        "轮子",
	}
}

func (s *BeijingFactoty) CreateChetai() Product {
	return &Luntai{
		ProductName: s.Name,
		Name:        "轮胎",
	}
}

func (s *BeijingFactoty) CreateChejia() Product {
	return &Chejia{
		ProductName: s.Name,
		Name:        "车架",
	}
}

// 实现产品
// 轮子
type Lunzi struct {
	ProductName string
	Name        string
}

type Luntai struct {
	ProductName string
	Name        string
}

type Chejia struct {
	ProductName string
	Name        string
}

func (l *Lunzi) Creating() {
	fmt.Println(fmt.Sprintf("%s的工厂生产%s中", l.ProductName, l.Name))
}

func (l *Luntai) Creating() {
	fmt.Println(fmt.Sprintf("%s的工厂生产%s中", l.ProductName, l.Name))
}

func (l *Chejia) Creating() {
	fmt.Println(fmt.Sprintf("%s的工厂生产%s中", l.ProductName, l.Name))
}
