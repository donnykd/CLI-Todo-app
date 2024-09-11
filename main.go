package main

func main() {
	list := Todos{}

	list.add("walk")
	list.toggle(0)

	list.print()
}
