package nmea

const (
	// TypeWPL type for WPL sentences
	TypeWPL = "WPL"
)

// WPL contains information about a waypoint location
type WPL struct {
	BaseSentence
	Latitude  Float64 // Latitude
	Longitude Float64 // Longitude
	Ident     String  // Ident of nth waypoint
}

// newWPL constructor
func newWPL(s BaseSentence) (WPL, error) {
	p := NewParser(s)
	p.AssertType(TypeWPL)
	return WPL{
		BaseSentence: s,
		Latitude:     p.LatLong(0, 1, "latitude"),
		Longitude:    p.LatLong(2, 3, "longitude"),
		Ident:        p.String(4, "ident of nth waypoint"),
	}, p.Err()
}
