package bridge

type Drawing interface {
	drawLine(int, int, int, int)
	drawCircle(int, int, int)
}

type V1Drawing struct {
}

func NewV1Drawing() *V1Drawing {
	return &V1Drawing{}
}

func (d *V1Drawing) drawLine(pos_x, pos_y, start, end int) {
	draw_a_line(pos_x, pos_y, start, end)
}

func (d *V1Drawing) drawCircle(x, y, r int) {
	draw_a_circle(x, y, r)
}

type V2Drawing struct {
	dp DP2
}

func NewV2Drawing() *V2Drawing {
	return &V2Drawing{
		dp: DP2{},
	}
}

func (d *V2Drawing) drawLine(pos_x, pos_y, start, end int) {
	d.dp.drawline(pos_x, pos_y, start, end)
}

func (d *V2Drawing) drawCircle(x, y, r int) {
	d.dp.drawcircle(x, y, r)
}
