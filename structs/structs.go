package structs

import (
	"fmt"
	"strconv"
)

type Printer interface {
	Print()
}

type Employee struct {
	JobTitle string
	Person
}

func (e Employee) PrintInfo() {
	fmt.Println(e.Name + " is " + strconv.Itoa(e.Age) + " years old. He lives in " + e.City + ". He works in the job that has title " + e.JobTitle)
}

func (e Employee) Print() {
	e.PrintInfo()
}

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) PrintInfo() {
	fmt.Println(p.Name + " is " + strconv.Itoa(p.Age) + " years old. He lives in " + p.City)
}

func (p Person) Print() {
	p.PrintInfo()
}
