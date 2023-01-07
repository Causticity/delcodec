// Copyright Raul Vera 2022


package dei

import (
    "fmt"
    _ "sipp/sgrad"
)

type DEIstart struct {
    // Identifer sequence in the first three bytes of the file
    const ID string = `DEI`
    // Major version of the format this file conforms to. Major version changes
    // indicate a format change that is incompatible with previous versions, i.e.
    // software that can read a 1.x version file is expected to be able to read
    // any 1.y version file without error, though perhaps functionality specific
    // to the difference betwen x and y might be unavailable.
    MajorVersion byte
    // Minor version of the format this file conforms to. See above for the
    // distinction between major and minor version numbers.
    MinorVersion byte
    // The length, in bytes, of the metadata block immediately following this
    // header. The metadata may contain any arbitrary set of bytes used by a
    // specific implementation. If there is no such block the value is 0.
    MetadataSize uint32
}

type DEIimage struct {
    // Width of the image, in pixels.
    Width uint32
    // Height of the imnage, in pixels.
    Height uint32
    // Number of channels in the image.
    NumChannels uint8
}

type DEIchannel struct {
    // Bit depth of this channel
    BitDepth uint8
    // kernel for the x quincunx lattice for this channel
    Xkernel SippGradKernel
    // kernel for the y quincunx lattice for this channel
    Ykernel SippGradKernel
}