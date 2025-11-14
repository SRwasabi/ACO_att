package aco

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
	Cities          []City
	Cities_distance [][]float64
	Pheromones      [][]float64
}

//================================================================================

func create_CITY(id int, x, y float64) City {
	return City{ID: id, X: x, Y: y}
}

func CreateGRAPH() Graph {
	file, err := os.Open("coordinates/wi29.tsp")
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
			g.Cities = append(g.Cities, city)
		}
	}

	// Inicializa as matrizes de distancia e feromonio
	cities_qtty := len(g.Cities)
	g.Cities_distance = make([][]float64, cities_qtty)
	g.Pheromones = make([][]float64, cities_qtty)

	for from := 0; from < cities_qtty; from++ {
		g.Cities_distance[from] = make([]float64, cities_qtty)
		g.Pheromones[from] = make([]float64, cities_qtty)

		for to := 0; to < cities_qtty; to++ {
			if from == to {
				g.Cities_distance[from][to] = 0.0
			} else {
				g.Cities_distance[from][to] = distance(g.Cities[from], g.Cities[to])
			}

			// feromonio inicial: aleatorio entre 0 e 1
			g.Pheromones[from][to] = rand.Float64()
		}
	}

	return g
}
