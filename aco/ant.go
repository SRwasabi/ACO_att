package aco

import (
	"math/rand"
)

type Ant struct {
	Start         int
	Path          []int
	Cost          float64
	Actual        int
	Qtd_pheromone float64
	Visited       []bool
}

func Create_ANT(Grafo *Graph, Rng *rand.Rand) Ant {
	n := len(Grafo.Cities)
	Start := Rng.Intn(n)
	print("start ->", Start, "\n")
	visited := make([]bool, n)
	visited[Start] = true

	return Ant{
		Start:         Start,
		Path:          []int{Start},
		Cost:          0.0,
		Actual:        Start,
		Visited:       visited,
		Qtd_pheromone: 0.0,
	}
}
