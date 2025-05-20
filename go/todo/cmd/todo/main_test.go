package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var toolName string = "test"

var vault string = "test2"

func TestMain(m *testing.M) {
	storage := "testing_storage"

	os.Setenv("TODO_VAULT_STORAGE", storage)

	if runtime.GOOS == "windows" {
		toolName += ".exe"
	}

	if err := exec.Command("go", "build", "-o", toolName).Run(); err != nil {
		fmt.Println("Build failed:", err)
		os.Exit(1)
	}

	code := m.Run()

	err := os.Remove(toolName)

	if err != nil {
		fmt.Println("Cleanup failed")
		os.Exit(1)
	}

	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Println("Unable to get home dir")
		os.Exit(1)
	}

	err = os.RemoveAll(filepath.Join(home, storage))

	os.Exit(code)
}

func TestCli(t *testing.T) {
	task := "Task"
	t.Run("AddNewTasks", func(t *testing.T) {
		path, err := os.Getwd()

		if err != nil {
			t.Fatal("Unable to get working directory")
		}

		if out, err := exec.Command(filepath.Join(path, toolName), "-v", vault, task).CombinedOutput(); err != nil {
			t.Fatalf("Command failed to execute: %s", out)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		expected := task + "\n"

		path, err := os.Getwd()

		if err != nil {
			t.Fatal("Unable to get working directory")
		}

		out, err := exec.Command(filepath.Join(path, toolName), "-v", vault, "-l").CombinedOutput()

		if string(out) != expected {
			t.Errorf("Not matched %s", out)
		}

		if err != nil {
			t.Fatalf("Command failed to execute: %s", out)
		}
	})

	t.Run("ShowAllTasks", func(t *testing.T) {
		expected := task + "\n"

		path, err := os.Getwd()

		if err != nil {
			t.Fatal("Unable to get working directory")
		}

		out, err := exec.Command(filepath.Join(path, toolName), "-v", vault, "-c", "1", "-all").CombinedOutput()

		if err != nil {
			t.Fatalf("Command failed to execute: %s", out)
		}

		if string(out) != expected {
			t.Errorf("Not matched %s", out)
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {

		path, err := os.Getwd()

		if err != nil {
			t.Fatal("Unable to get working directory")
		}

		out, err := exec.Command(filepath.Join(path, toolName), "-v", vault, "-c", "1").CombinedOutput()

		if err != nil {
			t.Fatalf("Command failed to execute: %s", out)
		}
	})

}
