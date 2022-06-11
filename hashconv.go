/*
This package is a port of dustin/go-humanize but adjusted for hashrates
*/

package hashconv

import (
	"fmt"
	"math"
	"strconv"

	"github.com/dustin/go-humanize"
)

const (
	Hash  = 1
	KHash = Hash * 1000
	MHash = KHash * 1000
	GHash = MHash * 1000
	THash = GHash * 1000
	PHash = THash * 1000
	EHash = PHash * 1000
)

func Format(b int64) string {
	multiple := ""
	value := float64(b)

	switch {
	case b >= EHash:
		value /= EHash
		multiple = "EH"
	case b >= PHash:
		value /= PHash
		multiple = "PH"
	case b >= THash:
		value /= THash
		multiple = "TH"
	case b >= GHash:
		value /= GHash
		multiple = "GH"
	case b >= MHash:
		value /= MHash
		multiple = "MH"
	case b >= KHash:
		value /= KHash
		multiple = "kH"
	case b == 0:
		return "0 H"
	default:
		return strconv.FormatInt(b, 10) + " H"
	}

	// round to 3 decimal places
	value = math.Round(value*1000) / 1000

	return fmt.Sprintf("%s %s", humanize.Ftoa(value), multiple)
}
