package aco

// imports ================================================================================
import (
	"fmt"
	"math"
	"math/rand"
)

// ================================================================================
// Ant Colony Optimization (ACO) =====================================================
type Ant struct {
	Start         int
	Path          []int
	Cost          float64
	Actual        int
	Qtd_pheromone float64
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
		Start:         Start,
		Path:          []int{Start},
		Cost:          0.0,
		Actual:        Start,
		Qtd_pheromone: 0.0,
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

func NextCITY(ant *Ant, aco *ACO) {
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
			sum += math.Pow((1/aco.Grafo.Cities_distance[ant.Actual][to]), aco.Alpha) * math.Pow(aco.Grafo.Pheromones[ant.Actual][to], aco.Beta)
		}
	}

	for to := 0; to < len(aco.Grafo.Cities); to++ {
		if !visited[to] {
			if prob == 0.0 {

				prob = math.Pow((1/aco.Grafo.Cities_distance[ant.Actual][to]), aco.Alpha) * math.Pow(aco.Grafo.Pheromones[ant.Actual][to], aco.Beta) / sum
				index = to
			} else {
				aux := math.Pow((1/aco.Grafo.Cities_distance[ant.Actual][to]), aco.Alpha) * math.Pow(aco.Grafo.Pheromones[ant.Actual][to], aco.Beta) / sum
				if aux > prob {
					prob = aux
					index = to
				}
			}
		}
	}

	if !(index == -1) {
		ant.Path = append(ant.Path, index)
		ant.Actual = index
	}
}

func PathCOST(aco *ACO) {

	for a := 0; a < len(aco.Ants); a++ {
		aco.Ants[a].Cost = 0.0
		for c := 0; c < len(aco.Ants[a].Path); c++ {
			if c == len(aco.Ants[a].Path)-1 {
				aco.Ants[a].Cost += aco.Grafo.Cities_distance[aco.Ants[a].Path[c]][aco.Ants[a].Path[0]]
			} else {
				aco.Ants[a].Cost += aco.Grafo.Cities_distance[aco.Ants[a].Path[c]][aco.Ants[a].Path[c+1]]
			}
		}

		aco.Ants[a].Qtd_pheromone = aco.ConstatQ / aco.Ants[a].Cost
		print("Formiga ", a)
		fmt.Printf(" Custo do caminho da formiga: %.2f \n", aco.Ants[a].Cost)
		if aco.Ants[a].Cost < aco.BestCost {
			aco.BestCost = aco.Ants[a].Cost
			aco.BestPath = make([]int, len(aco.Ants[a].Path))
			copy(aco.BestPath, aco.Ants[a].Path)
		}
	}

	for i := 0; i < len(aco.BestPath); i++ {
		print(" ", aco.BestPath[i])
	}
	fmt.Printf("\nMelhor custo até agora: %.2f\n", aco.BestCost)

}

func UpdatePheromones(aco *ACO) {
	for from := 0; from < len(aco.Grafo.Pheromones); from++ {
		for to := 0; to < len(aco.Grafo.Pheromones); to++ {
			if from != to {
				aco.Grafo.Pheromones[from][to] *= (1 - aco.Evaporation)
			}
		}
	}

	for path := 0; path < len(aco.BestPath); path++ {
		if path < (len(aco.BestPath) - 1) {
			for a := 0; a < len(aco.Ants); a++ {
				for c := 0; c < len(aco.Ants[a].Path); c++ {
					if c < (len(aco.Ants[a].Path) - 1) {
						if aco.Ants[a].Path[c] == aco.BestPath[path] && aco.Ants[a].Path[c+1] == aco.BestPath[path+1] {
							aco.Grafo.Pheromones[path][path+1] += aco.Ants[a].Qtd_pheromone
							aco.Grafo.Pheromones[path+1][path] += aco.Ants[a].Qtd_pheromone
						}
					}
				}
			}
		}
	}
}
