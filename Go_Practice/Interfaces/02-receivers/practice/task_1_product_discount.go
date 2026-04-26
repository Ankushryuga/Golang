package main

type Product struct {
	Name  string
	Price float64
}

// pointer receiver -> modifies original price
func (p *Product) ApplyDiscount(percent float64) {
	p.Price = p.Price * (1 - percent/100)
}

// value recevier
func (p Product) FinalPrice() float64 {
	return p.Price
}
