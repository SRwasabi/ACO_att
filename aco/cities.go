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
	cities          []City
	cities_distance [][]float64
	pheromones      [][]float64
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

			// feromonio inicial: aleatorio entre 0 e 1
			g.pheromones[from][to] = rand.Float64()
		}
	}

	return g
}
