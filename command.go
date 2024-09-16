package main

import (
	"flag"
	"fmt"
	"strconv"
)

type cmd struct {
	Add    string
	Del    interface{}
	Edit   interface{}
	Toggle interface{}
	List   bool
}

func NewCmd() *cmd {
	c := cmd{}

	var delStr, editStr, toggleStr string

	flag.StringVar(&c.Add, "add", "", "Add a todo")
	flag.StringVar(&delStr, "del", "", "Delete a todo (accepts index or title)")
	flag.StringVar(&editStr, "edit", "", "Edit a todo (accepts index or title)")
	flag.StringVar(&toggleStr, "toggle", "", "Change a todo's status (accepts index or title)")
	flag.BoolVar(&c.List, "list", false, "list all todos")

	flag.Parse()

	//c.Del = CheckIfInt(delStr)
	c.Edit = editStr
	//c.Toggle = CheckIfInt(toggleStr)

	return &c
}

func CheckIfInt(s string) interface{} {
	i, err := strconv.Atoi(s)
	if err != nil {
		return s
	}

	return i
}

func (c *cmd) Execute(todos *Todos) {
	switch {
	case c.Add != "":
		todos.Add(c.Add)
	case c.List:
		todos.Print()
	case c.Toggle != nil:
		err := todos.Toggle(c.Toggle)
		if err != nil {
			fmt.Println("Error toggling todo: ", err)
		}
	default:
		fmt.Println("Invalid Commands")
	}
}
