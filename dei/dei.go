// Copyright Raul Vera 2022


package dei

import (
    //"fmt"
    . "github.com/Causticity/sipp/sgrad"
)

// The following structures are in-memory representations of the components of
// a Delentropy Encoded Image (DEI) file. A DEI file consists of a header
// followed by an implementation-specific metadata block, followed by
// channel data. Channels are not interleaved.

// All multi-byte values, including floats, are represented in the file in
// network byte order, i.e. big endian, MSB first. (Complex numbers are pairs of
// floats, real first, then imaginary.) Multi-byte values in these structs are
// in the byte order native to the machine, so readers and writers must take
// care to swab as necessary. See the encoding/binary standard library package.

// DEIheader is an in-memory representation of the header block at the beginning
// of a DEI file.
type DEIheader struct {
    // Identifer sequence "DEI" in the first three bytes of the file
    ID [3]uint8
    // Major version of the format this file conforms to. Major version changes
    // indicate a format change that is incompatible with previous versions, i.e.
    // software that can read a 1.x version file is expected to be able to read
    // any 1.y version file without error, though perhaps functionality specific
    // to the difference betwen x and y might be unavailable.
    MajorVersion uint8
    // Minor version of the format this file conforms to. See above for the
    // distinction between major and minor version numbers.
    MinorVersion uint8
    // Width of the image, in pixels.
    Width uint32
    // Height of the image, in pixels.
    Height uint32
    // Number of channels in the image.
    NumChannels uint8
    // The length, in bytes, of the metadata block immediately following this
    // header. The metadata may contain any arbitrary block of bytes used by a
    // specific implementation. If there is no such block the value is 0.
    MetadataSize uint32
}

// A DEI file consists of one of the above header structures, followed by
// DEIHeader.NumChannels instances of the following channel structure.

// DEIchannel is an in-memory structure corresponding to one channel of an image
// encoded in a DEI file.
type DEIchannel struct {
    // Bit depth of this channel
    BitDepth uint8
    // kernel for the x quincunx lattice for this channel.
    Xkernel SippGradKernel
    // kernel for the y quincunx lattice for this channel
    Ykernel SippGradKernel
    // Huffman dictionary for DC column
    // DC column (line of average column pixel values), Huffman coded. Must be
    	// accurate to .25 bit
    // Huffman dictionary for DC row
    // DC row (line of average row pixel values), Huffman coded. Must be
    	// accurate to .25 bit
    // Huffman dictionary for Nyquist column
    // Nyquist column (average column but summed with +-1 alternation), Huffman
    	// coded. Must be accurate to .25 bit
    // Huffman dictionary for Nyquist row
    // Nyquist row (average row but summed with +-1 alternation), Huffman coded.
    	// Must be accurate to .25 bit
    // Huffman dictionary for x lattice
    // Huffman-encoded data for x lattice
    // Huffman dictionary for y lattice
    // Huffman-encoded data for y lattice
}