package main

import (
	"crypto/sha1"
	"encoding/hex"
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
 * TODO: Print out entire tree.
 */
func main() {
	nodes := openFile(os.Args[1:], nil)
	topNode := buildTree(nodes)
	topHash := hex.EncodeToString(topNode.hash)
	fmt.Println(topHash)
}

/*
 * Opens the files using the pathnames and hashes them.
 */
func openFile(args []string, nodes []node) []node {
	if len(args) == 0 {
		return nodes
	}

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}

	fileHash := sha1.New()
	if _, err := io.Copy(fileHash, file); err != nil {
		log.Fatal(err)
	}

	nodes = append(nodes, node{hash: fileHash.Sum(nil)})

	return openFile(args[1:], nodes)
}

/*
 * Build the Merkle Tree from the bottom-up.
 */
func buildTree(nodes []node) node {
	if len(nodes) == 0 {
		log.Fatal("There should not be 0 nodes.")
	}

	if len(nodes) == 1 {
		return nodes[0]
	}

	var newLevel []node

	for i := 0; i < len(nodes); i += 2 {
		if i+1 < len(nodes) {
			combinedHash := append(nodes[i].hash, nodes[i+1].hash...)
			fileHash := sha1.New()
			fileHash.Write(combinedHash)
			newLevel = append(newLevel, node{hash: fileHash.Sum(nil), left: &nodes[i], right: &nodes[i+1]})
		} else {
			// Use duplicate of last node if there is an odd number.
			combinedHash := append(nodes[i].hash, nodes[i].hash...)
			fileHash := sha1.New()
			fileHash.Write(combinedHash)
			newLevel = append(newLevel, node{hash: fileHash.Sum(nil), left: &nodes[i]})
		}
	}

	//newLevel := buildNextLevel(nodes, nil)

	return buildTree(newLevel)
}

/*
 * Recursive version of creating the next level of nodes.
 * Written just to try it out.
 */
/*
func buildNextLevel(nodes, newLevel []node) []node {
	switch len(nodes) {
	case 0:
		return newLevel

	case 1:
		combinedHash := append(nodes[0].hash, nodes[0].hash...)
		fileHash := sha1.New()
		fileHash.Write(combinedHash)
		newLevel = append(newLevel, node{hash: fileHash.Sum(nil), left: &nodes[0]})
		return buildNextLevel(nil, newLevel)
	default:
		combinedHash := append(nodes[0].hash, nodes[1].hash...)
		fileHash := sha1.New()
		fileHash.Write(combinedHash)
		newLevel = append(newLevel, node{hash: fileHash.Sum(nil), left: &nodes[0], right: &nodes[1]})
	}

	return buildNextLevel(nodes[2:], newLevel)
}
*/
