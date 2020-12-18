package snippet

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const SNIPPET_DIR = "$HOME/.config/snippet-man"

type Snippets struct {
	Snippets []Snippet `toml:"snippets"`
}

type Snippet struct {
	Description string `toml:"description"`
	Command     string `toml:"command"`
}

func (snippets *Snippets) Init() error {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "snippet-man")
	fullFilePath := filepath.Join(configDir, "snippets.toml")

	// Check for existence of snippet files
	_, err := os.Stat(fullFilePath)
	if err == nil {
		return fmt.Errorf("Initialization has already been done.")
	}

	// Create snippet file directory
	_, err = os.Stat(configDir)
	if err != nil {
		os.MkdirAll(configDir, 0755)
	}

	// Create snippet file
	_, err = os.Create(fullFilePath)
	if err != nil {
		return fmt.Errorf("Failed to create config file. %v", err)
	}

	return nil
}

func (snippets *Snippets) Load() error {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "snippet-man")
	fullFilePath := filepath.Join(configDir, "snippets.toml")

	_, err := toml.DecodeFile(fullFilePath, &snippets)
	if err != nil {
		return fmt.Errorf("Failed to load config file. %v", err)
	}

	for _, s := range snippets.Snippets {
		fmt.Printf("Name: %s\n", s.Command)
	}

	return nil
}
