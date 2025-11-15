package aco

// imports ================================================================================
import (
	"bufio"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// ================================================================================
// Graphs and Cities ==================================================================

type Graph struct {
	Cities          []City
	Cities_distance [][]float64
	Pheromones      [][]float64
}

//================================================================================

func CreateGRAPH(Rng *rand.Rand) *Graph {
	const tspFilePath = "coordinates/wi29.tsp"

	cities := loadCitiesFromTSP(tspFilePath)
	distances := buildDistanceMatrix(cities)
	pheromones := buildInitialPheromoneMatrix(len(cities), Rng)

	return &Graph{
		Cities:          cities,
		Cities_distance: distances,
		Pheromones:      pheromones,
	}
}

func loadCitiesFromTSP(path string) []City {
	file := openFile(path)
	defer file.Close()

	return parseCitiesFromReader(file)
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func parseCitiesFromReader(r io.Reader) []City {
	scanner := bufio.NewScanner(r)

	var (
		cities    []City
		inSection bool
	)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if line == "NODE_COORD_SECTION"{
			inSection = true
			continue
		}

		if line == "EOF" {
			break
		}

		if !inSection {
			continue
		}

		city, ok := parseCityLine(line)
		if !ok {
			continue
		}

		cities = append(cities, city)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return cities
}

func parseCityLine(line string) (City, bool) {
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return City{}, false
	}

	id, _ := strconv.Atoi(fields[0])
	x, _ := strconv.ParseFloat(fields[1], 64)
	y, _ := strconv.ParseFloat(fields[2], 64)
	return Create_CITY(id, x, y), true
}

func buildDistanceMatrix(cities []City) [][]float64 {
	n := len(cities)
	matrix := makeSquareMatrix(n)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				continue
			}
			d := distance(cities[i], cities[j])
			matrix[i][j] = d
			matrix[j][i] = d
		}
	}

	return matrix
}

func buildInitialPheromoneMatrix(size int, Rng *rand.Rand) [][]float64 {
	matrix := makeSquareMatrix(size)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			value := Rng.Float64()
			matrix[i][j] = value
			matrix[j][i] = value
		}
	}

	return matrix
}

func makeSquareMatrix(n int) [][]float64 {
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
	}
	return matrix
}

func distance(a, b City) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}