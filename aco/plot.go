package aco

import (
	"fmt"
	"math"
	"sort"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/palette/moreland"
)

// ====================== CONVERGÊNCIA ======================

func (aco *ACO) SaveConvergencePlot(filename string) error {
	p := plot.New()

	p.Title.Text = "ACO Convergence (Best Cost)"
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Best Cost"

	p.Add(plotter.NewGrid())

	points := make(plotter.XYs, len(aco.BestCostHistory))
	for i := range aco.BestCostHistory {
		points[i].X = float64(i)
		points[i].Y = aco.BestCostHistory[i]
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		return err
	}
	line.LineStyle.Width = vg.Points(1.5)
	line.LineStyle.Color = color.RGBA{R: 30, G: 144, B: 255, A: 255} // dodger blue

	p.Add(line)
	p.Legend.Add("Best cost", line)
	p.Legend.Top = true

	return p.Save(10*vg.Inch, 5*vg.Inch, filename)
}

// ====================== MELHOR CAMINHO ======================

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

	// fecha o ciclo
	first := aco.Grafo.Cities[aco.BestPath[0]]
	pathPoints[n].X = first.X
	pathPoints[n].Y = first.Y

	line, err := plotter.NewLine(pathPoints)
	if err != nil {
		return err
	}
	line.LineStyle.Width = vg.Points(1.3)
	line.LineStyle.Color = color.RGBA{R: 255, G: 99, B: 71, A: 255} // tomato

	scatter, err := plotter.NewScatter(cityPoints)
	if err != nil {
		return err
	}
	scatter.Radius = vg.Points(2.5)
	scatter.GlyphStyle.Color = color.RGBA{R: 25, G: 25, B: 112, A: 255} // midnight blue

	p.Add(line, scatter)
	p.Legend.Add("Best path", line)
	p.Legend.Add("Cities", scatter)
	p.Legend.Top = true

	return p.Save(8*vg.Inch, 8*vg.Inch, filename)
}

func (aco *ACO) SaveCostStatsPlot(filename string) error {
	n := len(aco.BestCostHistory)
	if n == 0 || len(aco.MeanCostHistory) != n || len(aco.MaxCostHistory) != n {
		return fmt.Errorf("históricos de custo inconsistentes ou vazios")
	}

	p := plot.New()
	p.Title.Text = "ACO Costs per Iteration"
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Cost"

	p.Add(plotter.NewGrid())

	bestPts := make(plotter.XYs, n)
	meanPts := make(plotter.XYs, n)
	maxPts := make(plotter.XYs, n)

	for i := 0; i < n; i++ {
		x := float64(i)
		bestPts[i].X = x
		bestPts[i].Y = aco.BestCostHistory[i]

		meanPts[i].X = x
		meanPts[i].Y = aco.MeanCostHistory[i]

		maxPts[i].X = x
		maxPts[i].Y = aco.MaxCostHistory[i]
	}

	bestLine, err := plotter.NewLine(bestPts)
	if err != nil {
		return err
	}
	meanLine, err := plotter.NewLine(meanPts)
	if err != nil {
		return err
	}
	maxLine, err := plotter.NewLine(maxPts)
	if err != nil {
		return err
	}

	bestLine.LineStyle.Width = vg.Points(1.8)
	bestLine.LineStyle.Color = color.RGBA{R: 30, G: 144, B: 255, A: 255} // azul

	meanLine.LineStyle.Width = vg.Points(1.3)
	meanLine.LineStyle.Color = color.RGBA{R: 60, G: 179, B: 113, A: 255} // verde

	maxLine.LineStyle.Width = vg.Points(1.0)
	maxLine.LineStyle.Dashes = []vg.Length{vg.Points(3), vg.Points(3)}
	maxLine.LineStyle.Color = color.RGBA{R: 220, G: 20, B: 60, A: 255} // vermelho

	p.Add(bestLine, meanLine, maxLine)
	p.Legend.Add("Best", bestLine)
	p.Legend.Add("Mean", meanLine)
	p.Legend.Add("Max", maxLine)
	p.Legend.Top = true

	return p.Save(10*vg.Inch, 5*vg.Inch, filename)
}

// ====================== TIMING (BARRAS) ======================

func (aco *ACO) SaveTimingPlot(filename string) error {
    n := len(aco.TimeConstructHistory)
    if n == 0 || len(aco.TimeCostHistory) != n || len(aco.TimePheromoneHistory) != n {
        return fmt.Errorf("históricos de tempo inconsistentes ou vazios")
    }

    var totalConstruct, totalCost, totalPhero float64
    for i := 0; i < n; i++ {
        totalConstruct += aco.TimeConstructHistory[i]
        totalCost += aco.TimeCostHistory[i]
        totalPhero += aco.TimePheromoneHistory[i]
    }

    avgConstruct := totalConstruct / float64(n)
    avgCost := totalCost / float64(n)
    avgPhero := totalPhero / float64(n)

    p := plot.New()
    p.Title.Text = "Average Time per Iteration by Phase"
    p.Y.Label.Text = "Time (seconds per iteration)"
    p.X.Label.Text = ""

    p.NominalX("Construct\nSolutions", "PathCOST", "Update\nPheromones")

    values := plotter.Values{avgConstruct, avgCost, avgPhero}

    barWidth := vg.Points(30)
    bars, err := plotter.NewBarChart(values, barWidth)
    if err != nil {
        return err
    }

    bars.LineStyle.Width = vg.Points(0.5)
    bars.LineStyle.Color = color.RGBA{R: 70, G: 130, B: 180, A: 255} // steel blue
    bars.Color = color.RGBA{R: 135, G: 206, B: 235, A: 255}          // light sky blue

    p.Add(bars)
    p.Add(plotter.NewGrid())

    xy := make(plotter.XYs, len(values))
    labelStrings := []string{
        fmt.Sprintf("%.4f", avgConstruct),
        fmt.Sprintf("%.4f", avgCost),
        fmt.Sprintf("%.4f", avgPhero),
    }

    for i, v := range values {
        xy[i].X = float64(i)
        xy[i].Y = v + (v * 0.05)
        if v == 0 {
            xy[i].Y = 0.01
        }
    }

    labels, err := plotter.NewLabels(plotter.XYLabels{
        XYs:    xy,
        Labels: labelStrings,
    })
    if err == nil {
        p.Add(labels)
    }

    return p.Save(8*vg.Inch, 5*vg.Inch, filename)
}



