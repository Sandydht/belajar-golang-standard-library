package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("sandy"))
	fmt.Println(regex.MatchString("dwi"))
	fmt.Println(regex.MatchString("sAndy"))

	fmt.Println(regex.FindAllString("eko edo egi ego e10 eto", 10))
}
