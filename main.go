package main

import "fmt"

func main() {
	list := Todos{}

	list.add("walk")

	list.toggle(0)

	fmt.Println(list[0].CompletedAt)
}