type pheromoneGrid struct {
	data [][]float64
}

func (g pheromoneGrid) Dims() (c, r int) {
	if len(g.data) == 0 {
		return 0, 0
	}
	return len(g.data[0]), len(g.data)
}

func (g pheromoneGrid) Z(c, r int) float64 {
	return g.data[r][c]
}

func (g pheromoneGrid) X(c int) float64 {
	return float64(c)
}

func (g pheromoneGrid) Y(r int) float64 {
	return float64(r)
}

func pheromoneStats(matrix [][]float64) (minVal, maxVal, mean float64) {
    n := len(matrix)
    if n == 0 {
        return 0, 0, 0
    }

    minVal = math.Inf(1)
    maxVal = math.Inf(-1)
    sum := 0.0
    count := 0

    for i := 0; i < n; i++ {
        row := matrix[i]
        for j := 0; j < len(row); j++ {
            if i == j {
                continue
            }
            v := row[j]
            if v < minVal {
                minVal = v
            }
            if v > maxVal {
                maxVal = v
            }
            sum += v
            count++
        }
    }

    if count == 0 {
        return 0, 0, 0
    }
    mean = sum / float64(count)
    return minVal, maxVal, mean
}

func pheromoneStatsWithQuartiles(matrix [][]float64) (minVal, maxVal, mean, q1, q2, q3 float64) {
    n := len(matrix)
    if n == 0 {
        return 0, 0, 0, 0, 0, 0
    }

    vals := make([]float64, 0, n*n)

    minVal = math.Inf(1)
    maxVal = math.Inf(-1)
    sum := 0.0

    for i := 0; i < n; i++ {
        row := matrix[i]
        for j := 0; j < len(row); j++ {
            if i == j {
                continue
            }
            v := row[j]
            vals = append(vals, v)
            if v < minVal {
                minVal = v
            }
            if v > maxVal {
                maxVal = v
            }
            sum += v
        }
    }

    if len(vals) == 0 {
        return 0, 0, 0, 0, 0, 0
    }

    mean = sum / float64(len(vals))

    sort.Float64s(vals)
    idxQ1 := int(0.25 * float64(len(vals)))
    idxQ2 := int(0.50 * float64(len(vals)))
    idxQ3 := int(0.75 * float64(len(vals)))

    q1 = vals[idxQ1]
    q2 = vals[idxQ2]
    q3 = vals[idxQ3]

    return minVal, maxVal, mean, q1, q2, q3
}

func savePheromoneHeatmap(matrix [][]float64, filename, title string) error {
    if len(matrix) == 0 {
        return fmt.Errorf("matriz de feromônios vazia")
    }

    minVal, maxVal, mean, q1, q2, q3 := pheromoneStatsWithQuartiles(matrix)

    p := plot.New()

    p.Title.Text = fmt.Sprintf(
        "%s\nmin=%.4g  max=%.4g  mean=%.4g\nBlue≈[%.3g,%.3g]  Blue→White≈[%.3g,%.3g]  White→Red≈[%.3g,%.3g]  Red≈[%.3g,%.3g]",
        title, minVal, maxVal, mean,
        minVal, q1,
        q1, q2,
        q2, q3,
        q3, maxVal,
    )

    p.X.Label.Text = "City index"
    p.Y.Label.Text = "City index"

    grid := pheromoneGrid{data: matrix}

    // aqui você está usando:
    pal := moreland.SmoothBlueRed().Palette(255)

    hm := plotter.NewHeatMap(grid, pal)
    p.Add(hm)
    p.Add(plotter.NewGrid())

    return p.Save(8*vg.Inch, 8*vg.Inch, filename)
}



func (aco *ACO) SaveInitialPheromoneHeatmap(filename string) error {
	if aco.Grafo == nil || len(aco.Grafo.InitialPheromones) == 0 {
		return fmt.Errorf("InitialPheromones vazio ou grafo nil")
	}
	return savePheromoneHeatmap(aco.Grafo.InitialPheromones, filename, "Initial Pheromone Levels")
}

func (aco *ACO) SaveFinalPheromoneHeatmap(filename string) error {
	if aco.Grafo == nil || len(aco.Grafo.Pheromones) == 0 {
		return fmt.Errorf("Pheromones vazio ou grafo nil")
	}
	return savePheromoneHeatmap(aco.Grafo.Pheromones, filename, "Final Pheromone Levels")
}
