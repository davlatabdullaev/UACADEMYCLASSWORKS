package main

import (
	"fmt"
	"math/rand"
)

const chances = 5

func main() {
	for true {
		player := Player{}

		game := Game{}

		name := getPlayerName()

		newPlayer := player.NewPlayer(name, chances)

		newGame := game.NewGame(newPlayer)

		newGame.StartGame()
	}
}

type Game struct {
	RandomNumber int
	Player       Player
}

func (g Game) NewGame(player Player) Game {
	return Game{
		RandomNumber: rand.Intn(50),
		Player:       player,
	}
}
func (g Game) StartGame() {
	fmt.Printf("Salom %s\n", g.Player.Name)
	fmt.Println(" Bu kompyuter o'ylagan sonni topish o'yini \n Siz 1 dan 50 gacha bolgan sonlarni kiritishingiz mumkin \n Sizda 5 ta imkoniyat bor ")
	for i := 0; i < g.Player.Chances; i++ {
		var n int
		fmt.Printf("%d - urinish \n son kiriting: ", i+1)
		fmt.Scan(&n)

		if n == g.RandomNumber {
			fmt.Println("TABRIKLAYMIZ")
			return
		} else if n > g.RandomNumber {
			fmt.Println("Random number siz kiritgan sondan kichkina")
		} else if n < g.RandomNumber {
			fmt.Println("Random number siz kiritgan sondam katta ")
		} else {
			fmt.Println("Incorrect")
		}
	}
	fmt.Println(" Yutqazdingiz\n kompyuter o'ylagan son ", g.RandomNumber, " edi")
}

type Player struct {
	Name    string
	Chances int
}

func (p Player) NewPlayer(name string, chances int) Player {
	return Player{
		Name:    name,
		Chances: chances,
	}
}
func getPlayerName() string {
	var (
		name string
	)
	fmt.Print("Ismingizni kiriting : ")
	fmt.Scan(&name)

	return name
}
