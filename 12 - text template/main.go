package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name   string
	Secret string
}

func main() {
	secret := secret{"Wallace", "I'm a ninja"}

	templatePath := "/home/matheus/Documents/Golang/12 - text template/template.txt"

	t, err := template.New("template.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, secret)

	if err != nil {
		fmt.Println(err)
	}
}
