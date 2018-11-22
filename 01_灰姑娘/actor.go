package main

import "fmt"

// Sex -- 性别
type Sex int

//
const (
	_ Sex = iota
	Woman
	Man
)

// Actor -- 登场角色
type Actor struct {
	Name  string
	Age   int
	Sex   Sex
	Cos   Costume
	Shoes *Shoes
}

// NewActor -- 创建角色
func NewActor(n string, age int, s Sex) *Actor {
	return &Actor{
		Name: n,
		Age:  age,
		Sex:  s,
	}
}

// Say -- 角色-说
func (a *Actor) Say(s string) {
	fmt.Printf("%v: %v\n", a.Name, s)
}

// SetCostume -- 角色-穿服装
func (a *Actor) SetCostume(c Costume) {
	a.Cos = c
}

// SetShoes -- 角色-穿鞋
func (a *Actor) SetShoes(s *Shoes) {
	a.Shoes = s
}
