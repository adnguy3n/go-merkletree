package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

/*
 * Merkle Tree node.
 */
type node struct {
	hash  []byte
	left  *node
	right *node
}

/*
 * Takes file pathnames as command line arguments and computes their top hash.
 */
func main() {
	fmt.Println(os.Args[1:])

	nodes := openFile(os.Args[1:], nil)
	fmt.Println(nodes)
}

/*
 * Opens the files using the pathnames and hashes them.
 */
func openFile(args []string, nodes []node) []node {
	if len(args) == 0 {
		return nodes
	}

	file, err := os.Open(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileHash := sha1.New()
	if _, err := io.Copy(fileHash, file); err != nil {
		log.Fatal(err)
	}

	nodes = append(nodes, node{hash: fileHash.Sum(nil)})

	return openFile(args[1:], nodes)
}
