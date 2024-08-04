package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func UpdatePerson(p *Person, newName string, newAge int) {

	p.Name = newName
	p.Age = newAge
}

func main() {
	person := Person{
		Name: "zelis",
		Age:  20,
	}

	fmt.Println("Kişi Bilgileri (Başlangıç):")
	fmt.Println("Adı:", person.Name)
	fmt.Println("Yaşı:", person.Age)

	UpdatePerson(&person, "Boz", 4)

	fmt.Println("\nKişi Bilgileri (Güncellenmiş):")
	fmt.Println("Adı:", person.Name)
	fmt.Println("Yaşı:", person.Age)
}
