package main

import "fmt"

func main() {
	list := Todos{}

	list.add("walk")
	list.add("talk")
	fmt.Println(len(list))
}
