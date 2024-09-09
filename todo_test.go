package main

import "testing"

func Test_add(t *testing.T) {
	list := Todos{}
	list.add("walk")

	if len(list) != 1 {
		t.Errorf("Test Failed! Expected %v but instead got %v", true, false)
	}
}
