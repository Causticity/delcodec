// Copyright Raul Vera 2020

// This program implements an encoder/decoder pair for encoding and decoding
// grescale images using the delentropy method described in Kieran Larkin's
// paper: https://arxiv.org/pdf/1609.01117v1.pdf
// Encoded images have a .dei suffix, an acronym for "[d]elentropy [e]ncoded [i]mage".
// TODO: perhaps there should be two separate programs, one for encoding and
// one for decoding? For now it just uses the suffix of the input file to 
// decide what to do, and outputs the converted file with the corresponding
// suffix. Decoded files are written only as PNG to start with.

package main

import (
	"errors"
	"flag"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

import (
    simage "github.com/Causticity/sipp/simage"
//   . "github.com/Causticity/sipp/sgrad"
//   . "github.com/Causticity/sipp/shist"
//   . "github.com/Causticity/sipp/sfft"
)

func main() {

	flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
        fmt.Println()
        fmt.Println("This program uses FFTW (http://www.fftw.org/), licensed under the GPL.")
        fmt.Println("Consequently this program is also licensed under the GPL v3 (http://www.gnu.org/licenses/gpl.html)")
        fmt.Println("Source code for this program may be found at (https://github.com/Causticity/delcodec)")
    }

	var in = flag.String("in", "", "Input image filename; must be either grayscale png or dei encoded")

	flag.Parse()
	
	if *in == "" {
		flag.Usage()
		os.Exit(0)
	}

	out, process, err := validateFile(*in)

    if err != nil {
        fmt.Println("Error validating file <", *in, ">: ", err)
        os.Exit (1)
    }

    // Check that the input file exists
	if _, err := os.Stat(*in); os.IsNotExist(err) {
	    fmt.Println("Input file does not exist: " + *in)
        os.Exit (1)
	}

    err = process(*in, out)
    if err != nil {
        fmt.Println("Error processing file <", *in, ">: ", err)
        os.Exit (1)
    }
}

// Verify that the input file name is valid (has either a .png or .dei suffix).
// Returns the name of the output file, the function to use to process (either
// encode or decode), and an error code or nil.
func validateFile(in string) (string, (func(string, string) error), error) {
    inext := filepath.Ext(in)
    var proc procFunc
    var out string
    var err error
    if strings.EqualFold(inext, ".png") {
        proc = encode
        out = strings.TrimSuffix(in, inext) + ".dei"
        err = nil
    } else if strings.EqualFold(inext, ".dei") {
        proc = decode
        out = strings.TrimSuffix(in, inext) + ".png"
        err = nil
    } else {
        out = ""
        proc = nil
        if inext == "" {
        	err = errors.New("Input file has no extension")
        } else {
        	err = errors.New("Invalid input file extension: " + inext[1:])
        }
    }
    return out, proc, err
}

type procFunc func(string, string) error

func encode(in string, out string) error {
    fmt.Println("Encoding " + in + " to " + out)
    fmt.Println("unimplemented")
    image, err := simage.Read(in)
    if err != nil {
    	return err
    }
    dst := encodeImage(image)
	writer, err := os.Create(out)
	if err != nil {
		return err
	}
	_, err = writer.Write(dst)
	if err != nil {
		return err
	}
    return nil
}

// TODO: return an in-memory DEI object?
func encodeImage(image simage.SippImage) []byte {
    // Set up quincunx lattice (just use one)
    // Take the fft and extract the DC and Nyquist rows and columns
    // compute gradient image on the lattice only
    // compute deldensity of the gradient
    // Huffman gradient based on deldensity
    // Huffman of DC and Nyquist based on their own line statistics
    return nil
}

func decode(in string, out string) error {
    fmt.Println("Decoding " + in + " to " + out)
    fmt.Println("unimplemented")
    // Read dei file into memory
    // Write image
    return nil
}

func decodeImage(dei []byte) simage.SippImage {
    // Extract and Huffman decode DC and Nyquist lines
    // Huffman decode gradient
    // Create gradient image with interspersed zeroes
    // In Fourier domain, remove aliases, undo gradient, add DC and Nyquist
    // Inverse FFT to get image
    return nil
}