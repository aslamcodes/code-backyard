package todo

import (
	"encoding/json"
	"fmt"
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

func (l *List) Complete(i int) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("Invalid Id")
	}

	(*l)[i-1].Done = true
	(*l)[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Save(vault string) error {
	// This is code is hard for testing without any side effects, as it requires a file to created at a definite place
	js, err := json.Marshal(*l)

	if err != nil {
		return err
	}

	if err := writeToVault(js, vault); err != nil {
		return err
	}

	return nil
}

func (l *List) Load(vault string) error {
	js, err := readFromVault(vault)

	if err != nil {
		return err
	}

	json.Unmarshal(js, l)

	return nil
}
