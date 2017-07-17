package printer_driver

import "fmt"

type PrinterDriver interface {
	RealPrint(string)
}

type LRPD struct {
	PrinterDriver
}

func NewLRPD() *LRPD {
	return &LRPD{}
}

func (lrpd LRPD) RealPrint(s string) {
	fmt.Println(s + "LowResPrinting...")
}

type HRPD struct {
	PrinterDriver
}

func NewHRPD() *HRPD {
	return &HRPD{}
}

func (lrpd HRPD) RealPrint(s string) {
	fmt.Println(s + "HighResPrinting!!!")
}
