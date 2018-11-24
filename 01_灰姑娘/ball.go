package main

import (
	"fmt"
	"sync"
)

// Ball -- 舞会
type Ball struct {
	m          sync.Mutex
	Entries    []*Actor
	Clock      int //时间
	FinishedAt int
}

// NewBall ...
func NewBall(startedAt, finishedAt int) *Ball {
	return &Ball{
		Clock:      startedAt,
		FinishedAt: finishedAt,
	}
}

// Start ...
func (b *Ball) Start() {
	fmt.Printf("舞会开始")
}

// Dancing ...
func (b *Ball) Dancing() {
	b.m.Lock()
	defer b.m.Unlock()
	fmt.Printf("现在%d点\n", b.Clock)
	for _, a := range b.Entries {
		fmt.Printf("%v 在跳舞\n", a.Name)
	}
	b.Clock++
}

// Finish ...
func (b *Ball) Finish() {
	fmt.Println("舞会结束")
}

// IsFinished ...
func (b *Ball) IsFinished() bool {
	return b.Clock >= b.FinishedAt
}

// Entry ...
func (b *Ball) Entry(a *Actor) {
	if a.Cos != nil {
		b.Entries = append(b.Entries, a)
		fmt.Printf("%v 参加了舞会\n", a.Name)
	} else {
		fmt.Println("没有服装不能参加舞会")
		fmt.Printf("%v 不能参加舞会\n", a.Name)
	}
}

// Exit ...
func (b *Ball) Exit(a *Actor) {
	b.m.Lock()
	defer b.m.Unlock()
	var entries []*Actor
	for _, e := range b.Entries {
		if e != a {
			entries = append(entries, e)
		}
	}

	b.Entries = entries
	fmt.Printf("%v 离开舞会回家了\n", a.Name)
}
