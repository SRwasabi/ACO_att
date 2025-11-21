package aco

import (
    "fmt"
    "math"
    "math/rand"
    "strings"
    "sync"
    "time"
    "github.com/SRwasabi/ACO_att/pkg/config"
)

type ACO struct {
    Cfg config.RunConfig

    Grafo                  *Graph
    Ants                    []Ant

    BestPath                []int
    BestCost                float64
    
    // Historicos 
    BestCostHistory         []float64
    MeanCostHistory         []float64
    MaxCostHistory          []float64
    TimeConstructHistory    []float64
    TimeCostHistory         []float64
    TimePheromoneHistory    []float64

    Rng *rand.Rand
}

func NewAco(Grafo *Graph, cfg config.RunConfig, Rng *rand.Rand) *ACO {
    Ants := make([]Ant, cfg.NumAnts)
    for i := 0; i < cfg.NumAnts; i++ {
        Ants[i] = NewAnt(Grafo, Rng)
    }

    Grafo.PrecomputeHeuristics(cfg.Alpha)
    return &ACO{
        Cfg:             cfg,
        Grafo:           Grafo,
        Ants:            Ants,
        BestCost:        math.Inf(1),
        
        BestCostHistory:        make([]float64, 0, cfg.Iterations),
        MeanCostHistory:        make([]float64, 0, cfg.Iterations),
        MaxCostHistory:         make([]float64, 0, cfg.Iterations),
        TimeConstructHistory:   make([]float64, 0, cfg.Iterations),
        TimeCostHistory:        make([]float64, 0, cfg.Iterations),
        TimePheromoneHistory:   make([]float64, 0, cfg.Iterations),

        Rng: Rng,
    }
}

func (aco *ACO) Run() {
    barWidth := 50

    targetUpdates := aco.Cfg.Iterations
    step := aco.Cfg.Iterations / targetUpdates
    if step <= 0 {
        step = 1
    }

    for iter := 0; iter < aco.Cfg.Iterations; iter++ {
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

        if (iter+1)%step == 0 || iter == aco.Cfg.Iterations-1 {
            printProgressBar(iter+1, aco.Cfg.Iterations, barWidth)
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
