package main

type Celsius float64

// value receiver
func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

// value receiver
func (c Celsius) ToKelvin() float64 {
	return float64(c) + 273.15
}
