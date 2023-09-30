package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, err
}

func main() {
	input, err := readFile("/Users/adrea/.zshrc")
	if err != nil {
		return
	}
	for _, s := range input {
		fmt.Println(s)

	}
}
