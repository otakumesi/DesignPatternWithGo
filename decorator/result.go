package main

import (
	d "./decorator"
)

func main() {
	takuro := d.NewPerson("Takuro")
	takuro.Greet()

	business_takuro := d.NewBusinessGreetDecorator(d.NewPerson("Takuro"))
	business_takuro.Greet()
}
