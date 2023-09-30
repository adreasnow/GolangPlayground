package main

import "fmt"

// Throws an error if you input the value 3
func funcName(n int) (int, error) {
	if n == 3 {
		return 0, fmt.Errorf("how DARE you use the number 3")
	}
	return n, nil
}

// Catches the error from funcName()
func testFunc(n int) {
	num1, err := funcName(n)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("num1=%#v\n", num1)
}

func main() {
	testFunc(3)
	testFunc(5)
}
