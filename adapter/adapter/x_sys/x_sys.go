package x_sys

type XColor string

type XCircle struct {
	color XColor
}

func NewXCircle(c XColor) *XCircle {
	return &XCircle{color: c}
}

func (x *XCircle) SetItsColor(c XColor) {
	x.color = c
}

func (x *XCircle) DisplayIt() string {
	return string(x.color) + "Circle"
}
