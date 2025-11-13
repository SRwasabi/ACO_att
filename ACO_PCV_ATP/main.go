package main

// imports ================================================================================
import (
	"bufio"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ================================================================================
// Graphs and Cities ==================================================================
type City struct {
	ID int
	X  float64
	Y  float64
}

type Graph struct {
	cities			[]City
	cities_distance	[][]float64
	pheromones		[][]float64
}

// ================================================================================
// Ant Colony Optimization (ACO) =====================================================
type Ant struct {
	start		int
	path       []int
	cost       float64
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
}

//================================================================================

func create_CITY(id int, x, y float64) City {
	return City{ID: id, X: x, Y: y}
}

func create_GRAPH() Graph {
	file, err := os.Open("wi29.tsp")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var g Graph
	scanner := bufio.NewScanner(file)
	inSection := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		// Start reading coordinates
		if line == "NODE_COORD_SECTION" {
			inSection = true
			continue
		}

		// Stop at EOF
		if line == "EOF" {
			break
		}

		if inSection {
			fields := strings.Fields(line)
			if len(fields) < 3 {
				continue
			}
			id, _ := strconv.Atoi(fields[0])
			x, _ := strconv.ParseFloat(fields[1], 64)
			y, _ := strconv.ParseFloat(fields[2], 64)
			city := create_CITY(id, x, y)
			g.cities = append(g.cities, city)
		}
	}

	// Inicializa as matrizes de distancia e feromonio
	cities_qtty := len(g.cities)
	g.cities_distance = make([][]float64, cities_qtty)
	g.pheromones = make([][]float64, cities_qtty)

	for from := 0; from < cities_qtty; from++ {
		g.cities_distance[from] = make([]float64, cities_qtty)
		g.pheromones[from] = make([]float64, cities_qtty)

		for to := 0; to < cities_qtty; to++ {
			if from == to {
				g.cities_distance[from][to] = 9999999.0
			} else {
				g.cities_distance[from][to] = distance(g.cities[from], g.cities[to])
			}

			// feromonio inicial: 1.0 pra todas as arestas
			g.pheromones[from][to] = 1.0
		}
	}

	return g
}

func create_ANT(grafo *Graph) Ant {
    n := len(grafo.cities)
    start := rand.Intn(n)

    return Ant{
        start: start,
        path:  []int{start},
        cost:  0.0,
    }
}

func create_ACO(grafo *Graph, num_ants int, alpha, beta, evaporation, constatQ float64, iterations int) ACO {
	ants := make([]Ant, num_ants)
	for i := 0; i < num_ants; i++ {
		ants[i] = create_ANT(grafo)
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
	}
}

func distance(a, b City) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var alpha float64 = 0.5
	var beta float64 = 0.5
	var evaporation float64 = 0.5
	var constatQ float64 = 0.5
	var iteretions int = 10
	var ants int = 5

	g := create_GRAPH()
	println("Loaded cities:", len(g.cities))

	aco := create_ACO(&g, ants, alpha, beta, evaporation, constatQ, iteretions)

	for i := 0; i < ants; i++ {
		startIdx := aco.ants[i].start
		cityID := g.cities[startIdx].ID
		println("Ant", i, "start city ID:", cityID)
	}

}
