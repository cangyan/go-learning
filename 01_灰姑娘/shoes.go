package main

// Shoes -- 鞋子
type Shoes struct {
	Owner *Actor
}

// NewShoes -- 生成鞋子
func NewShoes(a *Actor) *Shoes {
	return &Shoes{
		Owner: a,
	}
}

// Wear -- 鞋子只能持有者能穿
func (s *Shoes) Wear(a *Actor) bool {
	return a == s.Owner
}
