package todo_test

import (
	// "fmt"
	// "os"
	// "path/filepath"
	"fmt"
	"os"
	"path/filepath"
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

	vault := "temp"

	err := l.Save(vault)

	if err != nil {
		t.Errorf("Error %s", err)
	}

	l2 := todo.List{}

	err = l2.Load(vault)

	if err != nil {
		t.Errorf("Error %s", err)
	}

	if l[0].Title != l2[0].Title {
		t.Errorf("List not matches")
	}

	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("Error getting the home dir")
	}

	os.Remove(filepath.Join(homeDir, ".todo", fmt.Sprintf("%s.json", vault)))
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	l.Add("t1")

	id := 1

	err := l.Complete(id)

	if err != nil {
		t.Error(err)
	}

	if !l[id-1].Done {
		t.Error("Task not completed")
	}
}
