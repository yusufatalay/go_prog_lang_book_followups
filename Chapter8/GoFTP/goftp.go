package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	// created an input loop
	for input.Scan() {
		line := strings.Split(input.Text(), " ")
		command := line[0]
		args := line[1:]
		switch command {
		case "pwd":
			path, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
				continue
			}
			c.Write([]byte(path + "\n"))
		case "ls":
			// read the current dir
			files, err := os.ReadDir(".")
			if err != nil {
				log.Fatal(err)
				continue
			}
			for _, file := range files {
				fileinfo, _ := file.Info()
				filesize := fileinfo.Size()
				isdir := "file"
				if file.IsDir() {
					isdir = "directory"
				}
				line := fmt.Sprintf("%s %d  %s\n", file.Name(), filesize, isdir)
				c.Write([]byte(line))
			}

		case "cd":
			err := os.Chdir(args[0])
			if err != nil {
				log.Fatal(err)
				continue
			}
		case "get":
			for _, file := range args {
				data, err := os.Open(file)

				if err != nil {
					log.Fatal(err)
					continue
				}
				mustCopy(c, data)
			}
		case "close":
			return
		}
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // connection aborted
			continue
		}
		go handleConn(conn)
	}
}
