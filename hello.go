package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run hello.go [path]")
		return
	}
	rootDir := os.Args[1]

	targetFolder := "node_modules"

	err := findFolder(rootDir, targetFolder)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func findFolder(dir string, target string) error {

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == target {
			cmd := exec.Command("rm", "-rf", path)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error removing folder:", path, err)
			}	
			fmt.Println("Successfully removed folder:", path)
			return filepath.SkipDir

		}
		return nil
	})
}
