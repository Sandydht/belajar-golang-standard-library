package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Sandy")
	data.PushBack("Dwi")
	data.PushBack("Handoko")
	data.PushBack("Trapsilo")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
