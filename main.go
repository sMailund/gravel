package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	if strings.Compare(args[0], "-s") == 0 {
		search(args[1])
	}
}

func search(term string) {
	e := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		searchFile(path, term)
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
}

func searchFile(path string, term string) {
	f, err := os.Open(path)
	if err != nil {
		panic("")
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic("")
	}

	if fileInfo.IsDir() {
		return
	}

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, term) {
			println(text)
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
