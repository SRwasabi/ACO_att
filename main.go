package main

import (
	"github.com/SRwasabi/ACO_att/aco"
	"math/rand"
	// "time"
)

func main() {
	var alpha float64 = 0.5
	var beta float64 = 0.5
	var evaporation float64 = 0.2
	var constatQ float64 = 0.5
	var iteretions int = 10
	var ants int = 5
	// Rng := rand.NewSource(time.Now().UnixNano())
	Rng := rand.New(rand.NewSource(1))

	g := aco.CreateGRAPH(Rng)
	println("Loaded Cities:", len(g.Cities))

	colony := aco.CreateACO(g, ants, alpha, beta, evaporation, constatQ, iteretions, Rng)
	colony.Run()



}
