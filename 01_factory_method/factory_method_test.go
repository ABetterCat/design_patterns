package factory_method

import (
	"fmt"
	"testing"
)

func TestBJ(t *testing.T) {
	var factory FactoryOperator
	factory = BJFactory{}
	a := factory.Create()
	fmt.Println(a.GetName())
	if a.GetName() != "北京" {
		t.Fatal("错误")
	}
}

func TestGY(t *testing.T) {
	var factory FactoryOperator
	factory = GYFactory{}
	a := factory.Create()
	fmt.Println(a.GetName())
	if a.GetName() != "贵阳" {
		t.Fatal("错误")
	}
}
