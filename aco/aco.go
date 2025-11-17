package aco

import (
	"math"
	"math/rand"
)

type ACO struct {
	Grafo       *Graph
	Ants        []Ant
	Alpha       float64
	Beta        float64
	Evaporation float64
	ConstatQ    float64
	Iterations  int

	BestPath        []int
	BestCost        float64
	BestCostHistory []float64
	Rng             *rand.Rand
}

func CreateACO(Grafo *Graph, num_Ants int, Alpha, Beta, Evaporation, ConstatQ float64, Iterations int, Rng *rand.Rand) ACO {
	Ants := make([]Ant, num_Ants)
	for i := 0; i < num_Ants; i++ {
		Ants[i] = Create_ANT(Grafo, Rng)
	}

	return ACO{
		Grafo:           Grafo,
		Ants:            Ants,
		Alpha:           Alpha,
		Beta:            Beta,
		Evaporation:     Evaporation,
		ConstatQ:        ConstatQ,
		Iterations:      Iterations,
		BestCost:        math.Inf(1),
		BestCostHistory: make([]float64, 0, Iterations),
		Rng:             Rng,
	}
}

func (aco *ACO) Run() {
	for iter := 0; iter < aco.Iterations; iter++ {
		aco.resetAnts()
		aco.constructSolutions()

		PathCOST(aco)
		UpdatePheromones(aco)
		aco.BestCostHistory = append(aco.BestCostHistory, aco.BestCost)
	}
}

func (aco *ACO) resetAnts() {
	for i := range aco.Ants {
		aco.Ants[i] = Create_ANT(aco.Grafo, aco.Rng)
	}
}

func (aco *ACO) constructSolutions() {
	numCities := len(aco.Grafo.Cities)
	for i := range aco.Ants {
		for step := 0; step < numCities-1; step++ {
			NextCITY(&aco.Ants[i], aco)
		}
	}
}
