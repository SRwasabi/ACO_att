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
	Rng           *rand.Rand
}

func Create_ANT(Grafo *Graph, baseRng *rand.Rand) Ant {
    n := len(Grafo.Cities)
    start := baseRng.Intn(n)

    visited := make([]bool, n)
    visited[start] = true

    seed := baseRng.Int63()
    antRng := rand.New(rand.NewSource(seed))

    return Ant{
        Start:         start,
        Path:          []int{start},
        Cost:          0.0,
        Actual:        start,
        Visited:       visited,
        Qtd_pheromone: 0.0,
        Rng:           antRng,
    }
}

