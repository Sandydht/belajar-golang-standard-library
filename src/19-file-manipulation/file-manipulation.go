package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		message += string(line)
		if err == io.EOF {
			break
		}
	}

	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	createNewFile("example.txt", "Hello, world\n")
	readFile("example.txt")
	addToFile("example.txt", "Appending this line.\n")
}
