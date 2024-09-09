package main

import "time"

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos[] Todo

func (todos *Todos) add(title string){
	todo := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) delete(title string){
	t := *todos
	for i, value := range t{
		if value.Title == title{
			*todos = append(t[:i], t[i+1:]...)
		}
	}
}
