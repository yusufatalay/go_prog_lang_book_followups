package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {

	var sha_256 = flag.Bool("s256", false, "Cypher given message with SHA256")
	var sha_384 = flag.Bool("s384", false, "Cypher given message with SHA384")
	var sha_512 = flag.Bool("s512", false, "Cypher given message with SHA512")

	flag.Parse()

	input := []byte(os.Args[2]) // NOTE: ignoring the presence of the arg[2] I assume that user will always provide an input
	var sha256Input [32]byte
	var sha384Input [48]byte
	var sha512Input [64]byte

	if *sha_256 {
		sha256Input = sha256.Sum256(input)
		fmt.Printf("%x\n", sha256Input)
	} else if *sha_384 {
		sha384Input = sha512.Sum384(input)
		fmt.Printf("%x\n", sha384Input)
	} else if *sha_512 {
		sha512Input = sha512.Sum512(input)
		fmt.Printf("%x\n", sha512Input)
	}

}
