package factory

import (
	dd "./display_driver"
	pd "./printer_driver"
	fac "./res_factory"
)

type ApControl struct {
	printer   pd.PrinterDriver
	displayer dd.DisplayDriver
}

func NewApControl(rf fac.ResFactory) *ApControl {
	return &ApControl{
		printer:   rf.PrinterDriver(),
		displayer: rf.DisplayDriver(),
	}
}

func (ap ApControl) PrintOutput(s string) {
	ap.printer.RealPrint(s)
}

func (ap ApControl) DisplayOutput(s string) {
	ap.displayer.Display(s)
}
