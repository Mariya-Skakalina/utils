package utils

import "fmt"

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			fmt.Print("True")
			return true
		}
	}
	fmt.Print("False")
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			fmt.Print("True")
			return true
		}
	}
	fmt.Print("False")
	return false
}
