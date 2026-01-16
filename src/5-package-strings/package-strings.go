package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Sandy Dwi Handoko Trapsilo", "Sandy"))
	fmt.Println(strings.Split("Sandy Dwi Handoko Trapsilo", " "))
	fmt.Println(strings.ToLower("Sandy Dwi Handoko Trapsilo"))
	fmt.Println(strings.ToUpper("Sandy Dwi Handoko Trapsilo"))
	fmt.Println(strings.Trim("   Sandy Dwi Handoko Trapsilo  ", " "))
	fmt.Println(strings.ReplaceAll("Sandy Sandy Sandy", "Sandy", "Dwi"))
}
