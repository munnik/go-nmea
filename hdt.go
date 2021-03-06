package nmea

import (
	"fmt"

	"github.com/martinlindhe/unit"
)

const (
	// TypeHDT type for HDT sentences
	TypeHDT = "HDT"
)

// HDT is the Actual vessel heading in degrees True.
// http://aprs.gids.nl/nmea/#hdt
type HDT struct {
	BaseSentence
	Heading Float64 // Heading in degrees
	True    bool    // Heading is relative to true north
}

// newHDT constructor
func newHDT(s BaseSentence) (HDT, error) {
	p := NewParser(s)
	p.AssertType(TypeHDT)
	m := HDT{
		BaseSentence: s,
		Heading:      p.Float64(0, "heading"),
		True:         p.EnumString(1, "true", "T").Value == "T",
	}
	return m, p.Err()
}

// GetTrueHeading retrieves the true heading from the sentence
func (s HDT) GetTrueHeading() (float64, error) {
	if v, err := s.Heading.GetValue(); err == nil && s.True {
		return (unit.Angle(v) * unit.Degree).Radians(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
