package abstract_factory

import "testing"

func TestBeijing(t *testing.T) {
	var factory Factory
	factory = &BeijingFactoty{
		Name: "北京",
	}
	factory.CreateChejia().Creating()
	factory.CreateChetai().Creating()
	factory.CreateLunzi().Creating()
}

func TestShanghai(t *testing.T) {
	var factory Factory
	factory = &ShanghaiFactoty{
		Name: "上海",
	}
	factory.CreateChejia().Creating()
	factory.CreateChetai().Creating()
	factory.CreateLunzi().Creating()
}
