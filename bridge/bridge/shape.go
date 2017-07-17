package bridge

type Shape interface {
	Draw()
}

type Rectangle struct {
	d      Drawing
	x      int
	y      int
	height int
	width  int
}

func NewRectangle(d Drawing, x, y, height, width int) *Rectangle {
	return &Rectangle{
		d:      d,
		x:      x,
		y:      y,
		height: height,
		width:  width,
	}
}

func (r *Rectangle) Draw() {
	r.d.drawLine(r.x, r.y, r.x, r.x+r.width)
	r.d.drawLine(r.x, r.y, r.y, r.y+r.height)
	r.d.drawLine(r.x, r.y+r.height, r.x, r.x+r.width)
	r.d.drawLine(r.x+r.width, r.y, r.y, r.y+r.height)
}

type Circle struct {
	d Drawing
	x int
	y int
	r int
}

func NewCircle(d Drawing, x, y, r int) *Circle {
	return &Circle{
		d: d,
		x: x,
		y: y,
		r: r,
	}
}

func (c *Circle) Draw() {
	c.d.drawCircle(c.x, c.y, c.r)
}
