package nmea

// Latitude / longitude representation.

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	// Degrees value
	Degrees = '\u00B0'
	// Minutes value
	Minutes = '\''
	// Seconds value
	Seconds = '"'
	// Point value
	Point = '.'
	// North value
	North = "N"
	// South value
	South = "S"
	// East value
	East = "E"
	// West value
	West = "W"
)

// ParseLatLong parses the supplied string into the LatLong.
//
// Supported formats are:
// - DMS (e.g. 33° 23' 22")
// - Decimal (e.g. 33.23454)
// - GPS (e.g 15113.4322S)
//
func ParseLatLong(s string) (Float64, error) {
	if v, err := ParseDMS(s); err == nil {
		return v, nil
	}
	if v, err := ParseGPS(s); err == nil {
		return v, nil
	}
	if v, err := ParseDecimal(s); err == nil {
		return v, nil
	}
	return Float64{}, fmt.Errorf("cannot parse [%s], unknown format", s)

}

// ParseGPS parses a GPS/NMEA coordinate.
// e.g 15113.4322S
func ParseGPS(s string) (Float64, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return Float64{}, fmt.Errorf("invalid format: %s", s)
	}
	dir := parts[1]
	value, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return Float64{}, fmt.Errorf("parse error: %s", err.Error())
	}

	degrees := math.Floor(value / 100)
	minutes := value - (degrees * 100)
	value = degrees + minutes/60

	if dir == North || dir == East {
		return Float64{Valid: true, Value: value}, nil
	}
	if dir == South || dir == West {
		return Float64{Valid: true, Value: 0 - value}, nil
	}
	return Float64{}, fmt.Errorf("invalid direction [%s]", dir)
}

// FormatGPS formats a GPS/NMEA coordinate
func FormatGPS(l Float64) string {
	padding := ""
	degrees := math.Floor(math.Abs(l.Value))
	fraction := (math.Abs(l.Value) - degrees) * 60
	if fraction < 10 {
		padding = "0"
	}
	return fmt.Sprintf("%d%s%.4f", int(degrees), padding, fraction)
}

// ParseDecimal parses a decimal format coordinate.
// e.g: 151.196019
func ParseDecimal(s string) (Float64, error) {
	// Make sure it parses as a float.
	l, err := strconv.ParseFloat(s, 64)
	if err != nil || s[0] != '-' && len(strings.Split(s, ".")[0]) > 3 {
		return Float64{}, errors.New("parse error (not decimal coordinate)")
	}
	return Float64{Valid: true, Value: l}, nil
}

// ParseDMS parses a coordinate in degrees, minutes, seconds.
// - e.g. 33° 23' 22"
func ParseDMS(s string) (Float64, error) {
	degrees := 0
	minutes := 0
	seconds := 0.0
	// Whether a number has finished parsing (i.e whitespace after it)
	endNumber := false
	// Temporary parse buffer.
	tmpBytes := []byte{}
	var err error

	for i, r := range s {
		switch {
		case unicode.IsNumber(r) || r == '.':
			if !endNumber {
				tmpBytes = append(tmpBytes, s[i])
			} else {
				return Float64{}, errors.New("parse error (no delimiter)")
			}
		case unicode.IsSpace(r) && len(tmpBytes) > 0:
			endNumber = true
		case r == Degrees:
			if degrees, err = strconv.Atoi(string(tmpBytes)); err != nil {
				return Float64{}, errors.New("parse error (degrees)")
			}
			tmpBytes = tmpBytes[:0]
			endNumber = false
		case s[i] == Minutes:
			if minutes, err = strconv.Atoi(string(tmpBytes)); err != nil {
				return Float64{}, errors.New("parse error (minutes)")
			}
			tmpBytes = tmpBytes[:0]
			endNumber = false
		case s[i] == Seconds:
			if seconds, err = strconv.ParseFloat(string(tmpBytes), 64); err != nil {
				return Float64{}, errors.New("parse error (seconds)")
			}
			tmpBytes = tmpBytes[:0]
			endNumber = false
		case unicode.IsSpace(r) && len(tmpBytes) == 0:
			continue
		default:
			return Float64{}, fmt.Errorf("parse error (unknown symbol [%d])", s[i])
		}
	}
	if len(tmpBytes) > 0 {
		return Float64{}, fmt.Errorf("parse error (trailing data [%s])", string(tmpBytes))
	}
	val := float64(degrees) + (float64(minutes) / 60.0) + (float64(seconds) / 60.0 / 60.0)
	return Float64{Valid: true, Value: val}, nil
}

