package main

func chooseFasterSerie(pos int, series ...[]AsaDriver) []AsaDriver {
	fasterTime := make([]float64, len(series))
	fromSerie := make([]int, len(series))

	for i := 0; i < len(series); i++ {
	JLoop:
		for j := 0; j < len(series[i]); j++ {
			if series[i][j].Position == 1 {
				fasterTime[i] = series[i][j].FinishTime
				fromSerie[i] = i
				break JLoop
			}
		}
	}

	// Sort faster
	for i := 0; i < len(fasterTime); i++ {
		for j := i; j < len(fasterTime)-1; j++ {
			if fasterTime[j] > fasterTime[j+1] {
				a1 := fasterTime[j]
				fasterTime[j] = fasterTime[j+1]
				fasterTime[j+1] = a1

				a2 := fromSerie[j]
				fromSerie[j] = fromSerie[j+1]
				fromSerie[j+1] = a2
			}
		}
	}

	return series[fromSerie[pos]]
}
