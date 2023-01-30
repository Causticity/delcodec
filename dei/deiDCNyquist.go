package dei

import (
	//"fmt"
	"math"
	. "github.com/Causticity/sipp/simage"
)


// This file contains structures and code associated with the DC-Nyquist side
// channels of a DEI image.

// A DCvalue is the DC value of the frequency range of a row or column of a
// SippImage, computed as a fixed-point average over the row or column.
// We use int32s to accomodate both 8-bit and 16-bit images. Each value is a
// SippImage.Bpp().2 fixed-point number. Values are always positive, but int32
// rather than uint32 makes this parallel to the Nyquist case.
type DCvalue int32

// A NyguistValue is the value at the Nyquist frequency of a row or column of a
// SippImage, represented as a fixed-point sum over the row or column, using
// alternating +-1 modulation during the summation.
type NyquistValue int32

var fpScale = 4.0	// Equivalent to shift left by 2. We compute the shifted
					// average in floating-point, then round.

// DCcolumn returns the average values of each row of an image, as a slice of
// DCvalues, one per image row. 
func DCcolumn(src SippImage) (ret []DCvalue) {
	ret = make([]DCvalue, src.Bounds().Dy())
	var avg float64
	width := src.Bounds().Dx()
	for y, _ := range ret {
		avg = 0.0
		for x := 0; x < width; x++ {
			avg = avg + (src.Val(x, y) * fpScale)
		}
		ret[y] = DCvalue(math.Round(avg / float64(width)))
	}
	return
}

// DCrow returns the average values of each column of an image, as a slice of
// DCvalues, one per image column. 
func DCrow(src SippImage) (ret []DCvalue) {
	ret = make([]DCvalue, src.Bounds().Dx())
	var avg float64
	height := src.Bounds().Dy()
	for x, _ := range ret {
		avg = 0.0
		for y := 0; y < height; y++ {
			avg = avg + (src.Val(x, y) * fpScale)
		}
		ret[x] = DCvalue(math.Round(avg / float64(height)))
	}
	return
}