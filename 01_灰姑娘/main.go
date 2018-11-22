package main

func main() {
	setMother := NewActor("继母", 52, Woman)
	sisterA := NewActor("大姐", 23, Woman)
	sisterB := NewActor("二姐", 20, Woman)
	cinderella := NewActor("辛德瑞拉", 18, Woman)

	setMother.Say("今天欺负一下辛德瑞拉～～～")
	sisterA.Say("今天欺负一下辛德瑞拉～～～")
	sisterB.Say("今天欺负一下辛德瑞拉～～～")
	cinderella.Say("orz...")
}
