package coupler

import (
	"fmt"
	"strconv"
	"strings"
)

type FloatRange struct {
	Lower float64
	Upper float64
}

// Ator converts a string to a FloatRange
//
// Supported styles:
// - "10/20.5"
func Ator(s string) (*FloatRange, error) {
	fr := &FloatRange{}

	ss := strings.Split(s, "/")
	if len(ss) != 2 {
		return nil, fmt.Errorf("Invalid format detected in \"%s\"", s)
	}

	v1, err := strconv.ParseFloat(ss[0], 64)
	if err != nil {
		return nil, err
	}

	v2, err := strconv.ParseFloat(ss[1], 64)
	if err != nil {
		return nil, err
	}

	if v1 < v2 {
		fr.Lower = v1
		fr.Upper = v2
		return fr, nil
	}

	fr.Lower = v2
	fr.Upper = v1
	return fr, nil
}
