package main

import (
	"github.com/SRwasabi/ACO_att/aco"
)

func main() {
	var alpha float64 = 0.5
	var beta float64 = 0.5
	var evaporation float64 = 0.5
	var constatQ float64 = 0.5
	var iteretions int = 10
	var ants int = 5

	g := aco.CreateGRAPH()
	println("Loaded Cities:", len(g.Cities))

	aco := aco.CreateACO(&g, ants, alpha, beta, evaporation, constatQ, iteretions)

	for i := 0; i < ants; i++ {
		startIdx := aco.Ants[i].Start
		cityID := g.Cities[startIdx].ID
		println("Ant", i, "start city ID:", cityID)
	}

}
