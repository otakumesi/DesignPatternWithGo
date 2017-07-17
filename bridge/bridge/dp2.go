package bridge

import "fmt"

type DP2 struct {
}

func (dp *DP2) drawline(pos_x, pos_y, start, end int) {
	fmt.Println(pos_x, pos_y, start, end)
}

func (dp *DP2) drawcircle(x, y, r int) {
	fmt.Println(x, y, r)
}
