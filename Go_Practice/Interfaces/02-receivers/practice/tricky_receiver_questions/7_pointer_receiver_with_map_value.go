package main

type Pointer_Receiver_With_Map_Value struct {
	Name string
}

func (p *Pointer_Receiver_With_Map_Value) ChangeName(name string) {
	p.Name = name
}
