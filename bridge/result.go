package main

import (
	"./bridge"
)

func main() {
	var shapes []bridge.Shape
	d := bridge.NewV1Drawing()
	shapes = append(shapes, bridge.NewRectangle(d, 50, 50, 100, 100))

	d2 := bridge.NewV2Drawing()
	shapes = append(shapes, bridge.NewCircle(d2, 50, 50, 15))
	for _, shape := range shapes {
		shape.Draw()
	}
}
