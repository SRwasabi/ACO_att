package aco

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func (aco *ACO) SaveConvergencePlot(filename string) error {
	p := plot.New()

	p.Title.Text = "ACO Convergence"
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Best Cost"

	points := make(plotter.XYs, len(aco.BestCostHistory))

	for i := range aco.BestCostHistory {
		points[i].X = float64(i)
		points[i].Y = aco.BestCostHistory[i]
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		return err
	}
	p.Add(line)

	return p.Save(10*vg.Inch, 5*vg.Inch, filename)
}
