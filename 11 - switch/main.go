package main

import "fmt"

func main() {

	example1()
	example2()
}

func greeting(greeting string) {
	fmt.Println(greeting)
}

func example1() {
	venues := []string{"Home", "School", "Bar", "Gym"}

	for _, venue := range venues {
		if venue == "Home" {
			greeting("Mom, i'm hungry")
		} else if venue == "Work" || venue == "School" {
			greeting("Weather is great today")
		} else if venue == "Bar" {
			greeting("Hey i like your dress, but it's a little to blue")
		} else {
			greeting("Sup bois")
		}
	}
}

func example2() {
	venues := []string{"Home", "School", "Bar", "Gym"}

	for _, venue := range venues {
		switch venue {
		case "Home":
			greeting("Mom, i'm hungry")
		case "Work", "School":
			greeting("Weather is great today")
		case "Bar":
			greeting("Hey i like your dress, but it's a little to blue")
		default:
			greeting("Sup bois")
		}
	}
}
