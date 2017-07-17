package strategy

type Item struct {
	Name  string
	Price int
}

type SalesOrder struct {
	item     Item
	quantity int
	calcTax  CalcTax
}

func NewSalesOrder(item Item, quantity int, calcTax CalcTax) *SalesOrder {
	return &SalesOrder{item: item, quantity: quantity, calcTax: calcTax}
}

func (s *SalesOrder) CalcBill() int {
	return s.item.Price*s.quantity + int(s.calcTax.taxAmount())
}

type CalcTax interface {
	taxAmount() float64
}

type JapanCalcTax struct {
	price    int
	quantity int
}

func NewJapanTax(price int, quantity int) *JapanCalcTax {
	return &JapanCalcTax{price: price, quantity: quantity}
}

func (c *JapanCalcTax) taxAmount() float64 {
	return float64(c.price*c.quantity) * 0.08
}
