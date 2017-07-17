package main

import (
	"fmt"

	"./strategy"
)

func main() {
	item := strategy.Item{Name: "Yeah", Price: 1000}
	tax := strategy.NewJapanTax(item.Price, 10)
	order := strategy.NewSalesOrder(item, 10, tax)
	fmt.Println(order.CalcBill())
}
