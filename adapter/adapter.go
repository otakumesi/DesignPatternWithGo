package main

import "fmt"

type Shape interface {
	Display() string
	SetColor(Color)
}

type Color string

type Point struct {
	color Color
}

type Square struct {
	color Color
}

/* 他のシステムを利用するCircleオブジェクト */

type Circle struct {
	xCircle *XCircle
}

func NewCircle(c Color) *Circle {
	x := NewXCircle(c)
	return &Circle{xCircle: x}
}

func (ci *Circle) SetColor(c Color) {
	ci.xCircle.SetItsColor(c)
}

func NewSquare(c Color) *Square {
	return &Square{color: c}
}

func (ci *Circle) Display() string {
	return ci.xCircle.DisplayIt()
}

func NewPoint(c Color) *Point {
	return &Point{color: c}
}

func (p *Point) SetColor(c Color) {
	p.color = c
}

func (p *Point) Display() string {
	return string(p.color) + "Point"
}

func (p *Square) SetColor(c Color) {
	p.color = c
}

func (p *Square) Display() string {
	return string(p.color) + "Square"
}

/* 他のシステムのStruct */

type XCircle struct {
	color Color
}

func NewXCircle(c Color) *XCircle {
	return &XCircle{color: c}
}

func (x *XCircle) SetItsColor(c Color) {
	x.color = c
}

func (x *XCircle) DisplayIt() string {
	return string(x.color) + "Circle"
}

func main() {
	color := Color("Red")
	var shapes []Shape
	shapes = append(shapes, NewPoint(color))
	shapes = append(shapes, NewSquare(color))
	shapes = append(shapes, NewCircle(color))
	for _, s := range shapes {
		fmt.Println(s.Display())
	}
}
