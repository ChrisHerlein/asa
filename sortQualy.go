package main

import "fmt"

func removeExtraLaps(results *RFactorXML) *RFactorXML {
	for i := 0; i < len(results.RaceResults.Qualify.Drivers); i++ {
		d := results.RaceResults.Qualify.Drivers[i]
		if len(d.Lap) > 3 {
			results.RaceResults.Qualify.Drivers[i].BestLapTime = d.Lap[1].Text
		}
	}

	sorted := false
	for i := 0; i < len(results.RaceResults.Qualify.Drivers) && !sorted; i++ {
		sorted = true
		for j := i; j < len(results.RaceResults.Qualify.Drivers)-1; j++ {
			if results.RaceResults.Qualify.Drivers[j].BestLapTime > results.RaceResults.Qualify.Drivers[j+1].BestLapTime &&
				results.RaceResults.Qualify.Drivers[j].BestLapTime != "" && results.RaceResults.Qualify.Drivers[j+1].BestLapTime != "" {
				sorted = false
				aux := results.RaceResults.Qualify.Drivers[j]
				results.RaceResults.Qualify.Drivers[j] = results.RaceResults.Qualify.Drivers[j+1]
				results.RaceResults.Qualify.Drivers[j+1] = aux
			}
		}
	}

	for i := 0; i < len(results.RaceResults.Qualify.Drivers); i++ {
		results.RaceResults.Qualify.Drivers[i].Position = fmt.Sprintf("%d", i+1)
	}

	return results
}

func removeInvalidPole(results *RFactorXML) *RFactorXML {
	for i := 0; i < poleExclude; i++ {
		poleSitter := results.RaceResults.Qualify.Drivers[0]
		drivers := make([]Driver, len(results.RaceResults.Qualify.Drivers))
		for j := 1; j < len(results.RaceResults.Qualify.Drivers); j++ {
			drivers[j-1] = results.RaceResults.Qualify.Drivers[j]
		}
		drivers[len(results.RaceResults.Qualify.Drivers)-1] = poleSitter
		results.RaceResults.Qualify.Drivers = drivers
	}
	return results
}
