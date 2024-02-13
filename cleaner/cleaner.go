package cleaner

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	name string
	d    bool
	f    bool
}

func NewConfig(name string, d bool, f bool) (*Config, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if d && f {
		return nil, fmt.Errorf("only one of -d or -f can be used")
	}

	return &Config{name, d, f}, nil
}

func FindAndDeleteAll(cfg *Config) error {
	found, err := find(cfg)
	if err != nil {
		return err
	}

	return delete(cfg, found)
}

func find(cfg *Config) ([]string, error) {
	var found []string

	rootDir := "."
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() != cfg.name {
			return nil
		}
		if cfg.d && !info.IsDir() {
			return nil
		} else if cfg.f && info.IsDir() {
			return nil
		}

		found = append(found, path)
		fmt.Println(path)

		if info.IsDir() {
			return filepath.SkipDir
		}

		return nil
	})

	return found, err
}

func delete(cfg *Config, found []string) error {
	if len(found) == 0 {
		fmt.Println("not found.")
		return nil
	}

	var confirmation string
	fmt.Printf("Only if you want to delete all (%d)\nPress 'y', otherwise press any key:\n", len(found))
	fmt.Scan(&confirmation)

	if confirmation != "y" {
		fmt.Println("canceled.")
		return nil
	}

	for _, p := range found {
		err := os.RemoveAll(p)
		if err != nil {
			return err
		}
	}

	var msg string
	if cfg.d {
		msg = "directories"
	} else if cfg.f {
		msg = "files"
	} else {
		msg = "directories and files"
	}

	fmt.Printf("all %s deleted successfully.\n", msg)

	return nil
}
