package template_method

import "fmt"

type BaseMethods interface {
	SomeMethod()
	concleteMethod1()
	concleteMethod2()
}

type BaseClass struct {
	BaseMethods
}

func SomeMethod(b BaseMethods) {
	fmt.Println("TemplateMethod start")
	b.concleteMethod1()
	fmt.Println("TempleteMethod middle...")
	b.concleteMethod2()
	fmt.Println("TemplateMethod finish...")
}

type ConcleteClassA struct {
	BaseClass
}

func NewConcleteClassA() *ConcleteClassA {
	return &ConcleteClassA{}
}

func (c ConcleteClassA) concleteMethod1() {
	fmt.Println("ConcleteMethod1 powererd by ConcleteClassA")
}

func (c ConcleteClassA) concleteMethod2() {
	fmt.Println("ConcleteMethod2 powererd bby ConcleteClassA")
}

type ConcleteClassB struct {
	BaseClass
}

func NewConcleteClassB() *ConcleteClassB {
	return &ConcleteClassB{}
}

func (c ConcleteClassB) concleteMethod1() {
	fmt.Println("ConcleteMethod1 powererd by ConcleteClassB")
}

func (c ConcleteClassB) concleteMethod2() {
	fmt.Println("ConcleteMethod2 powererd bby ConcleteClassB")
}
