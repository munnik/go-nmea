package nmea_test

import (
	"testing"

	. "github.com/munnik/go-nmea"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var gnstests = []struct {
	name string
	raw  string
	err  string
	msg  GNS
}{
	{
		name: "good sentence A",
		raw:  "$GNGNS,014035.00,4332.69262,S,17235.48549,E,RR,13,0.9,25.63,11.24,,*70",
		msg: GNS{
			Time:       Time{true, 1, 40, 35, 0},
			Latitude:   MustParseGPS("4332.69262 S"),
			Longitude:  MustParseGPS("17235.48549 E"),
			Mode:       []string{"R", "R"},
			SVs:        NewInt64(13),
			HDOP:       NewFloat64(0.9),
			Altitude:   NewFloat64(25.63),
			Separation: NewFloat64(11.24),
			Age:        Float64{},
			Station:    Int64{},
		},
	},
	{
		name: "good sentence B",
		raw:  "$GNGNS,094821.0,4849.931307,N,00216.053323,E,AA,14,0.6,161.5,48.0,,*6D",
		msg: GNS{
			Time:       Time{true, 9, 48, 21, 0},
			Latitude:   MustParseGPS("4849.931307 N"),
			Longitude:  MustParseGPS("00216.053323 E"),
			Mode:       []string{"A", "A"},
			SVs:        NewInt64(14),
			HDOP:       NewFloat64(0.6),
			Altitude:   NewFloat64(161.5),
			Separation: NewFloat64(48.0),
			Age:        Float64{},
			Station:    Int64{},
		},
	},
	{
		name: "good sentence B",
		raw:  "$GNGNS,094821.0,4849.931307,N,00216.053323,E,AAN,14,0.6,161.5,48.0,,*23",
		msg: GNS{
			Time:       Time{true, 9, 48, 21, 0},
			Latitude:   MustParseGPS("4849.931307 N"),
			Longitude:  MustParseGPS("00216.053323 E"),
			Mode:       []string{"A", "A", "N"},
			SVs:        NewInt64(14),
			HDOP:       NewFloat64(0.6),
			Altitude:   NewFloat64(161.5),
			Separation: NewFloat64(48.0),
			Age:        Float64{},
			Station:    Int64{},
		},
	},
	{
		name: "bad sentence",
		raw:  "$GNGNS,094821.0,4849.931307,N,00216.053323,E,AAX,14,0.6,161.5,48.0,,*35",
		err:  "nmea: GNGNS invalid mode: AAX",
	},
}

func TestGNS(t *testing.T) {
	for _, tt := range gnstests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				gns := m.(GNS)
				gns.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, gns)
			}
		})
	}
}

var _ = Describe("GNS", func() {
	var (
		parsed GNS
	)
	Describe("Getting data from a $__GNS sentence", func() {
		BeforeEach(func() {
			parsed = GNS{
				Time:       Time{},
				Latitude:   NewFloat64(Latitude),
				Longitude:  NewFloat64(Longitude),
				Mode:       []string{SimulatorGNS},
				SVs:        Int64{},
				HDOP:       Float64{},
				Altitude:   NewFloat64(Altitude),
				Separation: Float64{},
				Age:        Float64{},
				Station:    Int64{},
			}
		})
		Context("When having a parsed sentence", func() {
			It("should give a valid position", func() {
				lat, lon, alt, _ := parsed.GetPosition3D()
				Expect(lat).To(Equal(Latitude))
				Expect(lon).To(Equal(Longitude))
				Expect(alt).To(Equal(Altitude))
			})
		})
	})
})
