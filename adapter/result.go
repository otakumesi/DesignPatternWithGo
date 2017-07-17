package main

import (
	"fmt"

	"./adapter"
)

func main() {
	color := adapter.Color("Red")
	var shapes []adapter.Shape
	shapes = append(shapes, adapter.NewPoint(color))
	shapes = append(shapes, adapter.NewSquare(color))
	shapes = append(shapes, adapter.NewCircle(color))
	for _, s := range shapes {
		fmt.Println(s.Display())
	}
}
