package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Print(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}
