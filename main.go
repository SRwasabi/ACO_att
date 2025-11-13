package main

func main() {
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