// FormatDMS returns the degrees, minutes, seconds format for the given LatLong.
func FormatDMS(l Float64) string {
	val := math.Abs(l.Value)
	degrees := int(math.Floor(val))
	minutes := int(math.Floor(60 * (val - float64(degrees))))
	seconds := 3600 * (val - float64(degrees) - (float64(minutes) / 60))
	return fmt.Sprintf("%d\u00B0 %d' %f\"", degrees, minutes, seconds)
}

// Time type
type Time struct {
	Valid       bool
	Hour        int
	Minute      int
	Second      int
	Millisecond int
}

// String representation of Time
func (t Time) String() string {
	seconds := float64(t.Second) + float64(t.Millisecond)/1000
	return fmt.Sprintf("%02d:%02d:%07.4f", t.Hour, t.Minute, seconds)
}

// timeRe is used to validate time strings
var timeRe = regexp.MustCompile(`^\d{6}(\.\d*)?$`)

// ParseTime parses wall clock time.
// e.g. hhmmss.ssss
// An empty time string will result in an invalid time.
func ParseTime(s string) (Time, error) {
	if s == "" {
		return Time{}, nil
	}
	if !timeRe.MatchString(s) {
		return Time{}, fmt.Errorf("parse time: expected hhmmss.ss format, got '%s'", s)
	}
	hour, _ := strconv.Atoi(s[:2])
	minute, _ := strconv.Atoi(s[2:4])
	second, _ := strconv.ParseFloat(s[4:], 64)
	whole, frac := math.Modf(second)
	return Time{true, hour, minute, int(whole), int(math.Round(frac * 1000))}, nil
}

// Date type
type Date struct {
	Valid bool
	DD    int
	MM    int
	YY    int
}

// String representation of date
func (d Date) String() string {
	return fmt.Sprintf("%02d/%02d/%02d", d.DD, d.MM, d.YY)
}

// ParseDate field ddmmyy format
func ParseDate(ddmmyy string) (Date, error) {
	if ddmmyy == "" {
		return Date{}, nil
	}
	if len(ddmmyy) != 6 {
		return Date{}, fmt.Errorf("parse date: exptected ddmmyy format, got '%s'", ddmmyy)
	}
	dd, err := strconv.Atoi(ddmmyy[0:2])
	if err != nil {
		return Date{}, errors.New(ddmmyy)
	}
	mm, err := strconv.Atoi(ddmmyy[2:4])
	if err != nil {
		return Date{}, errors.New(ddmmyy)
	}
	yy, err := strconv.Atoi(ddmmyy[4:6])
	if err != nil {
		return Date{}, errors.New(ddmmyy)
	}
	return Date{true, dd, mm, yy}, nil
}

// LatDir returns the latitude direction symbol
func LatDir(l float64) string {
	if l < 0.0 {
		return South
	}
	return North
}

// LonDir returns the longitude direction symbol
func LonDir(l float64) string {
	if l < 0.0 {
		return East
	}
	return West
}

type Float64 struct {
	Valid bool
	Value float64
}

func (v Float64) GetValue() (float64, error) {
	if v.Valid {
		return v.Value, nil
	}
	return 0, fmt.Errorf("the value is nil")
}

type Int64 struct {
	Valid bool
	Value int64
}

func (v Int64) GetValue() (int64, error) {
	if v.Valid {
		return v.Value, nil
	}
	return 0, fmt.Errorf("the value is nil")
}
