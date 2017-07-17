package main

import (
	f "./factory"
	fac "./factory/res_factory"
)

func main() {
	l_fac := fac.NewLowResFactory()
	l_ap_control := f.NewApControl(l_fac)
	l_ap_control.PrintOutput("Hello")
	l_ap_control.DisplayOutput("World")

	h_fac := fac.NewHighResFactory()
	h_ap_control := f.NewApControl(h_fac)
	h_ap_control.PrintOutput("Hello")
	h_ap_control.DisplayOutput("World")
}
