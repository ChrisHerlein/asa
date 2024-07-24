package main

import (
	"encoding/xml"
	"strconv"
)

type RFactorXML struct {
	XMLName     xml.Name    `xml:"rFactorXML"`
	Text        string      `xml:",chardata"`
	Version     string      `xml:"version,attr"`
	RaceResults RaceResults `xml:"RaceResults"`
}

type ConnectionType struct {
	Text     string `xml:",chardata"`
	Upload   string `xml:"upload,attr"`
	Download string `xml:"download,attr"`
}

type AsaDriver struct {
	Name         string
	Laps         int
	Position     int
	FinishTime   float64
	FinishStatus string
}

func (d *Driver) toAsa() *AsaDriver {
	laps, _ := strconv.Atoi(d.Laps)
	finishTime, _ := strconv.ParseFloat(d.FinishTime, 64)
	pos, _ := strconv.Atoi(d.Position)
	return &AsaDriver{
		Name:         d.Name,
		Laps:         laps,
		Position:     pos,
		FinishTime:   finishTime,
		FinishStatus: d.FinishStatus,
	}
}

type Driver struct {
	Text                   string `xml:",chardata"`
	Name                   string `xml:"Name"`
	Connected              string `xml:"Connected"`
	VehFile                string `xml:"VehFile"`
	UpgradeCode            string `xml:"UpgradeCode"`
	VehName                string `xml:"VehName"`
	CarType                string `xml:"CarType"`
	CarClass               string `xml:"CarClass"`
	CarNumber              string `xml:"CarNumber"`
	TeamName               string `xml:"TeamName"`
	IsPlayer               string `xml:"isPlayer"`
	Position               string `xml:"Position"`
	ClassPosition          string `xml:"ClassPosition"`
	Points                 string `xml:"Points"`
	ClassPoints            string `xml:"ClassPoints"`
	LapRankIncludingDiscos string `xml:"LapRankIncludingDiscos"`
	Lap                    []struct {
		Text string `xml:",chardata"`
		Num  string `xml:"num,attr"`
		P    string `xml:"p,attr"`
		Et   string `xml:"et,attr"`
		Fuel string `xml:"fuel,attr"`
		S1   string `xml:"s1,attr"`
		S2   string `xml:"s2,attr"`
		S3   string `xml:"s3,attr"`
	} `xml:"Lap"`
	BestLapTime    string `xml:"BestLapTime"`
	Laps           string `xml:"Laps"`
	Pitstops       string `xml:"Pitstops"`
	FinishStatus   string `xml:"FinishStatus"`
	FinishTime     string `xml:"FinishTime"`
	ControlAndAids []struct {
		Text     string `xml:",chardata"`
		StartLap string `xml:"startLap,attr"`
		EndLap   string `xml:"endLap,attr"`
	} `xml:"ControlAndAids"`
	Swap struct {
		Text     string `xml:",chardata"`
		StartLap string `xml:"startLap,attr"`
		EndLap   string `xml:"endLap,attr"`
	} `xml:"Swap"`
	PenaltyMass string `xml:"PenaltyMass"`
}

type Qualy struct {
	Text       string `xml:",chardata"`
	DateTime   string `xml:"DateTime"`
	TimeString string `xml:"TimeString"`
	Laps       string `xml:"Laps"`
	Minutes    string `xml:"Minutes"`
	Stream     struct {
		Text string `xml:",chardata"`
		Chat []struct {
			Text string `xml:",chardata"`
			Et   string `xml:"et,attr"`
		} `xml:"Chat"`
		Score []struct {
			Text string `xml:",chardata"`
			Et   string `xml:"et,attr"`
		} `xml:"Score"`
		Incident []struct {
			Text string `xml:",chardata"`
			Et   string `xml:"et,attr"`
		} `xml:"Incident"`
	} `xml:"Stream"`
	MostLapsCompleted string   `xml:"MostLapsCompleted"`
	Drivers           []Driver `xml:"Driver"`
}

type RaceResults struct {
	Text            string         `xml:",chardata"`
	Setting         string         `xml:"Setting"`
	ServerName      string         `xml:"ServerName"`
	PlayerFile      string         `xml:"PlayerFile"`
	DateTime        string         `xml:"DateTime"`
	TimeString      string         `xml:"TimeString"`
	Mod             string         `xml:"Mod"`
	Season          string         `xml:"Season"`
	TrackVenue      string         `xml:"TrackVenue"`
	TrackCourse     string         `xml:"TrackCourse"`
	TrackEvent      string         `xml:"TrackEvent"`
	TrackLength     string         `xml:"TrackLength"`
	GameVersion     string         `xml:"GameVersion"`
	Dedicated       string         `xml:"Dedicated"`
	ConnectionType  ConnectionType `xml:"ConnectionType"`
	RaceLaps        string         `xml:"RaceLaps"`
	RaceTime        string         `xml:"RaceTime"`
	MechFailRate    string         `xml:"MechFailRate"`
	DamageMult      string         `xml:"DamageMult"`
	FuelMult        string         `xml:"FuelMult"`
	TireMult        string         `xml:"TireMult"`
	VehiclesAllowed string         `xml:"VehiclesAllowed"`
	ParcFerme       string         `xml:"ParcFerme"`
	FixedSetups     string         `xml:"FixedSetups"`
	FreeSettings    string         `xml:"FreeSettings"`
	FixedUpgrades   string         `xml:"FixedUpgrades"`
	Qualify         Qualy          `xml:"Qualify"`
	Race            Race           `xml:"Race"`
}

type Race struct {
	Text              string   `xml:",chardata"`
	DateTime          string   `xml:"DateTime"`
	TimeString        string   `xml:"TimeString"`
	Laps              string   `xml:"Laps"`
	Minutes           string   `xml:"Minutes"`
	MostLapsCompleted string   `xml:"MostLapsCompleted"`
	Length            float64  `xml:"-"`
	Drivers           []Driver `xml:"Driver"`
}
