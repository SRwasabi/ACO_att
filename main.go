package main

import (
	"math/rand"

	"time"

	"github.com/SRwasabi/ACO_att/aco"
)

func main() {
	var alpha float64 = 0.4
	var beta float64 = 0.3
	var evaporation float64 = 0.5
	var constatQ float64 = 0.3
	var iteretions int = 100
	var ants int = 50
	Rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	//Rng := rand.New(rand.NewSource(1))

	g := aco.CreateGRAPH(Rng)
	println("Loaded Cities:", len(g.Cities))

	colony := aco.CreateACO(g, ants, alpha, beta, evaporation, constatQ, iteretions, Rng)
	colony.Run()
	colony.SaveConvergencePlot("convergence.png")
}
