package aco

import (
	"fmt"
	"math"
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
	desirabilities := computeDesirabilities(ant, aco)
	total := sum(desirabilities)
	if total == 0 {
		return -1
	}

	bestIndex := -1
	bestProb := -1.0

	for city, desirability := range desirabilities {
		if desirability == 0 {
			continue
		}
		prob := desirability / total
		if prob > bestProb {
			bestProb = prob
			bestIndex = city
		}
	}

	return bestIndex
}

func computeDesirabilities(ant *Ant, aco *ACO) []float64 {
	n := len(aco.Grafo.Cities)
	desirabilities := make([]float64, n)

	for to := 0; to < n; to++ {
		if ant.Visited[to] || to == ant.Actual {
			continue
		}
		desirabilities[to] = transitionDesirability(ant.Actual, to, aco)
	}

	return desirabilities
}

func transitionDesirability(from, to int, aco *ACO) float64 {
	dist := aco.Grafo.Cities_distance[from][to]
	if dist <= 0 {
		return 0
	}

	visibility := 1.0 / dist
	pheromone := aco.Grafo.Pheromones[from][to]

	return math.Pow(visibility, aco.Alpha) * math.Pow(pheromone, aco.Beta)
}

func moveAntToCity(ant *Ant, cityIndex int) {
	ant.Path = append(ant.Path, cityIndex)
	ant.Actual = cityIndex
	ant.Visited[cityIndex] = true
}

func sum(values []float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total
}

func PathCOST(aco *ACO) {
	for i := range aco.Ants {
		ant := &aco.Ants[i]

		ant.Cost = computePathCost(ant.Path, aco.Grafo.Cities_distance)
		ant.Qtd_pheromone = pheromoneAmount(aco.ConstatQ, ant.Cost)

		logAntCost(i, ant.Cost)
		updateBestSolution(aco, ant)
	}

	logBestSolution(aco.BestPath, aco.BestCost)
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
	factor := 1 - aco.Evaporation

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
			delta += ant.Qtd_pheromone
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