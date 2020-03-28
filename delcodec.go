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
	"flag"
    "fmt"
    "os"
)

import (
//    "github.com/Causticity/sipp/simage"
//    "github.com/Causticity/sipp/sgrad"
//    "github.com/Causticity/sipp/shist"
//    "github.com/Causticity/sipp/sfft"
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
    
    flag.Usage()
}