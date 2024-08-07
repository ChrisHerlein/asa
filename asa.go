package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var action string
var resultName string
var poleExclude int

func main() {
	flag.StringVar(&action, "action", "", "series to sort qualy into series, final to sort series into final")
	flag.StringVar(&resultName, "resultName", "", "name for result file to be writen")
	flag.IntVar(&poleExclude, "exclude", 0, "amount of polemen to exclude by invalid lap")
	flag.Parse()

	if action != "series" && action != "final" && action != "result" {
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
	case "result":
		err := setResultTable()
		if err != nil {
			fmt.Println("Error generating table:", err)
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

func sortFinal(series ...[]AsaDriver) error {
	// merge series
	var drivers = make([]AsaDriver, 0)

	var maxDrivers int
	for i := 0; i < len(series); i++ {
		if len(series[i]) > maxDrivers {
			maxDrivers = len(series[i])
		}
	}

	for i := 0; i <= maxDrivers*len(series); i++ {
		idx := i % len(series)
		if len(series[idx]) > i/len(series) {
			drivers = append(drivers, series[idx][i/len(series)])
		}
	}

	var lines = make([]string, 0)
	lines = append(lines, "/racelength 1 12")

	for pos, d := range drivers {
		line := fmt.Sprintf("/editgrid %d %s", pos+1, d.Name)
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

func sortSerie(sn int, drivers []AsaDriver) error {
	var lines = make([]string, 0)
	lines = append(lines, "/racelength 1 5")
	var sidx = sn - 1

	var pos = 1
	for i := sidx; i < len(drivers); i = i + 2 {
		line := fmt.Sprintf("/editgrid %d %s", pos, drivers[i].Name)
		lines = append(lines, line)
		pos++
	}

	return os.WriteFile(
		fmt.Sprintf("serie%d.ini", sn),
		[]byte(strings.Join(lines, "\n")),
		0644,
	)
}

func rfDriversToAsa(ds []Driver) []*AsaDriver {
	drivers := make([]*AsaDriver, len(ds))
	for i := 0; i < len(ds); i++ {
		drivers[i] = ds[i].toAsa()
	}
	return drivers
}
