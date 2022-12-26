package search

import (
	"fmt"
	"os"
)

type FileAndContentPrinter struct {
}

func (f FileAndContentPrinter) print(fileInfo os.FileInfo, matches []match) {
	if len(matches) > 0 {
		fmt.Printf("%v\n", fileInfo.Name())
		for _, match := range matches {
			fmt.Printf("%v -> %v\n", match.lineNumber, match.text)
		}
		fmt.Println()
	}
}
