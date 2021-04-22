package nmea_test

import (
	"testing"

	. "github.com/munnik/go-nmea"
	"github.com/stretchr/testify/assert"
)

var mtktests = []struct {
	name string
	raw  string
	err  string
	msg  MTK
}{
	{
		name: "good: Packet Type: 001 PMTK_ACK",
		raw:  "$PMTK001,604,3*" + Checksum("PMTK001,604,3"),
		msg: MTK{
			Cmd:  NewInt64(604),
			Flag: NewInt64(3),
		},
	},
	{
		name: "missing flag",
		raw:  "$PMTK001,604*" + Checksum("PMTK001,604"),
		err:  "nmea: PMTK001 invalid flag: index out of range",
	},
	{
		name: "missing cmd",
		raw:  "$PMTK001*" + Checksum("PMTK001"),
		err:  "nmea: PMTK001 invalid command: index out of range",
	},
}

func TestMTK(t *testing.T) {
	for _, tt := range mtktests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				mtk := m.(MTK)
				mtk.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, mtk)
			}
		})
	}
}
