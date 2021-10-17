package main

import "fmt"

type weapon interface {
	attack()
}

func attack(weapon weapon) {
	weapon.attack()
}

type sword struct {
	owner string
}

func (sword) attack() {
	fmt.Println("Sword attack!!!")
}

type axe struct {
	owner string
}

func (axe) attack() {
	fmt.Println("Axe attack!!!")
}

func main() {
	weapons := []weapon{
		sword{"Jhon"},
		axe{"Ragnar"},
	}

	for _, weapon := range weapons {
		attack(weapon)
	}

	for _, weapon := range weapons {
		weapon.attack()
	}
}
