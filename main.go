package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var Reset = "\033[0m"
var Green = "\033[32m"
var Cyan = "\033[36m"

func byteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func printListing(entry os.FileInfo, depth int) {
	indent := strings.Repeat("|   ", depth)

	name := entry.Name()
	size := entry.Size()
	mode := entry.Mode()

	fmt.Printf("%s|-- %s [%s | %s] \n", indent, Green+name+Reset, Cyan+byteCountSI(size)+Reset, mode)
}

func printDirectory(path string, current string, depth int) {
	deepPath := filepath.Join(path, current)
	entries, err := ioutil.ReadDir(deepPath)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			printListing(entry, depth)
			printDirectory(deepPath, entry.Name(), depth+1)
		} else {
			printListing(entry, depth)
		}
	}
}

func main() {
	var path = "./"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	dir, err := filepath.Abs(filepath.Dir(path))

	if err != nil {
		fmt.Println("Error to get abs path: ", err)
		dir = path
	}

	fmt.Println("Welcome to FS Tree Viewer 0.1")
	fmt.Println("Starting to scan your FS from path: ", dir)

	printDirectory(path, "", 0)
}
