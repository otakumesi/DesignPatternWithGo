package main

import (
	tm "./template_method"
)

func main() {
	tmA := tm.NewConcleteClassA()
	tm.SomeMethod(tmA)
	tmB := tm.NewConcleteClassB()
	tm.SomeMethod(tmB)
}
