// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standart output.
// includes exercise 10.1
package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// If no flag were supplied , then input will be converted to JPEG
var outformat = flag.String("format", "", "accepted values are : png, gif, jpeg")

func main() {

	flag.Parse()
	if err := convert(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "conversion: %v\n", err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer) error {
	img, format, err := image.Decode(in) // determine the format of the img
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", format)

	if format == "jpeg" && *outformat == "jpeg" {
		return errors.New("Input and desired output has same format(JPEG)")
	} else if *outformat == "gif" {

		return gif.Encode(out, img, nil)
	} else if *outformat == "png" {
		return png.Encode(out, img)
	} else {

		return jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
	}
}
