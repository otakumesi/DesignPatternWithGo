package decorator

import "fmt"

type Greeter interface {
	Greet()
}

type Person struct {
	Greeter
	Name string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Greet() {
	fmt.Println("はじめまして！" + p.Name + "です。")
}

type BusinessGreetDecorator struct {
	Greeter
	g Greeter
}

func NewBusinessGreetDecorator(g Greeter) *BusinessGreetDecorator {
	return &BusinessGreetDecorator{
		g: g,
	}
}

func (b *BusinessGreetDecorator) Greet() {
	b.g.Greet()
	fmt.Println("名刺を交換しましょう！")
}
