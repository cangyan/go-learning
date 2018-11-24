package main

// Costume -- 服装接口
type Costume interface {
	Wear(a *Actor) bool
}

// Dress -- 连衣裙
type Dress struct {
	Owner *Actor
}

// NewDress -- 生成连衣裙
func NewDress(a *Actor) Costume {
	return &Dress{
		Owner: a,
	}
}

// Wear -- 只有女性且所有者才能穿上
func (d *Dress) Wear(a *Actor) bool {
	return a == d.Owner && a.Sex == Woman
}

// Tailcoat -- 燕尾服
type Tailcoat struct {
	Owner *Actor
}

// NewTailcoat -- 生成燕尾服
func NewTailcoat(a *Actor) Costume {
	return &Tailcoat{
		Owner: a,
	}
}

// Wear -- 只有男性且所有者才能穿上
func (t *Tailcoat) Wear(a *Actor) bool {
	return a == t.Owner && a.Sex == Man
}
