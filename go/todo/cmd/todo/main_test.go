package main_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"
	"todo"
)

var toolName string = "test"

var vault string = "test"

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		toolName += ".exe"
	}

	if err := exec.Command("go", "build", "-o", toolName).Run(); err != nil {
		fmt.Println("Build failed:", err)
		os.Exit(1)
	}

	code := m.Run()

	time.Sleep(time.Second)
	err := os.Remove(toolName)

	if err != nil {
		fmt.Println("Cleanup failed")

		os.Exit(1)
	}

	os.Exit(code)
}

func TestCli(t *testing.T) {
	t.Run("AddNewTasks", func(t *testing.T) {
		path, err := os.Getwd()

		if err != nil {
			t.Fatal("Unable to get working directory")
		}

		if out, err := exec.Command(filepath.Join(path, toolName), "-v", vault, "Task").CombinedOutput(); err != nil {
			t.Fatalf("Command failed to execute: %s", out)
		}

		home, err := os.UserHomeDir()

		if err != nil {
			t.Error("Home Dir is not available")
		}

		l := todo.List{}

		content, err := os.ReadFile(filepath.Join(home, todo.DEFAULT_VAULT_STORAGE, fmt.Sprintf("%s.json", vault)))

		if err != nil {
			t.Errorf("Failed to parse json: %s", err)
		}

		json.Unmarshal(content, &l)

		if l[0].Title != "Task" {
			t.Error("Tasks not matched")
		}

	})

	t.Run("ListTasks", func(t *testing.T) {})
}
