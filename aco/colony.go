package aco

// imports ================================================================================
import (
	"math"
	"math/rand"
)

// ================================================================================
// Ant Colony Optimization (ACO) =====================================================
type Ant struct {
	Start  int
	Path   []int
	Cost   float64
	Actual int
}

type ACO struct {
	Grafo       *Graph
	Ants        []Ant
	Alpha       float64
	Beta        float64
	Evaporation float64
	ConstatQ    float64
	Iterations  int

	BestPath []int
	BestCost float64
	Rng      *rand.Rand
}

//================================================================================

func create_ANT(Grafo *Graph, Rng *rand.Rand) Ant {
	n := len(Grafo.Cities)
	Start := Rng.Intn(n)

	return Ant{
		Start: Start,
		Path:  []int{Start},
		Cost:  0.0,
	}
}

func CreateACO(Grafo *Graph, num_Ants int, Alpha, Beta, Evaporation, ConstatQ float64, Iterations int) ACO {
	Rng := rand.New(rand.NewSource(1))

	Ants := make([]Ant, num_Ants)
	for i := 0; i < num_Ants; i++ {
		Ants[i] = create_ANT(Grafo, Rng)
	}

	return ACO{
		Grafo:       Grafo,
		Ants:        Ants,
		Alpha:       Alpha,
		Beta:        Beta,
		Evaporation: Evaporation,
		ConstatQ:    ConstatQ,
		Iterations:  Iterations,
		BestCost:    math.Inf(1),
		Rng:         Rng,
	}
}

func distance(a, b City) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func nextCity(ant *Ant, aco *ACO) int {
	prob := 0.0
	sum := 0.0
	index := -1
	visited := make(map[int]bool)
	//map para ver se a cidade já foi visitada
	for _, city := range ant.Path {
		visited[city] = true
	}

	for to := 0; to < len(aco.Grafo.Cities); to++ {
		if !visited[to] {
			sum += (1 / aco.Grafo.Cities_distance[ant.Actual][to]) * aco.Grafo.Pheromones[ant.Actual][to]
		}
	}

	for to := 0; to < len(aco.Grafo.Cities); to++ {
		if !visited[to] {
			if prob == 0.0 {
				prob = ((1 / aco.Grafo.Cities_distance[ant.Actual][to]) * aco.Grafo.Pheromones[ant.Actual][to]) / sum
				index = to
			} else {
				aux := ((1 / aco.Grafo.Cities_distance[ant.Actual][to]) * aco.Grafo.Pheromones[ant.Actual][to]) / sum
				if aux > prob {
					prob = aux
					index = to
				}
			}
		}
	}
	return index
}
