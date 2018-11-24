package main

import "fmt"

// Magic ...
type Magic struct {
	Target *Actor
	Broken chan int
}

// NewMagic ...
func NewMagic(a *Actor) *Magic {
	return &Magic{
		Target: a,
		Broken: make(chan int, 1),
	}
}

// GenerateDress ...
func (m *Magic) GenerateDress() Costume {
	fmt.Printf("%v 获得了魔法道具连衣裙\n", m.Target.Name)
	return NewDress(m.Target)
}

// GenerateGlassShoes ...
func (m *Magic) GenerateGlassShoes() *Shoes {
	fmt.Printf("%v 获得了魔法道具水晶鞋\n", m.Target.Name)
	return NewShoes(m.Target)
}

// Limit ...
func (m *Magic) Limit(limit chan int) {
	<-limit
	fmt.Println("快到晚上0点了")
	fmt.Println("魔法要被解除了")
	m.Broken <- 1
}
