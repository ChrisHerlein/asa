package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func readResult(fname string) ([]AsaDriver, error) {
	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var resultFile RFactorXML
	err = xml.Unmarshal(data, &resultFile)
	return sortDrivers(&resultFile), err
}

func sortDrivers(rf *RFactorXML) []AsaDriver {
	var drivers []AsaDriver
	// if qualy
	if len(rf.RaceResults.Qualify.Drivers) > 0 {
		rf = removeExtraLaps(rf)
		for _, d := range rf.RaceResults.Qualify.Drivers {
			drivers = append(drivers, *d.toAsa())
		}
		ADQualySort(drivers)
	}

	// if race
	if len(rf.RaceResults.Race.Drivers) > 0 {
		for _, d := range rf.RaceResults.Race.Drivers {
			drivers = append(drivers, *d.toAsa())
		}
	}

	drivers = ADClearDNS(drivers)
	ADSort(drivers)

	return drivers
}

const resultsPath = "./UserData/LOG/Results"

func chooseFileName(guide string) (string, error) {
	ls, err := os.ReadDir(resultsPath)
	if err != nil {
		return "", nil
	}

	fmt.Println("Choose file for", guide)
	for i := 0; i < len(ls); i++ {
		fmt.Printf("[%d] %s\n", i, ls[i].Name())
	}

	var choosen int
	fmt.Printf("Choose file number: ")
	fmt.Scanf("%d", &choosen)

	if choosen > len(ls) || choosen < 0 {
		return "", fmt.Errorf("Invalid option #%d", choosen)
	}

	return fmt.Sprintf("%s/%s", resultsPath, ls[choosen].Name()), nil
}
