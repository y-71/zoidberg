package zoidberg

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

const localNode string = "localhost:5001"

/**
* handle: handles error at different level
* should be moved to utilities module. 🔧
**/
func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
	/**
	* TODO: implement different error handling levels.
	**/
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %s", err)
	// 	os.Exit(1)
	// }
}

/**
 * readFile: reads a file 📖
 * TODO: should be moved to utils 🔧
 */
func readFile(file string) io.Reader {
	content, err := os.ReadFile(file)
	handle(err)
	return bytes.NewReader(content)
}

/**
* store: store a file to IPFS 🚀🌏
**/
func Deploy(file string) {
	data := readFile(file)
	sh := shell.NewShell(localNode)
	cid, err := sh.Add(data)
	handle(err)
	fmt.Printf("added %s", cid)
}
