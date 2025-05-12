package todo_test

import (
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	l.Add("Do Go!")

	if l[0].Title != "Do Go!" {
		t.Errorf("Title is not matching")
	}
}

func TestRemove(t *testing.T) {
	l := todo.List{}

	l.Add("Do Go")

	l.Remove(1)

	if len(l) != 0 {
		t.Errorf("Item is not removed")
	}
}

func TestSave(t *testing.T) {
	l := todo.List{}

	l.Add("t1")
	l.Add("t2")
	l.Add("t3")
	l.Add("t4")

	vault := "home"

	l.Save(vault)
}

func TestLoad(t *testing.T) {
}
