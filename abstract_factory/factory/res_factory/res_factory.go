package res_factory

import (
	dd "../display_driver"
	pd "../printer_driver"
)

type ResFactory interface {
	DisplayDriver() dd.DisplayDriver
	PrinterDriver() pd.PrinterDriver
}

type LowResFactory struct {
	ResFactory
}

func NewLowResFactory() *LowResFactory {
	return &LowResFactory{}
}

func (lf LowResFactory) DisplayDriver() dd.DisplayDriver {
	return dd.NewLRDD()
}

func (lf LowResFactory) PrinterDriver() pd.PrinterDriver {
	return pd.NewLRPD()
}

type HighResFactory struct {
	ResFactory
}

func NewHighResFactory() *LowResFactory {
	return &LowResFactory{}
}

func (hf HighResFactory) DisplayDriver() dd.DisplayDriver {
	return dd.NewHRDD()
}

func (hf HighResFactory) PrinterDriver() pd.PrinterDriver {
	return pd.NewHRPD()
}
