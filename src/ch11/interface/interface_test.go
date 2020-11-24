package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string // 方法名和返回值类型
}

type GoProgrammer struct {
}

// WriteHelloWorld 是方法签名必须和定义的接口一致
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello go world\")"
}

func TestInterface(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
