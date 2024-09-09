package main

import (
	"errors"
	"strings"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index > len(*todos) {
		err := errors.New("index out of bounds")
		return err
	}
	return nil
}

func (todos *Todos) delete(input interface{}) error {
	t := *todos

	switch v := input.(type) {
	case string:
		for i, value := range t {
			if strings.EqualFold(value.Title, v) {
				*todos = append(t[:i], t[i+1:]...)
				return nil
			}
		}
		return errors.New("todo not found")

	case int:
		err := t.validateIndex(v)
		if err != nil {
			return err
		}

		*todos = append(t[:v], t[v+1:]...)

		return nil

	default:
		return errors.New("invalid input type")
	}
}
