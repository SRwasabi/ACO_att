package aco

import (
    "fmt"
    "math"
    "math/rand"
    "strings"
    "sync"
    "time"
)

type ACO struct {
    Grafo                  *Graph
    Ants                    []Ant
    Alpha                   float64
    Beta                    float64
    Evaporation             float64
    ConstatQ                float64
    Iterations              int
    RouletteSelection       bool
    UseKNN                  bool

    BestPath                []int
    BestCost                float64
    
    BestCostHistory         []float64
    MeanCostHistory         []float64
    MaxCostHistory          []float64

    TimeConstructHistory    []float64
    TimeCostHistory         []float64
    TimePheromoneHistory    []float64

    Rng *rand.Rand
}

func CreateACO(Grafo *Graph, num_Ants int, Alpha, Beta, Evaporation, ConstatQ float64, Iterations int, Rng *rand.Rand, rouletteSelection bool, useKNN bool) ACO {
    Ants := make([]Ant, num_Ants)
    for i := 0; i < num_Ants; i++ {
        Ants[i] = Create_ANT(Grafo, Rng)
    }

    Grafo.PrecomputeHeuristics(Alpha)
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
        MeanCostHistory: make([]float64, 0, Iterations),
        MaxCostHistory:  make([]float64, 0, Iterations),

        TimeConstructHistory: make([]float64, 0, Iterations),
        TimeCostHistory:      make([]float64, 0, Iterations),
        TimePheromoneHistory: make([]float64, 0, Iterations),
        RouletteSelection:    rouletteSelection,
        UseKNN:               useKNN,
        Rng: Rng,
    }
}

func (aco *ACO) Run() {
    barWidth := 50

    targetUpdates := aco.Iterations
    step := aco.Iterations / targetUpdates
    if step <= 0 {
        step = 1
    }

    for iter := 0; iter < aco.Iterations; iter++ {
        tStartConstruct := time.Now() // PARA MEDIR TEMPO 
            aco.resetAnts()
            aco.constructSolutions()
        tConstruct := time.Since(tStartConstruct).Seconds() // PARA MEDIR TEMPO 
        aco.TimeConstructHistory = append(aco.TimeConstructHistory, tConstruct) // PARA MEDIR TEMPO 

        tStartCost := time.Now() // PARA MEDIR TEMPO 
            meanCost, maxCost := PathCOST(aco) 
            aco.updateCostHistories(meanCost, maxCost)
        tCost := time.Since(tStartCost).Seconds() // PARA MEDIR TEMPO 
        aco.TimeCostHistory = append(aco.TimeCostHistory, tCost) // PARA MEDIR TEMPO 


        tStartPhero := time.Now() // PARA MEDIR TEMPO 
            UpdatePheromones(aco)
            tPhero := time.Since(tStartPhero).Seconds() // PARA MEDIR TEMPO 
        aco.TimePheromoneHistory = append(aco.TimePheromoneHistory, tPhero) // PARA MEDIR TEMPO 

        if (iter+1)%step == 0 || iter == aco.Iterations-1 {
            printProgressBar(iter+1, aco.Iterations, barWidth)
        }
    }
}

func (aco *ACO) updateCostHistories(meanCost, maxCost float64) {
    aco.BestCostHistory = append(aco.BestCostHistory, aco.BestCost)
    aco.MeanCostHistory = append(aco.MeanCostHistory, meanCost)
    aco.MaxCostHistory = append(aco.MaxCostHistory, maxCost)
}

func (aco *ACO) resetAnts() {
    n := len(aco.Grafo.Cities)

	for i := range aco.Ants {
		newStart := aco.Ants[i].Rng.Intn(n)
        aco.Ants[i].Reset(newStart)
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
