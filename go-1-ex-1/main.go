package main

import "fmt"

var firstName string = "Jacob"
var lastName string = "Sigrist"
var dayOfBirth int = 24
var monthOfBirth int = 5
var yearOfBirth int = 2007
var numberOfSiblings int = 2
var heightInMeters float32 = 1.8
var zodiacSign = '\u264A'

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
