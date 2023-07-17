package main

import (
	"fmt"
	"os"
	"bufio"
)

type Player struct {
	Name string
	LifePoints int
	Class string
	Level int
	Race string
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) SetName(newName string) {
	p.Name = newName
}

func (p *Player) GetLifePoints() int {
	return p.LifePoints
}

func (p *Player) SetLifePoints(newLifePoints int) {
	p.LifePoints = newLifePoints
}

func (p *Player) GetClass() string {
	return p.Class
}

func (p *Player) SetClass(newClass string) {
	p.Class = newClass
}

func (p *Player) GetLevel() int {
	return p.Level
}

func (p *Player) SetLevel(newLevel int) {
	p.Level = newLevel
}

func (p *Player) GetRace() string {
	return p.Race
}

func (p *Player) SetRace(newRace string) {
	p.Race = newRace
} 

func CreatePlayer(){
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("What is the name of your Player: ")
	var nome, _ = reader.ReadString('\n')
	fmt.Print("What is the name of your Class: ")
	var class, _ = reader.ReadString('\n')
	fmt.Print("What is the name of your Race: ")
	var race, _ = reader.ReadString('\n')
	player := Player{
		Name: nome,
		LifePoints: 100,
		Level: 1,
		Class: class,
		Race: race,
	}
	fmt.Print(fmt.Sprintf("\nHello %s\n---Here are your stats:---\nLife: %d\nLevel: %d", player.Name, player.LifePoints, player.Level))
}