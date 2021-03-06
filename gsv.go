package nmea

import "fmt"

const (
	// TypeGSV type for GSV sentences
	TypeGSV = "GSV"
)

// GSV represents the GPS Satellites in view
// http://aprs.gids.nl/nmea/#glgsv
type GSV struct {
	BaseSentence
	TotalMessages   Int64     // Total number of messages of this type in this cycle
	MessageNumber   Int64     // Message number
	NumberSVsInView Int64     // Total number of SVs in view
	Info            []GSVInfo // visible satellite info (0-4 of these)
}

// GSVInfo represents information about a visible satellite
type GSVInfo struct {
	SVPRNNumber Int64 // SV PRN number, pseudo-random noise or gold code
	Elevation   Int64 // Elevation in degrees, 90 maximum
	Azimuth     Int64 // Azimuth, degrees from true north, 000 to 359
	SNR         Int64 // SNR, 00-99 dB (null when not tracking)
}

// newGSV constructor
func newGSV(s BaseSentence) (GSV, error) {
	p := NewParser(s)
	p.AssertType(TypeGSV)
	m := GSV{
		BaseSentence:    s,
		TotalMessages:   p.Int64(0, "total number of messages"),
		MessageNumber:   p.Int64(1, "message number"),
		NumberSVsInView: p.Int64(2, "number of SVs in view"),
	}
	for i := 0; i < 4; i++ {
		if 5*i+4 > len(m.Fields) {
			break
		}
		m.Info = append(m.Info, GSVInfo{
			SVPRNNumber: p.Int64(3+i*4, "SV prn number"),
			Elevation:   p.Int64(4+i*4, "elevation"),
			Azimuth:     p.Int64(5+i*4, "azimuth"),
			SNR:         p.Int64(6+i*4, "SNR"),
		})
	}
	return m, p.Err()
}

// GetNumberOfSatellites retrieves the number of satellites from the sentence
func (s GSV) GetNumberOfSatellites() (int64, error) {
	if v, err := s.NumberSVsInView.GetValue(); err == nil {
		return v, nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
