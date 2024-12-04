package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Programming Essentials",
		117: "Algorithm Design and Analysis",
		346: "Mastering Go Development",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 104)
	// TODO: add one
	modules[666] = "System-Database"
	// TODO: replace one
	modules[346] = "Programierung der Database"

	fmt.Println(modules)
}
