package main

func main() {
	setMother := NewActor("继母", 52, Woman)
	sisterA := NewActor("姐姐A", 23, Woman)
	sisterB := NewActor("姐姐B", 20, Woman)
	cinderella := NewActor("辛德瑞拉", 18, Woman)

	setMother.Say("今天欺负一下辛德瑞拉～～～")
	sisterA.Say("今天欺负一下辛德瑞拉～～～")
	sisterB.Say("今天欺负一下辛德瑞拉～～～")
	cinderella.Say("orz...")

	ball := NewBall(19, 27)

	//舞会服装准备
	dressRoom := NewDressRoom()
	dressRoom.Store(setMother, sisterA, sisterB)

	//继母有连衣裙
	dressRoom.GetDress(setMother)

	//继母参加舞会
	ball.Entry(setMother)

	//姐姐A，B有连衣裙，参加舞会
	dressRoom.GetDress(sisterA)
	ball.Entry(sisterA)
	dressRoom.GetDress(sisterB)
	ball.Entry(sisterB)

	//辛德瑞拉没有连衣裙

	dressRoom.GetDress(cinderella)
	ball.Entry(cinderella)

	ball.Start()
}
