package main

import "fmt"

func isLeapsYear(year int) bool {
	// A year is a leap year if it is divisible by 4
	// but years divisible by 100 are not leap years unless they are also divisible by 400
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
	var year int
	fmt.Print("Enter a year to check if it's a leap year: ")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Printf("Leap Year")
	} else {
		fmt.Printf("Not Leap Year")
	}
}
