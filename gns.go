package nmea

import "fmt"

const (
	// TypeGNS type for GNS sentences
	TypeGNS = "GNS"
	// NoFixGNS Character
	NoFixGNS = "N"
	// AutonomousGNS Character
	AutonomousGNS = "A"
	// DifferentialGNS Character
	DifferentialGNS = "D"
	// PreciseGNS Character
	PreciseGNS = "P"
	// RealTimeKinematicGNS Character
	RealTimeKinematicGNS = "R"
	// FloatRTKGNS RealTime Kinematic Character
	FloatRTKGNS = "F"
	// EstimatedGNS Fix Character
	EstimatedGNS = "E"
	// ManualGNS Fix Character
	ManualGNS = "M"
	// SimulatorGNS Character
	SimulatorGNS = "S"
)

// GNS is standard GNSS sentance that combined multiple constellations
type GNS struct {
	BaseSentence
	Time       Time
	Latitude   Float64
	Longitude  Float64
	Mode       StringList
	SVs        Int64
	HDOP       Float64
	Altitude   Float64
	Separation Float64
	Age        Float64
	Station    Int64
}

// newGNS Constructor
func newGNS(s BaseSentence) (GNS, error) {
	p := NewParser(s)
	p.AssertType(TypeGNS)
	m := GNS{
		BaseSentence: s,
		Time:         p.Time(0, "time"),
		Latitude:     p.LatLong(1, 2, "latitude"),
		Longitude:    p.LatLong(3, 4, "longitude"),
		Mode:         p.EnumChars(5, "mode", NoFixGNS, AutonomousGNS, DifferentialGNS, PreciseGNS, RealTimeKinematicGNS, FloatRTKGNS, EstimatedGNS, ManualGNS, SimulatorGNS),
		SVs:          p.Int64(6, "SVs"),
		HDOP:         p.Float64(7, "HDOP"),
		Altitude:     p.Float64(8, "altitude"),
		Separation:   p.Float64(9, "separation"),
		Age:          p.Float64(10, "age"),
		Station:      p.Int64(11, "station"),
	}
	return m, p.Err()
}

// GetPosition3D retrieves the 3D position from the sentence
func (s GNS) GetPosition3D() (float64, float64, float64, error) {
	validModi := map[string]interface{}{
		AutonomousGNS:        nil,
		DifferentialGNS:      nil,
		PreciseGNS:           nil,
		RealTimeKinematicGNS: nil,
		FloatRTKGNS:          nil,
		EstimatedGNS:         nil,
		ManualGNS:            nil,
		SimulatorGNS:         nil,
	}
	if s.Mode.Valid {
		for _, m := range s.Mode.Values {
			if _, ok := validModi[m.Value]; ok && m.Valid {
				if vLat, err := s.Latitude.GetValue(); err == nil {
					if vLon, err := s.Longitude.GetValue(); err == nil {
						if vAlt, err := s.Altitude.GetValue(); err == nil {
							return vLat, vLon, vAlt, nil
						}
					}
				}
			}
		}
	}
	return 0, 0, 0, fmt.Errorf("value is unavailable")
}
