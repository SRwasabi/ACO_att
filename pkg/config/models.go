package config

import (
	"fmt"
)

func Default() MatrixConfig {
	return MatrixConfig{
		ExperimentName:    "Default",
		InputFile:         StringList{"coordinates/uy734.tsp"},
		NumAnts:           IntList{500},
		Iterations:        IntList{1200},
		Alpha:             FloatList{0.4},
		Beta:              FloatList{0.3},
		Evaporation:       FloatList{0.4},
		Q:                 FloatList{100.0},
		KNNSize:           IntList{20},
		UseKNN:            BoolList{true},
		RouletteSelection: BoolList{true},
		Seed:              0,
	}
}

func (m MatrixConfig) Expand() []RunConfig {
	var runs []RunConfig

	for _, file := range m.InputFile {
	for _, ants := range m.NumAnts {
	for _, iter := range m.Iterations {
	for _, a := range m.Alpha {
	for _, b := range m.Beta {
	for _, e := range m.Evaporation {
	for _, q := range m.Q {
	for _, k := range m.KNNSize {
	for _, useK := range m.UseKNN {
	for _, roul := range m.RouletteSelection {
		
		name := fmt.Sprintf("%s_Ants%d_Iter%d_Alpha%.2f_Beta%.2f_Evap%.2f_Q%.2f_KNN%d_UseKNN%t_Roulette%t_Seed%d", 
			m.ExperimentName, ants, iter, a, b, e, q, k, useK, roul, m.Seed)

		if len(m.InputFile) > 1 {
			name = fmt.Sprintf("%s_%s", name, file)
		}

		runs = append(runs, RunConfig{
			ExperimentName:    name,
			InputFile:         file,
			NumAnts:           ants,
			Iterations:        iter,
			Alpha:             a,
			Beta:              b,
			Evaporation:       e,
			Q:                 q,
			KNNSize:           k,
			UseKNN:            useK,
			RouletteSelection: roul,
			Seed:              m.Seed,
		})
	}}}}}}}}}} 

	return runs
}