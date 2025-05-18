package todo

import (
	"fmt"
	"os"
	"path/filepath"
)

const DEFAULT_VAULT_STORAGE string = ".todo"

func getRoot() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("Error %s", err)
	}

	return home, nil
}

func checkOrCreateVault(vault string) error {
	root, err := getRoot()

	if err != nil {
		return err
	}

	path := filepath.Join(root, DEFAULT_VAULT_STORAGE, fmt.Sprintf("%s.json", vault))

	if _, err := os.Stat(path); err != nil {
		if _, err := os.Create(path); err != nil {
			return err
		}

	}

	return nil
}

func createVaultStorage(storage string) error {
	root, err := getRoot()

	if err != nil {
		return err
	}

	path := filepath.Join(root, storage)

	err = os.MkdirAll(path, 0755)

	if err != nil {
		return err
	}

	return nil
}

func writeToVault(content []byte, vault string) error {
	root, err := getRoot()

	if err != nil {
		return err
	}

	path := filepath.Join(root, DEFAULT_VAULT_STORAGE, fmt.Sprintf("%s.json", vault))

	err = os.WriteFile(path, content, 0644)

	if err != nil {
		return err
	}

	return nil
}

func readFromVault(vault string) ([]byte, error) {
	root, err := getRoot()

	if err != nil {
		return nil, err
	}

	path := filepath.Join(root, DEFAULT_VAULT_STORAGE, fmt.Sprintf("%s.json", vault))

	content, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return content, nil
}

func Init(vault string) error {
	if err := createVaultStorage(DEFAULT_VAULT_STORAGE); err != nil {
		return err
	}

	if err := checkOrCreateVault(vault); err != nil {
		return err
	}

	return nil
}
