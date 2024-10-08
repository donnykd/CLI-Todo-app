package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
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

func (todos *Todos) Delete(input interface{}) error {
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
		return errors.New("invalid input")
	}
}

func (todos *Todos) Toggle(input interface{}) error {
	t := *todos

	switch v := input.(type) {
	case string:
		for i, value := range t {
			if strings.EqualFold(value.Title, v) {
				t[i].Completed = !value.Completed

				if !value.Completed {
					completedTime := time.Now()
					t[i].CompletedAt = &completedTime
				}

				return nil
			}
		}
		return errors.New("todo not found")

	case int:
		err := t.validateIndex(v)
		if err != nil {
			return err
		}

		if !t[v].Completed {
			completedTime := time.Now()
			t[v].CompletedAt = &completedTime
		}

		t[v].Completed = !t[v].Completed
		return nil

	default:
		return errors.New("invalid input")
	}
}

func (todos *Todos) Edit(input interface{}, title string) error {
	t := *todos

	switch v := input.(type) {
	case string:
		for i, value := range t {
			if strings.EqualFold(value.Title, v) {
				t[i].Title = title
				return nil
			}
		}
		return errors.New("todo not found")

	case int:
		err := t.validateIndex(v)
		if err != nil {
			return err
		}

		t[v].Title = title
		return nil

	default:
		return errors.New("invalid input")
	}
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for i, todo := range *todos {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(i), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
