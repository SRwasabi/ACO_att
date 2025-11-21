package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
    "os"
    
	"github.com/SRwasabi/ACO_att/pkg/aco"
	"github.com/SRwasabi/ACO_att/pkg/config"
    "github.com/SRwasabi/ACO_att/pkg/plotter"
    "github.com/SRwasabi/ACO_att/pkg/graph"
)

func main() {
	matrixCfg := config.Default()	
	jsonPath := config.ParseFlags(&matrixCfg)

	var matrices []config.MatrixConfig

	if jsonPath != "" {
		loaded, err := config.LoadMatrix(jsonPath)
		if err != nil {
			log.Fatalf("Erro JSON: %v", err)
		}
		for i := range loaded {
			config.ParseFlags(&loaded[i])
		}
		matrices = loaded
		fmt.Printf(">> Carregadas %d matrizes de configuração.\n", len(matrices))
	} else {
		matrices = []config.MatrixConfig{matrixCfg}
	}

	totalRuns := 0
	for _, m := range matrices {
		runConfigs := m.Expand()
		totalRuns += len(runConfigs)
		
		for i, runCfg := range runConfigs {
			if runCfg.Seed == 0 {
				runCfg.Seed = time.Now().UnixNano()
			}

			fmt.Printf("\n=== EXECUÇÃO %d/%d (Total: %d) ===\n", i+1, len(runConfigs), totalRuns)
            m.Print()

			executeRun(runCfg)
		}
	}
}

func executeRun(cfg config.RunConfig) {
	Rng := rand.New(rand.NewSource(cfg.Seed))   

    start := time.Now()
        g := graph.NewGraph(cfg.InputFile, Rng) 
        
        if cfg.UseKNN {
            g.ComputeNearestNeighbors(cfg.KNNSize)
        }
        fmt.Printf("\n\nCidades carregadas: %d \n", len(g.Cities))
        elapsed := time.Since(start)
	printDurationStats(elapsed)


    start = time.Now()
        colony := aco.NewAco(g, cfg, Rng)
        colony.Run()
        elapsed = time.Since(start)
    fmt.Print("-Tempo de execução ACO: ")
    printDurationStats(elapsed)
    
	prefix := fmt.Sprintf("%s_Results/", cfg.ExperimentName)
    os.MkdirAll(prefix, os.ModePerm)

	plotter.SaveConvergence(colony.BestCostHistory, prefix+"convergence.png")
	plotter.SaveBestPath(g.Cities, colony.BestPath, colony.BestCost, prefix+"best_path.png")
	plotter.SaveCostStats(colony.BestCostHistory, colony.MeanCostHistory, colony.MaxCostHistory, prefix+"costsStats.png")
	plotter.SaveTiming(colony.TimeConstructHistory, colony.TimeCostHistory, colony.TimePheromoneHistory, prefix+"timing.png")
	plotter.SaveHeatmap(g.InitialPheromones, prefix+"initial_pheromone_heatmap.png", "Initial Pheromone Levels")
	plotter.SaveHeatmap(g.Pheromones, prefix+"final_pheromone_heatmap.png", "Final Pheromone Levels")

	fmt.Printf("Resultados salvos com prefixo: '%s'\n", prefix)
}


func printDurationStats(d time.Duration) {
    total := d

    days := d / (24 * time.Hour)
    d -= days * 24 * time.Hour

    hours := d / time.Hour
    d -= hours * time.Hour

    minutes := d / time.Minute
    d -= minutes * time.Minute

    seconds := d / time.Second
    d -= seconds * time.Second

    milliseconds := d / time.Millisecond

    ticks := total.Nanoseconds() / 100
    totalDays := total.Hours() / 24
    totalHours := total.Hours()
    totalMinutes := total.Minutes()
    totalSeconds := total.Seconds()
    totalMilliseconds := float64(total.Milliseconds())

    fmt.Printf("Days              : %d\n", days)
    fmt.Printf("Hours             : %d\n", hours)
    fmt.Printf("Minutes           : %d\n", minutes)
    fmt.Printf("Seconds           : %d\n", seconds)
    fmt.Printf("Milliseconds      : %d\n", milliseconds)
    fmt.Printf("Ticks             : %d\n", ticks)
    fmt.Printf("TotalDays         : %f\n", totalDays)
    fmt.Printf("TotalHours        : %f\n", totalHours)
    fmt.Printf("TotalMinutes      : %f\n", totalMinutes)
    fmt.Printf("TotalSeconds      : %f\n", totalSeconds)
    fmt.Printf("TotalMilliseconds : %f\n", totalMilliseconds)
}
