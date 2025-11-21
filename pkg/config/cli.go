package config

import (
	"flag"
	"fmt"
	"os"
)

func ParseFlags(cfg *MatrixConfig) string {
	fs := flag.NewFlagSet("ACO Solver", flag.ExitOnError)

	configPath := fs.String("config", "", "Caminho para arquivo JSON de configuração")

	var (
		files    StringList
		ants     IntList
		iter     IntList
		alpha    FloatList
		beta     FloatList
		evap     FloatList
		q        FloatList
		knnSize  IntList
		useKNN   BoolList
		roulette BoolList
	)

	var seed int64

	fs.Var(&files, "file", "Lista de arquivos TSP (ex: data.tsp)")
	fs.Var(&ants, "ants", "Lista de formigas (ex: 50,100)")
	fs.Var(&iter, "iter", "Lista de iterações (ex: 100,200)")
	fs.Var(&alpha, "alpha", "Lista de Alpha (ex: 1.0,2.0)")
	fs.Var(&beta, "beta", "Lista de Beta (ex: 2.0,5.0)")
	fs.Var(&evap, "evaporation", "Lista de Evaporação (ex: 0.1,0.5)")
	fs.Var(&q, "q", "Lista de fator Q (ex: 100.0)")
	fs.Var(&knnSize, "knn_size", "Lista de tamanhos KNN (ex: 15,20)")
	fs.Var(&useKNN, "use_knn", "Lista de booleans para usar KNN (ex: true,false)")
	fs.Var(&roulette, "roulette", "Lista de booleans para seleção Roleta (ex: true,false)")
	
	fs.Int64Var(&seed, "seed", 0, "Seed para o RNG (0 = aleatório)")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUso do ACO Solver:\n")
		fmt.Fprintf(os.Stderr, "  Simples: go run . -ants 100 -iter 500\n")
		fmt.Fprintf(os.Stderr, "  Batch:   go run . -config matrix.json\n")
		fmt.Fprintf(os.Stderr, "  Híbrido: go run . -config matrix.json -beta \"2.0,5.0\" -q 200\n\n")
		fmt.Fprintf(os.Stderr, "Opções disponíveis:\n")
		fs.PrintDefaults()
	}

	fs.Parse(os.Args[1:])

	if len(files) > 0 { cfg.InputFile = files }
	if len(ants) > 0 { cfg.NumAnts = ants }
	if len(iter) > 0 { cfg.Iterations = iter }
	if len(alpha) > 0 { cfg.Alpha = alpha }
	if len(beta) > 0 { cfg.Beta = beta }
	if len(evap) > 0 { cfg.Evaporation = evap }
	if len(q) > 0 { cfg.Q = q }
	if len(knnSize) > 0 { cfg.KNNSize = knnSize }
	if len(useKNN) > 0 { cfg.UseKNN = useKNN }
	if len(roulette) > 0 { cfg.RouletteSelection = roulette }
	
	if seed != 0 {
		cfg.Seed = seed
	}

	return *configPath
}