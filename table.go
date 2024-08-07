package main

import (
	"fmt"
	"os"
	"strings"
)

func setResultTable() error {
	raceFile, err := chooseFileName("race")
	if err != nil {
		return err
	}

	race, err := readResult(raceFile)
	if err != nil {
		return err
	}

	if len(race.RaceResults.Race.Drivers) == 0 {
		return fmt.Errorf("is not a race")
	}

	drivers := rfDriversToAsa(race.RaceResults.Race.Drivers)
	drivers = ADClearDNS(drivers)
	ADSort(drivers)

	return resultToCsv(drivers)
}

func resultToCsv(drivers []*AsaDriver) error {
	var lines = make([]string, 0)
	lines = append(lines, "Position,Name,Laps,FinishTime,FinishStatus")
	for _, driver := range drivers {
		lines = append(lines,
			fmt.Sprintf("%d,%s,%d,%f,%s",
				driver.Position, driver.Name,
				driver.Laps, driver.FinishTime,
				driver.FinishStatus))
	}

	return os.WriteFile(
		fmt.Sprintf(resultName),
		[]byte(strings.Join(lines, "\n")),
		0644,
	)
}
