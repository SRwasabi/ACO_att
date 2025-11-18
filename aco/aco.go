package aco

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"sync"
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
    barWidth := 50

    targetUpdates := aco.Iterations

    step := aco.Iterations / targetUpdates
	fmt.Printf("Updating Progress bar at each %d iterations:\n", step)
    if step <= 0 {
        step = 1
    }

    for iter := 0; iter < aco.Iterations; iter++ {
        aco.resetAnts()
        aco.constructSolutions()

        PathCOST(aco)
        UpdatePheromones(aco)
        aco.BestCostHistory = append(aco.BestCostHistory, aco.BestCost)

        if (iter+1)%step == 0 || iter == aco.Iterations-1 {
            printProgressBar(iter+1, aco.Iterations, barWidth)
        }
    }

	logBestSolution(aco.BestPath, aco.BestCost)
}



func (aco *ACO) resetAnts() {
	for i := range aco.Ants {
		aco.Ants[i] = Create_ANT(aco.Grafo, aco.Rng)
	}
}

func (aco *ACO) constructSolutions() {
    numCities := len(aco.Grafo.Cities)

    var wg sync.WaitGroup
    wg.Add(len(aco.Ants))

    for i := range aco.Ants {
        go func(i int) {
            defer wg.Done()

            ant := &aco.Ants[i]
            for step := 0; step < numCities-1; step++ {
                NextCITY(ant, aco)
            }
        }(i)
    }

    wg.Wait()
}

func printProgressBar(current, total, width int) {
    if total <= 0 {
        return
    }

    percent := float64(current) / float64(total)
    filled := int(percent * float64(width))
    if filled > width {
        filled = width
    }

    bar := strings.Repeat("█", filled) + strings.Repeat(" ", width-filled)

    fmt.Printf("\r[%s] %6.2f%% (%d/%d)", bar, percent*100, current, total)

    if current == total {
        fmt.Println()
    }
}
