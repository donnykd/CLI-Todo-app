package main

func main() {
	list := Todos{}
	space := NewStorage[Todos]("todos.json")
	space.Load(&list)
	cmd := NewCmd()
	cmd.Execute(&list)
	space.Save(list)
}
