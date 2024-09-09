package main

import "testing"

func Test_add(t *testing.T) {
	list := Todos{}
	list.add("walk")

	if len(list) != 1 {
		t.Errorf("Test Failed! Expected %v but instead got %v", true, false)
	}
}

func Test_delete(t *testing.T) {
	list := Todos{}
	list.add("walk")
	list.add("talk")
	list.delete("walk")

	if len(list) != 1 {
		t.Errorf("Test Failed! Expected %v but instead got %v", true, false)
	}

	list.delete("talk")

	if len(list) != 0 {
		t.Errorf("Test Failed! Expected %v but instead got %v", true, false)
	}
}

func Test_case_sensitive_delete(t *testing.T) {
	list := Todos{}
	list.add("WALK")
	list.add("TALK")
	list.delete("walk")

	if len(list) != 1 {
		t.Errorf("Test Failed! Expected %v but instead got %v", true, false)
	}
}

func Test_delete_error(t *testing.T) {
	list := Todos{}
	list.add("WALK")

	not_err := list.delete("walk")

	if not_err == nil{
		t.Errorf("Test Failed!")
	}

	err := list.delete("talk")

	if err != nil{
		t.Errorf("Test Failed!")
	}
}
func Test_add_and_delete(t *testing.T) {
	list := Todos{}
	list.add("go buy some milk")
	list.add("take some photos at the park")
	list.add("return home")

	if len(list) != 3{
		t.Errorf("Test Failed!")
	}

	list.delete("Take SOME photos at the park")

	if len(list) != 2{
		t.Errorf("Test Failed!")
	}
}
