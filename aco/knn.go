package aco

// imports ================================================================================
import (
	"sort"
	"fmt"
)

// ================================================================================
// neighbor struct ==================================================================
type neighbor struct {
	index 	int
	dist 	float64
}

// ================================================================================

func (g *Graph) ComputeNearestNeighbors(k int) {
	n := len(g.Cities)
	g.NearestNeighbors = makeSquareMatrixInt(n)

	for i := 0; i < n; i++ {
		g.NearestNeighbors[i] = g.getKNearestForCity(i, k)
	}
}

func (g *Graph) getKNearestForCity(cityIndex, k int) []int {
	candidates := g.collectNeighbors(cityIndex)
	return selectTopKIndices(candidates, k)
}

func (g *Graph) collectNeighbors(cityIndex int) []neighbor {
	n := len(g.Cities)
	candidates := make([]neighbor, 0, n-1)

	for j := 0; j < n; j++ {
		if cityIndex == j {
			continue
		}
		candidates = append(candidates, neighbor{
			index: 	j,
			dist:	g.Cities_distance[cityIndex][j],
		})
	}
	return candidates
}

func selectTopKIndices(candidates []neighbor, k int) []int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].dist < candidates[j].dist
	})

	limit := k
	if limit > len(candidates) {
		limit = len(candidates)
	}

	result := make([]int, limit)
	for i := 0; i < limit; i++ {
		result[i] = candidates[i].index
	}

	return result
}

func (g *Graph) LogNeighbors(cityIndex int) {
    neighbors := g.NearestNeighbors[cityIndex]
    fmt.Printf("Vizinhos da cidade %d (Total: %d):\n", cityIndex, len(neighbors))
    
    for i, targetIndex := range neighbors {
        dist := g.Cities_distance[cityIndex][targetIndex]
        fmt.Printf("  %dº: Cidade %d (Dist: %.2f)\n", i+1, targetIndex, dist)
    }
}

func makeSquareMatrixInt(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}