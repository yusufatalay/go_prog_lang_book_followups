package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}

}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf((os.Stderr), "du1: %v\n", err)
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f GB \n", nfiles, float64((nbytes)/1e9))
}

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args() // get the non-flag command line params
	if len(roots) == 0 {
		roots = []string{"."} // if no path is given the use current one
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	// Print the results.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++ // increment the number of files
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}
