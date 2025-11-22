package aco

import (
	"fmt"
	"math"
	"sync"
)

// ===================== Escolha da próxima cidade ==============================

func NextCITY(ant *Ant, aco *ACO) {
	nextIndex := selectNextCity(ant, aco)
	if nextIndex == -1 {
		return
	}
	moveAntToCity(ant, nextIndex)
}

func selectNextCity(ant *Ant, aco *ACO) int {
	result := -1

	if aco.Cfg.UseKNN {
		result = trySelectKNN(ant, aco)
		if result != -1 {
			return result
		}
	}

	result = trySelectGlobal(ant, aco)
	return result
}

func trySelectKNN(ant *Ant, aco *ACO) int {
	neighbors := aco.Grafo.NearestNeighbors[ant.Actual]
	desirabilities := ant.DesirabilityCache
	total := 0.0

	candidates := make([]int, 0, len(neighbors))

	for _, neighborIndex := range neighbors {
		if ant.Visited[neighborIndex] {
			continue
		}

		val := transitionDesirability(ant.Actual, neighborIndex, aco)
		
		desirabilities[neighborIndex] = val
		total += val
		candidates = append(candidates, neighborIndex)
	}
	
	if len(candidates) == 0 {
		return -1
	}

	if aco.Cfg.RouletteSelection {
		return selectRouletteCity(ant, total, candidates)
	}
	
	return selectGreedyCity(ant, total, candidates)
}

func trySelectGlobal(ant *Ant, aco *ACO) int {
	total := computeDesirabilities(ant, aco)
	if total == 0 {
		return -1
	}
	
	if (aco.Cfg.RouletteSelection) {
		return selectRouletteCity(ant, total, nil)
	}
	
	return selectGreedyCity(ant, total, nil)
}

func selectRouletteCity(ant *Ant, total float64, candidates []int) int {
	r := ant.Rng.Float64() * total
	cumulative := 0.0
	desirabilities := ant.DesirabilityCache

	if candidates != nil {
		for _, i := range candidates {
			cumulative += desirabilities[i]
			if r <= cumulative {
				return i
			}
		}
		return candidates[len(candidates)-1]
	}

	for i, d := range desirabilities {
		if d <= 0 {
			continue
		}
		cumulative += d
		if r <= cumulative {
			return i
		}
	}

	return -1
}

func selectGreedyCity(ant *Ant, total float64, candidates []int) int {
	desirabilities := ant.DesirabilityCache
	bestProb := -1.0
	bestIndex := -1

	checkIndex := func(i int) {
		d := desirabilities[i]
		if d <= 0 {
			return
		}
		prob := d / total
		if prob > bestProb {
			bestProb = prob
			bestIndex = i
		}
	}
	
	if candidates != nil {
		for _, i := range candidates {
			checkIndex(i)
		}
	} else {
		for i := range desirabilities {
			checkIndex(i)
		}
	}

    return bestIndex
}


func computeDesirabilities(ant *Ant, aco *ACO) float64 {
	n := len(aco.Grafo.Cities)
	total := 0.0
	buffer := ant.DesirabilityCache

	for to := 0; to < n; to++ {
		buffer[to] = 0.0
		if ant.Visited[to] || to == ant.Actual {
			continue
		}
		val := transitionDesirability(ant.Actual, to, aco)
		
		buffer[to] = val
		total += val
	}

	return total
}

func transitionDesirability(from, to int, aco *ACO) float64 {
	heuristic := aco.Grafo.HeuristicMatrix[from][to]
	pheromone := aco.Grafo.Pheromones[from][to]

	if pheromone <= 0 {
		return 0
	}

	return heuristic * math.Pow(pheromone, aco.Cfg.Beta)
}

func moveAntToCity(ant *Ant, cityIndex int) {
	ant.Path = append(ant.Path, cityIndex)
	ant.Actual = cityIndex
	ant.Visited[cityIndex] = true
}

