package todo

import (
	"encoding/json"
	"errors"
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
	js, err := json.Marshal(*l)

	if err != nil {
		return err
	}

	home, _ := os.UserHomeDir()

	path := filepath.Join(home, ".todo", fmt.Sprintf("%s.json", vault))

	os.WriteFile(path, js, 0644)

	return nil
}

func (l *List) Load(vault string) error {
	home, _ := os.UserHomeDir()

	path := filepath.Join(home, ".todo", fmt.Sprintf("%s.json", vault))

	js, err := os.ReadFile(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	json.Unmarshal(js, l)

	return nil
}
