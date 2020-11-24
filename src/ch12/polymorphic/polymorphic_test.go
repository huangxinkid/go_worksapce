package interface_test

import (
	"fmt"
	"testing"
)

type Code string // custom type
type Programmer interface {
	WriteHelloWorld() Code // 方法名和返回值类型
}

type GoProgrammer struct {
}

// WriteHelloWorld 是方法签名必须和定义的接口一致
func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello world\")"
}

type JavaProgrammer struct {
}

// WriteHelloWorld 是方法签名必须和定义的接口一致
func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out(\"Hello world\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
