package search

import "os"

type ResultPrinter interface {
	print(fileInfo os.FileInfo, matches []match)
}
