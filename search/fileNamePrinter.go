package search

import (
	"fmt"
	"os"
)

type FileNamePrinter struct {
}

func (f FileNamePrinter) print(fileInfo os.FileInfo, matches []match) {
	if len(matches) > 0 {
		fmt.Printf("%v\n", fileInfo.Name())
	}
}
