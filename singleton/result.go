package main

import (
	"fmt"
	"reflect"

	"./singleton"
)

func main() {
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()
	if reflect.DeepEqual(s1, s2) {
		fmt.Println("OK!!!")
	}
}
