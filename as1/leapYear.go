package main

import "fmt"

/**
given a year N tell if it is leap year or not
*/

func isLeapYear(year int) bool {
	// A year is a leap year if it is divisible by 4
	// but years divisible by 100 are not leap years unless they are also divisible by 400
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func mainLeapYear() {

	if isLeapYear(2000) {
		fmt.Printf("Leap Year")
	} else {
		fmt.Printf("Not Leap Year")
	}
}
