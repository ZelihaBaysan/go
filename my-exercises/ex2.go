package main

import "fmt"

type Student struct {
	Name string
	Age  int
	ID   string
}

func NewStudent(name string, age int, id string) *Student {
	return &Student{Name: name, Age: age, ID: id}
}

func (s *Student) UpdateAge(newAge int) {
	s.Age = newAge
}

func (s *Student) DisplayInfo() {
	fmt.Printf("Name: %s\nAge: %d\nID: %s\n", s.Name, s.Age, s.ID)
}

func main2() {

	student := NewStudent("Zeliha Baysan", 20, "S12345")

	student.DisplayInfo()

	student.UpdateAge(21)

	student.DisplayInfo()
}
