package main

import (
	"fmt"
	"log"
	"os"
)

/*
 * Mekle Tree node.
 */
/*
type node struct {
	hash  []byte
	left  *node
	right *node
}*/

/*
 * Takes file pathnames as command line arguments and computes their top hash.
 */
func main() {
	fmt.Println(os.Args[1:])

	files := openFile(os.Args[1:], nil)
	fmt.Println(files)
}

/*
 * Opens the files using the pathnames.
 */
func openFile(args []string, files []*os.File) []*os.File {
	if len(args) == 0 {
		return files
	}

	file, err := os.Open(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}

	files = append(files, file)

	file.Close()

	return openFile(args[1:], files)
}
