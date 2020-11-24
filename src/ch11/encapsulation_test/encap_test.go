package encapsulation_test

import (
	"fmt"
	"testing"
)

type Employee struct {
	ID   string
	Name string
	Age  int
}

func (e Employee) Stringfy() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}

func (e *Employee) Stringfy2() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.ID, e.Name, e.Age)
}

func TestCreatEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.ID = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Logf("e2 is %T", e2)
}

func TestStructOperation(t *testing.T) {
	e := &Employee{"0", "Bob", 20}
	t.Log(e.Stringfy2())
}
