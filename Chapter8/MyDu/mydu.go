// Exercise 8.9
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type rootinfo struct {
	rootname string
	size     int64
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes
func walkDir(root, dir string, wg *sync.WaitGroup, fileSizes chan<- rootinfo) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			walkDir(root, subdir, wg, fileSizes)
		} else {
			fileSizes <- rootinfo{root, entry.Size()}
		}
	}

}

// sm is a counting semaphore for limiting concurrency in dirents
var sm = make(chan struct{}, 20) // only run 20 goroutines at a time

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sm <- struct{}{}        //acquire token
	defer func() { <-sm }() //release the token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf((os.Stderr), "du1: %v\n", err)
	}
	return entries
}
func printRootUsage(roots []string, subdirCount, rootsize map[string]int64) {

	for _, root := range roots {
		fmt.Printf("%s :", root)
		printDiskUsage(subdirCount[root], rootsize[root])
	}

}
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f GB \n", nfiles, float64((nbytes)/1e9))
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args() // get the non-flag command line params
	if len(roots) == 0 {
		roots = []string{"."} // if no path is given the use current one
	}

	// Traverse each root of the file tree in parallel
	fileSizes := make(chan rootinfo)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()
	// Print the results periodically
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
	subdirCount := make(map[string]int64)
	rootsize := make(map[string]int64)
loop:
	for {
		select {
		case root, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += root.size
			rootsize[root.rootname] += root.size
			subdirCount[root.rootname]++
		case <-tick:
			printRootUsage(roots, subdirCount, rootsize)
		}
	}
	printDiskUsage(nfiles, nbytes) // print totals

}
