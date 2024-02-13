package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	name := flag.String("n", "", "name to delete")
	d := flag.Bool("d", false, "only directory")
	f := flag.Bool("f", false, "only file")
	flag.Parse()

	if *name == "" {
		fmt.Println("name is required")
		flag.PrintDefaults()
		return
	}

	if *d && *f {
		fmt.Println("only one of -d or -f can be used")
		flag.PrintDefaults()
		return
	}

	var targets []string
	rootDir := "."
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() != *name {
			return nil
		}
		if *d && !info.IsDir() {
			return nil
		} else if *f && info.IsDir() {
			return nil
		}

		targets = append(targets, path)
		fmt.Println(path)

		if info.IsDir() {
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	targetName := *name
	if *d {
		targetName += " directories"
	} else if *f {
		targetName += " files"
	}

	fmt.Printf("Only if you want to delete all (%d)\nPress 'y', otherwise press any key:\n", len(targets))
	var confirmation string
	fmt.Scan(&confirmation)
	if confirmation != "y" {
		fmt.Println("canceled.")
		return
	}

	for _, p := range targets {
		err := os.RemoveAll(p)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("all %s deleted successfully.\n", targetName)
}
