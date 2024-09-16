package main

import (
	"testing"
)

func Test_add(t *testing.T) {
	list := Todos{}
	list.Add("walk")

	if len(list) != 1 {
		t.Errorf("Test Failed! Expected %v but instead got %v", true, false)
	}
}

func Test_delete(t *testing.T) {
	list := Todos{}
	list.Add("walk")
	list.Add("talk")
	list.Delete("walk")

	if len(list) != 1 {
		t.Errorf("Test Failed! Expected %v but instead got %v", true, false)
	}

	list.Delete("talk")

	if len(list) != 0 {
		t.Errorf("Test Failed!")
	}
}

func Test_case_sensitive_delete(t *testing.T) {
	list := Todos{}
	list.Add("WALK")
	list.Add("TALK")
	list.Delete("walk")

	if len(list) != 1 {
		t.Errorf("Test Failed!")
	}
}

func Test_delete_error(t *testing.T) {
	list := Todos{}
	list.Add("WALK")

	not_err := list.Delete("walk")

	if not_err != nil {
		t.Errorf("Test Failed!")
	}

	err := list.Delete("talk")

	if err == nil {
		t.Errorf("Test Failed!")
	}
}
func Test_add_and_delete(t *testing.T) {
	list := Todos{}
	list.Add("go buy some milk")
	list.Add("take some photos at the park")
	list.Add("return home")

	if len(list) != 3 {
		t.Errorf("Test Failed!")
	}

	list.Delete("Take SOME photos at the park")

	if len(list) != 2 {
		t.Errorf("Test Failed!")
	}
}

func Test_toggle_string(t *testing.T) {
	list := Todos{}
	list.Add("WALK")

	list.Toggle("walk")

	if list[0].Completed != true && list[0].CompletedAt == nil {
		t.Errorf("Test Failed!")
	}
}

func Test_toggle_index(t *testing.T) {
	list := Todos{}
	list.Add("WALK")

	list.Toggle(0)

	if list[0].Completed != true && list[0].CompletedAt == nil {
		t.Errorf("Test Failed!")
	}
}

func Test_edit_string(t *testing.T) {
	list := Todos{}
	list.Add("WALK")

	list.Edit("walk", "talk")

	if list[0].Title != "talk" {
		t.Errorf("Test Failed!")
	}
}
func Test_edit_index(t *testing.T) {
	list := Todos{}
	list.Add("WALK")

	list.Edit(0, "talk")

	if list[0].Title != "talk" {
		t.Errorf("Test Failed!")
	}
}
