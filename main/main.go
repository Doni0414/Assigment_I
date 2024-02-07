package main

import (
	"Assignment_I/structs"
	"Assignment_I/utils"
	"fmt"
	"strconv"
)

func main() {
	// Part 1: Go Syntax [20 points]
	// Variables and Types:
	// Declare a variable of type int and initialize it with a value of your choice.
	fmt.Println("Part 1")
	var a int = 7
	// Declare a variable of type string and initialize it with your name.
	var name string = "Nurdaulet"
	// Print the values of both variables.
	fmt.Println(a)
	fmt.Println(name)
	// Control Flow:
	// Write a function that takes an integer parameter and prints whether it's even or odd.
	// printEvenOrOdd(5)
	utils.PrintEvenOrOdd(5)
	// Use a loop to print numbers from 1 to 5.
	utils.PrintRange(1, 5)
	// Arrays and Slices:
	// Create an array of integers with values 1, 2, 3, 4, and 5.
	arr := [5]int{1, 2, 3, 4, 5}
	// Use a slice to extract a subset of the array (e.g., from index 1 to 3).
	slc := arr[1:3]
	// Print the elements of the slice.
	fmt.Println(slc)
	// Maps:
	// Create a map that represents a person's information with keys: "name", "age", and "city".
	person := map[string]any{
		"name": "Askar",
		"age":  20,
		"city": "Almaty",
	}
	// Print the person's information.
	fmt.Println(person["name"].(string) + " is " + strconv.Itoa(person["age"].(int)) + " years old. He lives in " + person["city"].(string))

	fmt.Println()
	fmt.Println("Part 2")
	// 	Part 2: OOP [40 points]
	// Define a Struct:
	// Create a struct named Person with fields: Name, Age, and City.
	// Methods:
	// Add a method to the Person struct named PrintInfo that prints the person's information.
	// Create Instances:
	// Create two instances of the Person struct with different information.
	person1 := &structs.Person{
		Name: "Ramazan", Age: 20, City: "Astana",
	}
	person2 := &structs.Person{
		Name: "Asset", Age: 30, City: "Pavlodar",
	}

	// Call the PrintInfo method for each instance to display their information.
	person1.PrintInfo()
	person2.PrintInfo()
	// Embedding:
	// Create a new struct named Employee that embeds the Person struct.
	// Add a new field to the Employee struct for the employee's job title.
	// Add a method to the Employee struct named PrintEmployeeInfo that prints both personal and job-related information.
	// Interfaces:
	// Define an interface named Printer with a method Print.
	// Make both Person and Employee structs implement the Printer interface by providing a Print method for each.
	// Polymorphism:
	// Create a function named DisplayInfo that takes a Printer interface as a parameter and calls its Print method.
	// Use this function to display information for both a Person and an Employee.
	employee := &structs.Employee{
		JobTitle: "Software Engineer",
		Person: structs.Person{
			Name: "Nuralibek", Age: 18, City: "Taraz",
		},
	}
	structs.DisplayInfo(person1)
	structs.DisplayInfo(employee)
}
