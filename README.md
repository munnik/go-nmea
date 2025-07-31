# Deprecated
Use https://github.com/munnik/go-signalk instead.

# go-nmea

[![Build Status](https://github.com/munnik/go-nmea/actions/workflows/go.yml/badge.svg)](https://github.com/munnik/go-nmea/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/munnik/go-nmea)](https://goreportcard.com/report/github.com/munnik/go-nmea) [![Coverage Status](https://coveralls.io/repos/github/munnik/go-nmea/badge.svg?branch=master)](https://coveralls.io/github/munnik/go-nmea?branch=master) ![GitHub](https://img.shields.io/github/license/munnik/go-nmea)

This is a fork of the https://github.com/adrianmo/go-nmea repository.

## Features

- All the futures of https://github.com/adrianmo/go-nmea
- All types (structs) have a `Valid` and `InvalidReason` property
- Supports additional nmea0183 sentences (GST, HEV, MDA, MWD, MWV, ROT, VWR)
- Implement SignalK interface
- Moved to Ginkgo tests

## Installing

To install go-nmea use `go get`:

```
go get github.com/munnik/go-nmea
```

This will then make the `github.com/munnik/go-nmea` package available to you.

### Staying up to date

To update go-nmea to the latest version, use `go get -u github.com/munnik/go-nmea`.

## Supported sentences

At this moment, this library supports the following sentence types:

| Sentence type                                                                       | Description                                                         |
| ----------------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| [DBS](https://gpsd.gitlab.io/gpsd/NMEA.html#_dbs_depth_below_surface)               | Depth Below Surface                                                 |
| [DBT](https://gpsd.gitlab.io/gpsd/NMEA.html#_dbt_depth_below_transducer)            | Depth below transducer                                              |
| [DPT](https://gpsd.gitlab.io/gpsd/NMEA.html#_dpt_depth_of_water)                    | Depth of Water                                                      |
| [GGA](http://aprs.gids.nl/nmea/#gga)                                                | GPS Positioning System Fix Data                                     |
| [GLL](http://aprs.gids.nl/nmea/#gll)                                                | Geographic Position, Latitude / Longitude and time                  |
| [GNS](https://www.trimble.com/oem_receiverhelp/v4.44/en/NMEA-0183messages_GNS.html) | Combined GPS fix for GPS, Glonass, Galileo, and BeiDou              |
| [GSA](http://aprs.gids.nl/nmea/#gsa)                                                | GPS DOP and active satellites                                       |
| [GST](https://gpsd.gitlab.io/gpsd/NMEA.html#_gst_gps_pseudorange_noise_statistics)  | GPS Pseudorange Noise Statistics                                    |
| [GSV](http://aprs.gids.nl/nmea/#gsv)                                                | GPS Satellites in view                                              |
| [HDT](http://aprs.gids.nl/nmea/#hdt)                                                | Actual vessel heading in degrees True                               |
| HEV                                                                                 | Heave in meters                                                     |
| [MDA](https://gpsd.gitlab.io/gpsd/NMEA.html#_mda_meteorological_composite)          | Meteorological Composite                                            |
| MWD                                                                                 | Wind Direction & Speed                                              |
| [MWV](https://gpsd.gitlab.io/gpsd/NMEA.html#_mwv_wind_speed_and_angle)              | Wind Speed and Angle                                                |
| [PGRME](http://aprs.gids.nl/nmea/#rme)                                              | Estimated Position Error (Garmin proprietary sentence)              |
| [PMTK](https://www.rhydolabz.com/documents/25/PMTK_A11.pdf)                         | Messages for setting and reading commands for MediaTek gps modules. |
| [RMC](http://aprs.gids.nl/nmea/#rmc)                                                | Recommended Minimum Specific GPS/Transit data                       |
| [ROT](https://gpsd.gitlab.io/gpsd/NMEA.html#_rot_rate_of_turn)                      | Rate Of Turn                                                        |
| [RTE](http://aprs.gids.nl/nmea/#rte)                                                | Route                                                               |
| [THS](http://www.nuovamarea.net/pytheas_9.html)                                     | Actual vessel heading in degrees True and status                    |
| [VDM/VDO](http://catb.org/gpsd/AIVDM.html)                                          | Encapsulated binary payload                                         |
| [VHW](https://www.tronico.fi/OH6NT/docs/NMEA0183.pdf)                               | Water Speed and Heading                                             |
| [VTG](http://aprs.gids.nl/nmea/#vtg)                                                | Track Made Good and Ground Speed                                    |
| [VWR](https://gpsd.gitlab.io/gpsd/NMEA.html#_vwr_relative_wind_speed_and_angle)     | Relative Wind Speed and Angle                                       |
| [WPL](http://aprs.gids.nl/nmea/#wpl)                                                | Waypoint location                                                   |
| [ZDA](http://aprs.gids.nl/nmea/#zda)                                                | Date & time data                                                    |


If you need to parse a message that contains an unsupported sentence type you can implement and register your own message parser and get yourself unblocked immediately. Check the example below to know how to [implement and register a custom message parser](#custom-message-parsing). However, if you think your custom message parser could be beneficial to other users we encourage you to contribute back to the library by submitting a PR and get it included in the list of supported sentences.

## Contributing

Please feel free to submit issues or fork the repository and send pull requests to update the library and fix bugs, implement support for new sentence types, refactor code, etc.

## License

Check [LICENSE](LICENSE).
