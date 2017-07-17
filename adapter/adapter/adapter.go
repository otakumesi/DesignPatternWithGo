package adapter

import (
	"./x_sys"
)

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
	xcr *x_sys.XCircle
}

func NewCircle(c Color) *Circle {
	xc := x_sys.XColor(string(c))
	x := x_sys.NewXCircle(xc)
	return &Circle{xcr: x}
}

func (ci *Circle) SetColor(c Color) {
	xc := x_sys.XColor(string(c))
	ci.xcr.SetItsColor(xc)
}

func NewSquare(c Color) *Square {
	return &Square{color: c}
}

func (ci *Circle) Display() string {
	return ci.xcr.DisplayIt()
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
