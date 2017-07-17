package display_driver

import "fmt"

type DisplayDriver interface {
	Display(string)
}

type LRDD struct {
	DisplayDriver
}

func NewLRDD() *LRDD {
	return &LRDD{}
}

func (lrdd LRDD) Display(s string) {
	fmt.Println(s + "Display LowRes...")
}

type HRDD struct {
	DisplayDriver
}

func NewHRDD() *HRDD {
	return &HRDD{}
}

func (hrdd HRDD) Display(s string) {
	fmt.Println(s + "Display HighRes!!!")
}
