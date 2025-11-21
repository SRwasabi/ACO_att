package main

import (
	"math/rand"
	"fmt"
	"time"

	"github.com/SRwasabi/ACO_att/aco"
)

func main() {
	var alpha float64 = 0.4
	var beta float64 = 0.3
	var evaporation float64 = 0.4
	var constatQ float64 = 100
	var iteretions int = 120
	var ants int = 50
	// Rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	Rng := rand.New(rand.NewSource(1))

	start := time.Now()
		g := aco.CreateGRAPH(Rng)
		println("Loaded Cities:", len(g.Cities))
    elapsed := time.Since(start) 
    fmt.Printf("Time taken to load cities: \n")
    printDurationStats(elapsed) 

	start = time.Now()
	colony := aco.CreateACO(g, ants, alpha, beta, evaporation, constatQ, iteretions, Rng)
	colony.Run()
    elapsed = time.Since(start) 
	fmt.Printf("Time taken to run ACO: \n")
    printDurationStats(elapsed) 

	colony.SaveConvergencePlot("convergence.png")
	colony.SaveBestPathPlot("best_path.png")
    colony.SaveCostStatsPlot("costsStats.png")
    colony.SaveTimingPlot("timing.png")
    colony.SaveInitialPheromoneHeatmap("initial_pheromone_heatmap.png")
    colony.SaveFinalPheromoneHeatmap("final_pheromone_heatmap.png")
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
