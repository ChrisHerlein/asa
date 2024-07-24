package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
)

func readResult(fname string) (*RFactorXML, error) {
	data, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	var resultFile RFactorXML
	return sortDrivers(&resultFile), xml.Unmarshal(data, &resultFile)
}

func sortDrivers(rf *RFactorXML) []AsaDriver {
	var drivers = make([]Driver, len(rf.RaceResults.Qualify.Drivers))
	for i := 0; i < len(rf.RaceResults.Qualify.Drivers); i++ {
		pos, _ := strconv.Atoi(rf.RaceResults.Qualify.Drivers[i].Position)
		drivers[pos-1] = rf.RaceResults.Qualify.Drivers[i]
	}
	rf.RaceResults.Qualify.Drivers = drivers

	var driversR = make([]Driver, len(rf.RaceResults.Race.Drivers))
	for i := 0; i < len(rf.RaceResults.Race.Drivers); i++ {
		pos, _ := strconv.Atoi(rf.RaceResults.Race.Drivers[i].Position)
		driversR[pos-1] = rf.RaceResults.Race.Drivers[i]
	}
	rf.RaceResults.Race.Drivers = driversR

	if len(driversR) > 0 {
		ln, _ := strconv.ParseFloat(driversR[0].FinishTime, 64)
		rf.RaceResults.Race.Length = ln
	}

	return rf
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
