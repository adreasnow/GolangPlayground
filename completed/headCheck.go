package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

type Head []byte

func readHead(fileName string) (Head, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := make(Head, 8)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, err
}

func compareHead(head Head) (string, error) {

	if bytes.Equal(head[:4], []byte{0x00, 0x00, 0x01, 0x00}) {
		return "ico", nil
	}
	if bytes.Equal(head[:6], []byte{0x47, 0x49, 0x46, 0x38, 0x37, 0x61}) {
		return "gif", nil
	}
	if bytes.Equal(head[:4], []byte{0xFF, 0xD8, 0xFF, 0xF0}) {
		return "jpg", nil
	}
	if bytes.Equal(head[:8], []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}) {
		return "png", nil
	}
	return "", fmt.Errorf("type not identified")
}

func main() {
	path := flag.String("file", "", "The file to cheack the header of")
	flag.Parse()

	head, err := readHead(*path)
	if err != nil {
		fmt.Println(err)
		return
	}
	format, err := compareHead(head)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s --> %s\n", *path, format)
}