func PathCOST(aco *ACO) (meanCost, maxCost float64) {
    computeAntCostsParallel(aco)
    meanCost, maxCost = aggregateCostsAndUpdateBest(aco)
    return meanCost, maxCost
}


func computeAntCostsParallel(aco *ACO) {
    var wg sync.WaitGroup

    for i := range aco.Ants {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            ant := &aco.Ants[i]
            ant.Cost = computePathCost(ant.Path, aco.Grafo.DistanceMatrix)
            ant.PheromoneAmount = pheromoneAmount(aco.Cfg.Q, ant.Cost)
        }(i)
    }

    wg.Wait()
}

func aggregateCostsAndUpdateBest(aco *ACO) (meanCost, maxCost float64) {
    aco.BestCost = math.Inf(1)
    aco.BestPath = nil

    sum := 0.0
    maxCost = 0.0

    for i := range aco.Ants {
        ant := &aco.Ants[i]

        sum += ant.Cost
        if ant.Cost > maxCost {
            maxCost = ant.Cost
        }

        updateBestSolution(aco, ant)
    }

    if len(aco.Ants) > 0 {
        meanCost = sum / float64(len(aco.Ants))
    } else {
        meanCost = 0
    }

    return meanCost, maxCost
}

func computePathCost(path []int, distances [][]float64) float64 {
	if len(path) == 0 {
		return 0
	}

	total := 0.0
	for i := 0; i < len(path); i++ {
		from := path[i]
		to := path[(i+1)%len(path)] // último volta para o primeiro
		total += distances[from][to]
	}

	return total
}

func pheromoneAmount(Q, cost float64) float64 {
	if cost == 0 {
		return 0
	}
	return Q / cost
}

func updateBestSolution(aco *ACO, ant *Ant) {
	if ant.Cost >= aco.BestCost {
		return
	}

	aco.BestCost = ant.Cost
	aco.BestPath = make([]int, len(ant.Path))
	copy(aco.BestPath, ant.Path)
}

func logAntCost(index int, cost float64) {
	fmt.Printf("Formiga %d - Custo do caminho: %.2f\n", index, cost)
}

func logBestSolution(bestPath []int, bestCost float64) {
	fmt.Print("Melhor caminho até agora: ")
	for _, city := range bestPath {
		fmt.Printf("%d ", city)
	}
	fmt.Printf("\nMelhor custo até agora: %.2f\n", bestCost)
}

func UpdatePheromones(aco *ACO) {
	evaporatePheromones(aco)
	reinforceBestPath(aco)
}

func evaporatePheromones(aco *ACO) {
	factor := 1 - aco.Cfg.Evaporation

	for i := range aco.Grafo.Pheromones {
		for j := range aco.Grafo.Pheromones[i] {
			if i == j {
				continue
			}
			aco.Grafo.Pheromones[i][j] *= factor
		}
	}
}

func reinforceBestPath(aco *ACO) {
	if len(aco.BestPath) < 2 {
		return
	}

	for k := 0; k < len(aco.BestPath)-1; k++ {
		fromCity := aco.BestPath[k]
		toCity := aco.BestPath[k+1]

		amount := allPheromoneAmountToDeposite(aco, fromCity, toCity)

		aco.Grafo.Pheromones[fromCity][toCity] += amount
		aco.Grafo.Pheromones[toCity][fromCity] += amount
	}
}

func allPheromoneAmountToDeposite(aco *ACO, fromCity, toCity int) float64 {
	delta := 0.0

	for i := range aco.Ants {
		ant := &aco.Ants[i]
		if antUsesEdge(ant.Path, fromCity, toCity) {
			delta += ant.PheromoneAmount
		}
	}

	return delta
}

func antUsesEdge(path []int, fromCity, toCity int) bool {
	for i := 0; i < len(path)-1; i++ {
		if path[i] == fromCity && path[i+1] == toCity {
			return true
		}
	}
	return false
}
