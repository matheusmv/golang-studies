package main

import (
	"log"
	"os"
	"path"
	"text/template"
)

type secret struct {
	Name   string
	Secret string
}

func main() {
	secret := secret{"Wallace", "I'm a ninja"}
	file := "template.txt"

	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err.Error())
	}

	templatePath := path.Join(currentDir, file)

	t, err := template.New(file).ParseFiles(templatePath)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = t.Execute(os.Stdout, secret)

	if err != nil {
		log.Fatal(err.Error())
	}
}
