package nmea_test

import (
	. "github.com/munnik/go-nmea"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("MWD", func() {
	var (
		sentence Sentence
		parsed   MWD
		err      error
		raw      string
	)
	Describe("Parsing", func() {
		JustBeforeEach(func() {
			sentence, err = Parse(raw)
			if sentence != nil {
				parsed = sentence.(MWD)
			} else {
				parsed = MWD{}
			}
		})
		Context("a valid sentence", func() {
			BeforeEach(func() {
				raw = "$WIMWD,351.1,T,350.8,M,8.4,N,4.3,M*59"
			})
			It("returns no errors", func() {
				Expect(err).NotTo(HaveOccurred())
			})
			It("equals a valid MWD struct", func() {
				Expect(parsed).To(MatchFields(IgnoreExtras, Fields{
					"WindDirectionTrue":          Equal(NewFloat64(351.1)),
					"WindDirectionMagnetic":      Equal(NewFloat64(350.8)),
					"WindSpeedInKnots":           Equal(NewFloat64(8.4)),
					"WindSpeedInMetersPerSecond": Equal(NewFloat64(4.3)),
				}))
			})
		})
		Context("a sentence with a bad checksum", func() {
			BeforeEach(func() {
				raw = "$WIMWD,351.1,T,350.8,M,8.4,N,4.3,M*60"
			})
			It("returns an error", func() {
				Expect(err).To(MatchError("nmea: sentence checksum mismatch [59 != 60]"))
			})
			It("returns nil", func() {
				Expect(sentence).To(BeNil())
			})
		})
	})
	Describe("Getting data from a MWD struct", func() {
		BeforeEach(func() {
			parsed = MWD{
				WindDirectionTrue:          NewFloat64(TrueDirectionDegrees),
				WindDirectionMagnetic:      NewFloat64(MagneticDirectionDegrees),
				WindSpeedInKnots:           NewFloat64(SpeedOverGroundKnots),
				WindSpeedInMetersPerSecond: NewFloat64(SpeedOverGroundMPS),
			}
		})
		Context("when having a complete struct", func() {
			It("returns a valid true wind direction", func() {
				Expect(parsed.GetTrueWindDirection()).To(Float64Equal(TrueDirectionRadians, 0.00001))
			})
			It("returns a valid magnetic wind direction", func() {
				Expect(parsed.GetMagneticWindDirection()).To(Float64Equal(MagneticDirectionRadians, 0.00001))
			})
			It("returns a valid wind speed", func() {
				Expect(parsed.GetWindSpeed()).To(Float64Equal(SpeedOverGroundMPS, 0.00001))
			})
		})
		Context("when having a struct with missing data", func() {
			JustBeforeEach(func() {
				parsed = MWD{}
			})
			It("returns an error", func() {
				_, err := parsed.GetMagneticWindDirection()
				Expect(err).To(HaveOccurred())
			})
			It("returns an error", func() {
				_, err := parsed.GetTrueWindDirection()
				Expect(err).To(HaveOccurred())
			})
			It("returns an error", func() {
				_, err := parsed.GetWindSpeed()
				Expect(err).To(HaveOccurred())
			})
		})
	})
})