package main

// imports ================================================================================
import (
	"math"
	"math/rand"
)

// ================================================================================
// Ant Colony Optimization (ACO) =====================================================
type Ant struct {
	start int
	path  []int
	cost  float64
}

type ACO struct {
	grafo       *Graph
	ants        []Ant
	alpha       float64
	beta        float64
	evaporation float64
	constatQ    float64
	iterations  int

	bestPath []int
	bestCost float64
	rng      *rand.Rand
}

//================================================================================

func create_ANT(grafo *Graph, rng *rand.Rand) Ant {
	n := len(grafo.cities)
	start := rng.Intn(n)

	return Ant{
		start: start,
		path:  []int{start},
		cost:  0.0,
	}
}

func create_ACO(grafo *Graph, num_ants int, alpha, beta, evaporation, constatQ float64, iterations int) ACO {
	rng := rand.New(rand.NewSource(1))

	ants := make([]Ant, num_ants)
	for i := 0; i < num_ants; i++ {
		ants[i] = create_ANT(grafo, rng)
	}

	return ACO{
		grafo:       grafo,
		ants:        ants,
		alpha:       alpha,
		beta:        beta,
		evaporation: evaporation,
		constatQ:    constatQ,
		iterations:  iterations,
		bestCost:    math.Inf(1),
		rng:         rng,
	}
}

func distance(a, b City) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}
