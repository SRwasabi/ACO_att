package main

// imports ================================================================================
import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// ================================================================================
// Graphs and Cities ==================================================================
type City struct {
	ID int
	X  float64
	Y  float64
}

type Graph struct {
	cities     []City
	pheromones map[City]map[City]float64
}

// ================================================================================
// Ant Colony Optimization (ACO) =====================================================
type Ant struct {
	start_city City
	path       []City
	cost       float64
}

type ACO struct {
	grafo       Graph
	ants        []Ant
	alpha       float64
	beta        float64
	evaporation float64
	constatQ    float64
	iterations  int
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

	g.pheromones = make(map[City]map[City]float64)

	for _, from := range g.cities {
		g.pheromones[from] = make(map[City]float64)

		for _, to := range g.cities {
			g.pheromones[from][to] = rand.Float64()
			if from == to {
				g.pheromones[from][to] = 9.99
			}

		}

	}

	return g
}

func create_ANT(grafo Graph) Ant {
	start_city := grafo.cities[rand.Intn(len(grafo.cities))]
	return Ant{start_city: start_city, path: []City{start_city}, cost: 0.0}
}

func create_ACO(grafo Graph, num_ants int, alpha, beta, evaporation, constatQ float64, iterations int) ACO {
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
	}
}

func main() {
	var alpha float64 = 0.5
	var beta float64 = 0.5
	var evaporation float64 = 0.5
	var constatQ float64 = 0.5
	var iteretions int = 10
	var ants int = 5

	g := create_GRAPH()
	println("Loaded cities:", len(g.cities))

	aco := create_ACO(g, ants, alpha, beta, evaporation, constatQ, iteretions)
	for i := 0; i < ants; i++ {
		println(aco.ants[i].start_city.ID)
	}
}
