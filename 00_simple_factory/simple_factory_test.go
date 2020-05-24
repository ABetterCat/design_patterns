package simple_factory

import (
	"fmt"
	"testing"
)

func TestBj(t *testing.T) {
	esRes := ResESStruct("false")
	res := esRes.GetName()
	fmt.Println(res)
	if res != "北京" {
		t.Fatal("失败")
	}
}

func TestGy(t *testing.T) {
	esRes := ResESStruct("true")
	if esRes.GetName() != "贵阳" {
		t.Fatal("失败")
	}
}
