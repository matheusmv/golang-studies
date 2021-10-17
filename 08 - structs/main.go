package main

import (
	"fmt"
	"time"
)

type address struct {
	street       string
	neighborhood string
	city         string
}

type person struct {
	name      string
	email     string
	password  string
	address   address
	createdAt string
}

func createPerson(name string, email string, password string) person {
	return person{
		name:      name,
		email:     email,
		password:  password,
		createdAt: time.Now().Format(time.UnixDate),
	}
}

func (p *person) setName(name string) {
	p.name = name
}

func (p *person) setEmail(email string) {
	p.email = email
}

func (p *person) setPassword(password string) {
	p.password = password
}

func (p *person) setCreatedAt(createdAt string) {
	p.createdAt = createdAt
}

func (p *person) setAddress(street string, neighborhood string, city string) {
	p.address.street = street
	p.address.neighborhood = neighborhood
	p.address.city = city
}

func (person) sayHello() {
	fmt.Println("Hello")
}

func (p person) toString() {
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Email: %s\n", p.email)
	fmt.Printf("Password: %s\n", p.password)
	fmt.Printf("Address: %s %s %s\n",
		p.address.city, p.address.neighborhood, p.address.street)
	fmt.Printf("createdAt: %s\n", p.createdAt)
}

func main() {

	// person1 := person{
	// 	name:      "Matheus Matias Viana",
	// 	email:     "matheus@email.com",
	// 	password:  "123456789",
	// 	createdAt: time.Now().Local().String(),
	// }

	person1 := createPerson("Matheus", "matheus@email", "123456")

	person1.toString()

	person1.setAddress("Rua 13", "Aeroporto", "Aracati")

	person1.toString()

	person1.address.city = "Fortaleza"

	person1.toString()

	person1.sayHello()

	person2 := new(person)

	person2.setName("Matheus 2")
	person2.setEmail("matheus2@email")
	person2.setPassword("123456")
	person2.setCreatedAt(time.Now().Format(time.UnixDate))
	person2.setAddress("Rua 13", "Aeroporto", "Aracati")
	person2.toString()

	fmt.Println(person1)
	fmt.Println(person2)

	example1()
	example2()
	example3()
	example4()
	example5()
}

func example1() {
	type publisher struct {
		name string
	}

	type game struct {
		name      string
		publisher *publisher
	}

	publisher1 := publisher{
		name: "Blizard",
	}

	game1 := game{
		name:      "CoD",
		publisher: &publisher1,
	}

	fmt.Println(game1.publisher.name)

	game1.publisher.name = "Test"

	fmt.Println(game1.publisher.name)
}

func example2() {
	type ninja struct {
		name    string
		weapons []string
		level   int
	}

	ninja1 := ninja{name: "Wallace"}

	fmt.Println(ninja1)

	ninja1 = ninja{"Wallace", []string{"Ninja Star"}, 1}

	fmt.Println(ninja1.name)
	ninja1.weapons = append(ninja1.weapons, "Ninja Sword")
	fmt.Println(ninja1.weapons)
	ninja1.level++
	fmt.Println(ninja1.level)
}

func example3() {
	type ninja struct {
		name    string
		weapons []string
		level   int
	}

	type dojo struct {
		name  string
		ninja ninja
	}

	ninja1 := ninja{"Wallace", []string{"Ninja Star", "Ninja Sword"}, 3}

	golangDojo := dojo{
		name:  "Golang Dojo",
		ninja: ninja1,
	}

	fmt.Println(golangDojo.name)
	fmt.Println(golangDojo.ninja.name)
	fmt.Println(golangDojo.ninja.weapons)
	fmt.Println(golangDojo.ninja.level)
}

func example4() {
	type ninja struct {
		name    string
		weapons []string
		level   int
	}

	type dojo struct {
		name  string
		ninja *ninja
	}

	ninja1 := ninja{"Wallace", []string{"Ninja Star", "Ninja Sword"}, 3}

	golangDojo := dojo{
		name:  "Golang Dojo",
		ninja: &ninja1,
	}

	fmt.Println(golangDojo)
	fmt.Println(golangDojo.ninja)
	fmt.Println(*golangDojo.ninja)
	fmt.Println(golangDojo.name)
	fmt.Println(golangDojo.ninja.name)
	fmt.Println(golangDojo.ninja.weapons)
	fmt.Println(golangDojo.ninja.level)

	golangDojo.ninja.level = 4

	fmt.Println(golangDojo.ninja.level)
	fmt.Println(ninja1.level)
	fmt.Printf("%p\n", &ninja1)
	fmt.Printf("%p\n", golangDojo.ninja)
	fmt.Println(&ninja1 == golangDojo.ninja)
}

func example5() {
	type ninja struct {
		name    string
		weapons []string
		level   int
	}

	type dojo struct {
		name  string
		ninja *ninja
	}

	ninja1 := new(ninja)

	ninja1.name = "Jonhy"
	ninja1.weapons = []string{"Ninja Start"}
	ninja1.level = 1

	golangDojo := dojo{
		name:  "Golang Dojo",
		ninja: ninja1,
	}

	fmt.Println(golangDojo)
	fmt.Println(*golangDojo.ninja)
	fmt.Printf("%p\n", ninja1)
	fmt.Printf("%p\n", *&golangDojo.ninja)
	fmt.Println(ninja1 == *&golangDojo.ninja)
}
