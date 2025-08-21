package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// now := time.Now()
	// fmt.Println(fmt.Sprintf("execution at %s", now))

	inputDir := "/mnt"
	// outputDir := "/out"

	entries, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatalf("Impossibile leggere la directory di input: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirToBackup := filepath.Join(inputDir, entry.Name())
			fmt.Println(fmt.Sprintf("dir to backup: %s", dirToBackup))
		}
	}
}
