package main

import (
	"github.com/SRwasabi/ACO_att/aco"
)

func main() {
	var alpha float64 = 0.5
	var beta float64 = 0.5
	var evaporation float64 = 0.2
	var constatQ float64 = 0.5
	var iteretions int = 10
	var ants int = 5

	g := aco.CreateGRAPH()
	println("Loaded Cities:", len(g.Cities))

	colony := aco.CreateACO(&g, ants, alpha, beta, evaporation, constatQ, iteretions)

	for i := 0; i < ants; i++ {
		startIdx := colony.Ants[i].Start
		cityID := g.Cities[startIdx].ID
		println("Ant", i, "start city ID:", cityID)
	}

	for i := 0; i < iteretions; i++ {
		for f := 0; f < ants; f++ {
			print("Iteration:", i, " Ant:", f, "\n")
			for c := 0; c < len(g.Cities); c++ {
				aco.NextCITY(&colony.Ants[f], &colony)
			}
		}
		aco.PathCOST(&colony)
		aco.UpdatePheromones(&colony)
	}
}
