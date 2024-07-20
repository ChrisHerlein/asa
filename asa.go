package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var action string
var poleExclude int

func main() {
	flag.StringVar(&action, "action", "", "series to sort qualy into series, final to sort series into final")
	flag.IntVar(&poleExclude, "exclude", 0, "amount of polemen to exclude by invalid lap")
	flag.Parse()

	if action != "series" && action != "final" {
		fmt.Println("Invalid action:", action)
		os.Exit(2)
	}

	switch action {
	case "series":
		err := setSeries()
		if err != nil {
			fmt.Println("Error generating series:", err)
			os.Exit(2)
		}
	case "final":
		err := setFinal()
		if err != nil {
			fmt.Println("Error generating final:", err)
			os.Exit(2)
		}
	}
	fmt.Println("Generated ok!")
}

func setFinal() error {
	serie1file, err := chooseFileName("serie1")
	if err != nil {
		return err
	}

	serie1, err := readResult(serie1file)
	if err != nil {
		return err
	}

	serie2file, err := chooseFileName("serie2")
	if err != nil {
		return err
	}

	serie2, err := readResult(serie2file)
	if err != nil {
		return err
	}

	s1 := chooseFasterSerie(0, serie1, serie2)
	s2 := chooseFasterSerie(1, serie1, serie2)

	return sortFinal(s1, s2)
}

func sortFinal(s1, s2 *RFactorXML) error {
	var lines = make([]string, 0)
	lines = append(lines, "/racelength 1 12")

	for i := 0; i < len(s1.RaceResults.Race.Drivers); i++ {
		if s1.RaceResults.Race.Drivers[i].FinishStatus == "None" {
			continue
		}
		driverPosStr := s1.RaceResults.Race.Drivers[i].Position
		driverPos, _ := strconv.Atoi(driverPosStr)
		line := fmt.Sprintf("/editgrid %d %s", driverPos*2-1, s1.RaceResults.Race.Drivers[i].Name)
		lines = append(lines, line)
	}

	for i := 0; i < len(s2.RaceResults.Race.Drivers); i++ {
		if s2.RaceResults.Race.Drivers[i].FinishStatus == "None" {
			continue
		}
		driverPosStr := s2.RaceResults.Race.Drivers[i].Position
		driverPos, _ := strconv.Atoi(driverPosStr)
		line := fmt.Sprintf("/editgrid %d %s", driverPos*2, s2.RaceResults.Race.Drivers[i].Name)
		lines = append(lines, line)
	}

	return os.WriteFile(
		fmt.Sprintf("final.ini"),
		[]byte(strings.Join(lines, "\n")),
		0644,
	)
}

func setSeries() error {
	fname, err := chooseFileName("qualy")
	if err != nil {
		return err
	}

	qualy, err := readResult(fname)
	if err != nil {
		return err
	}

	qualy = removeExtraLaps(qualy)
	qualy = removeInvalidPole(qualy)

	err = sortSerie(1, qualy)
	if err != nil {
		return err
	}

	err = sortSerie(2, qualy)
	if err != nil {
		return err
	}

	return nil
}

func sortSerie(sn int, q *RFactorXML) error {
	var lines = make([]string, 0)
	lines = append(lines, "/racelength 1 5")
	var sidx = sn - 1

	var pos = 1
	for i := sidx; i < len(q.RaceResults.Qualify.Drivers); i = i + 2 {
		line := fmt.Sprintf("/editgrid %d %s", pos, q.RaceResults.Qualify.Drivers[i].Name)
		lines = append(lines, line)
		pos++
	}

	return os.WriteFile(
		fmt.Sprintf("serie%d.ini", sn),
		[]byte(strings.Join(lines, "\n")),
		0644,
	)
}
