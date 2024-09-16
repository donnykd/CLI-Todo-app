package main

func main() {
	list := Todos{}
	space := NewStorage[Todos]("todos.json")
	list.add("walk")
	list.add("groceries")
	list.toggle(1)
	space.save(list)
	list.print()
}
