package main

import (
	"fmt"
	"time"
)

func main() {
	stepMother := NewActor("继母", 52, Woman)
	sisterA := NewActor("姐姐A", 23, Woman)
	sisterB := NewActor("姐姐B", 20, Woman)
	cinderella := NewActor("辛德瑞拉", 18, Woman)

	stepMother.Say("今天欺负一下辛德瑞拉～～～")
	sisterA.Say("今天欺负一下辛德瑞拉～～～")
	sisterB.Say("今天欺负一下辛德瑞拉～～～")
	cinderella.Say("orz...")

	ball := NewBall(19, 27)

	//舞会服装准备
	dressRoom := NewDressRoom()
	dressRoom.Store(stepMother, sisterA, sisterB)

	//继母有连衣裙
	dressRoom.GetDress(stepMother)

	//继母参加舞会
	ball.Entry(stepMother)

	//姐姐A，B有连衣裙，参加舞会
	dressRoom.GetDress(sisterA)
	ball.Entry(sisterA)
	dressRoom.GetDress(sisterB)
	ball.Entry(sisterB)

	//辛德瑞拉没有连衣裙
	dressRoom.GetDress(cinderella)
	ball.Entry(cinderella)

	//魔法变出连衣裙和水晶鞋
	magic := NewMagic(cinderella)
	cinderella.SetCostume(magic.GenerateDress())
	cinderella.SetShoes(magic.GenerateGlassShoes())

	limit := make(chan int, 1)
	go magic.Limit(limit)

	//辛德瑞拉参加了舞会
	ball.Entry(cinderella)

	//王子登场
	prince := NewActor("王子", 18, Man)
	tailcoat := NewTailcoat(prince)
	prince.SetCostume(tailcoat)
	ball.Entry(prince)

	//舞会开始
	ball.Start()
	finished := make(chan int, 1)
	go func() {
		for !ball.IsFinished() {
			<-time.After(1 * time.Second)
			ball.Dancing()

			if ball.Clock == 24 {
				limit <- 1
			}
		}
		ball.Finish()
		finished <- 1
	}()

	<-magic.Broken

	//辛德瑞拉提早回家
	ball.Exit(cinderella)

	//慌忙之中把水晶鞋给掉了
	falledShoes := cinderella.Shoes
	cinderella.Shoes = nil

	//王子捡到了水晶鞋
	foundShoes := falledShoes

	//舞会结束
	<-finished

	//王子开始寻找辛德瑞拉
	fmt.Println("王子开始寻找辛德瑞拉")

	//从参加舞会的女性中寻找, 辛德瑞拉中途离开，不在名单里
	for _, a := range ball.Entries {
		if a.Sex == Woman {
			if foundShoes.Wear(a) {
				fmt.Printf("找到了，是%v 的水晶鞋\n", a.Name)
			} else {
				fmt.Printf("%v: 不是%v 的水晶鞋\n", prince.Name, a.Name)
			}
		}
	}

	if foundShoes.Wear(cinderella) {
		fmt.Printf("找到了，是%v 的水晶鞋\n", cinderella.Name)
	}
}
