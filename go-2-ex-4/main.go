package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		FirstName string
		LastName  string
	}

	// TODO: declare a type for Class (consisting of multiple students)
	type Class []Student

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[string]Class{
		"Math": {
			{FirstName: "Liam", LastName: "Walker"},
			{FirstName: "Timmy", LastName: "Turner"},
		},
		"Science": {
			{FirstName: "Noah", LastName: "Carter"},
			{FirstName: "Emma", LastName: "Mitchell"},
		},
		"History": {
			{FirstName: "Oliver", LastName: "Harris"},
		},
	}

	// TODO: output everything using fmt.Println()
	for module, class := range modules {

		fmt.Println("Module:", module)
		for _, student := range class {
			fmt.Printf("- %s %s\n", student.FirstName, student.LastName)
		}

	}

}
