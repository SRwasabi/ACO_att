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

	return Graph{}
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

}
