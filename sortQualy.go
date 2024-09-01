package main

func removeExtraLaps(results *RFactorXML) *RFactorXML {
	for i := 0; i < len(results.RaceResults.Qualify.Drivers); i++ {
		d := results.RaceResults.Qualify.Drivers[i]
		if len(d.Lap) > 3 {
			results.RaceResults.Qualify.Drivers[i].BestLapTime = d.Lap[1].Text
		}
	}

	// ,pver todos los bestLapTime == 0 al fondo
	for i := 0; i < len(results.RaceResults.Qualify.Drivers); i++ {
		if results.RaceResults.Qualify.Drivers[i].BestLapTime == "" ||
			results.RaceResults.Qualify.Drivers[i].BestLapTime == "--.----" {
			driver := results.RaceResults.Qualify.Drivers[i]
			driver.BestLapTime = "999.9999"
			for j := i + 1; j < len(results.RaceResults.Qualify.Drivers); j++ {
				results.RaceResults.Qualify.Drivers[j-1] = results.RaceResults.Qualify.Drivers[j]
			}
			results.RaceResults.Qualify.Drivers[len(results.RaceResults.Qualify.Drivers)-1] = driver
		}
	}

	return results
}

func removeInvalidPole(original []AsaDriver) []AsaDriver {
	for i := 0; i < poleExclude; i++ {
		poleSitter := original[0]
		drivers := make([]AsaDriver, len(original))
		for j := 1; j < len(original); j++ {
			drivers[j-1] = original[j]
		}
		drivers[len(original)-1] = poleSitter
		original = drivers
	}
	return original
}
