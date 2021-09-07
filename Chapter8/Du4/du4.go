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

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes
func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			walkDir(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}

}

// sm is a counting semaphore for limiting concurrency in dirents
var sm = make(chan struct{}, 20) // only run 20 goroutines at a time

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case sm <- struct{}{}: //acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sm }() //release the token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf((os.Stderr), "du1: %v\n", err)
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f GB \n", nfiles, float64((nbytes)/1e9))
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args() // get the non-flag command line params
	if len(roots) == 0 {
		roots = []string{"."} // if no path is given the use current one
	}

	go func() {
		os.Stdin.Read(make([]byte, 1)) //	read a single byte
		close(done)
	}()

	// Traverse each root of the file tree in parallel
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
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
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish
			for range fileSizes {
				// Do nothing
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}

	}
	printDiskUsage(nfiles, nbytes) // print totals

}
