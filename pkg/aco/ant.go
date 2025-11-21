package aco

import (
	"math/rand"
    "github.com/SRwasabi/ACO_att/pkg/graph"
)

type Ant struct {
	Start               int
	Path                []int
	Cost                float64
	Actual              int
	PheromoneAmount       float64
	Visited             []bool
    DesirabilityCache   []float64
	Rng                 *rand.Rand
}

func NewAnt(Grafo *graph.Graph, baseRng *rand.Rand) Ant {
    n               := len(Grafo.Cities)
    seed            := baseRng.Int63()
    antRng          := rand.New(rand.NewSource(seed))

    start           := antRng.Intn(n)
    visited         := make([]bool, n)
    visited[start]  = true


    return Ant{
        Start:              start,
        Path:               make([]int, 0, n),
        Cost:               0.0,
        Actual:             start,
        Visited:            visited,
        PheromoneAmount:    0.0,
        DesirabilityCache:  make([]float64, n),
        Rng:                antRng,
    }
}

func (a *Ant) Reset(startNode int) {
    a.Start = startNode
    a.Actual = startNode
    a.Cost = 0.0
    a.PheromoneAmount = 0.0

    a.Path = a.Path[:0]
    a.Path = append(a.Path, startNode)

    for i := range a.Visited {
        a.Visited[i] = false
    }
    a.Visited[startNode] = true
}