package config

import (
	"fmt"
)
type MatrixConfig struct {
	ExperimentName    string      `json:"experiment_name"`
	InputFile         StringList  `json:"input_file"`
	NumAnts           IntList     `json:"num_ants"`
	Iterations        IntList     `json:"iterations"`
	Alpha             FloatList   `json:"alpha"`
	Beta              FloatList   `json:"beta"`
	Evaporation       FloatList   `json:"evaporation"`
	Q                 FloatList   `json:"q"`
	KNNSize           IntList     `json:"knn_size"`
	UseKNN            BoolList    `json:"use_knn"`
	RouletteSelection BoolList    `json:"roulette_selection"`
	Seed              int64       `json:"seed"`
}

type RunConfig struct {
	ExperimentName    string
	InputFile         string
	NumAnts           int
	Iterations        int
	Alpha             float64
	Beta              float64
	Evaporation       float64
	Q                 float64
	KNNSize           int
	UseKNN            bool
	RouletteSelection bool
	Seed              int64
}

func (c MatrixConfig) Print() {
	fmt.Printf("Config Matrix [%s]:\n", c.ExperimentName)
	fmt.Printf("    %-18s: %v\n", "InputFile", c.InputFile)
	fmt.Printf("    %-18s: %v\n", "NumAnts", c.NumAnts)
	fmt.Printf("    %-18s: %v\n", "Iterations", c.Iterations)
	fmt.Printf("    %-18s: %v\n", "Alpha", c.Alpha)
	fmt.Printf("    %-18s: %v\n", "Beta", c.Beta)
	fmt.Printf("    %-18s: %v\n", "Evaporation", c.Evaporation)
	fmt.Printf("    %-18s: %v\n", "Q", c.Q)
	fmt.Printf("    %-18s: %v\n", "KNNSize", c.KNNSize)
	fmt.Printf("    %-18s: %v\n", "UseKNN", c.UseKNN)
	fmt.Printf("    %-18s: %v\n", "RouletteSelection", c.RouletteSelection)
	fmt.Printf("    %-18s: %v\n", "Seed", c.Seed)
}