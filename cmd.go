
package main

import (
	"path/filepath"
)

func main() {


}

func openFile(path string) (string, error) {
	return filepath.Abs(path)
}