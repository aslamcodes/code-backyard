package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"time"
)

type item struct {
	Title       string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(title string) {
	i := item{
		Title:       title,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, i)
}

func (l *List) Remove(i int) error {

	if i <= 0 || i > len(*l) {
		return fmt.Errorf("Invalid Id")
	}

	*l = slices.Delete((*l), i-1, i)

	return nil
}

func (l *List) Save(vault string) error {
	// This is code is hard for testing without any side effects, as it requires a file to created at a definite place
	js, err := json.Marshal(*l)

	if err != nil {
		return err
	}

	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("Error %s", err)
	}

	path := filepath.Join(home, ".todo", fmt.Sprintf("%s.json", vault))

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, js, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (l *List) Load(vault string) error {
	home, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	path := filepath.Join(home, ".todo", fmt.Sprintf("%s.json", vault))

	js, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	json.Unmarshal(js, l)

	return nil
}
