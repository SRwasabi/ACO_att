package aco

import (
	"fmt"
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


func (aco *ACO) SaveBestPathPlot(filename string) error {
    if len(aco.BestPath) == 0 {
        return fmt.Errorf("BestPath está vazio, rode Run() antes de plotar")
    }

    p := plot.New()

    p.Title.Text = fmt.Sprintf("Best Tour (Cost = %.2f)", aco.BestCost)
    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"

    n := len(aco.BestPath)
    pathPoints := make(plotter.XYs, n+1)
    cityPoints := make(plotter.XYs, n)

    for i, idx := range aco.BestPath {
        c := aco.Grafo.Cities[idx]
        pathPoints[i].X = c.X
        pathPoints[i].Y = c.Y

        cityPoints[i].X = c.X
        cityPoints[i].Y = c.Y
    }

    first := aco.Grafo.Cities[aco.BestPath[0]]
    pathPoints[n].X = first.X
    pathPoints[n].Y = first.Y

    line, err := plotter.NewLine(pathPoints)
    if err != nil {
        return err
    }

    scatter, err := plotter.NewScatter(cityPoints)
    if err != nil {
        return err
    }

    p.Add(line, scatter)
    p.Legend.Add("best path", line)
    p.Legend.Add("cities", scatter)

    scatter.Radius = vg.Points(2)

    return p.Save(10*vg.Inch, 5*vg.Inch, filename)
}
