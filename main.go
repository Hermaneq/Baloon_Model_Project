package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Initializing data
	var volume float64 = 200
	var weight float64 = 0.241
	var balloonVolume []float64
	var balloonWeight []float64
	balloonVolume = append(balloonVolume, 200)
	balloonWeight = append(balloonWeight, 0)
	rand.Seed(time.Now().UnixNano())
	var time int = 0

	// Creating plots
	p := plot.New()
	w := plot.New()

	// Balloon pumping simulation
	balloonVolume, balloonWeight, time = pumpBalloon(volume, time, weight, balloonVolume, balloonWeight)

	// Creating plots data
	pts := make(plotter.XYs, time)
	wts := make(plotter.XYs, time)

	addData(pts, wts, time, balloonVolume, balloonWeight)
	createLines(pts, p, wts, w)
	setPlotDetails(p, w)

	// Saving plots to the files

	savePlots(p, w)

}

func createLines(pts plotter.XYs, p *plot.Plot, wts plotter.XYs, w *plot.Plot) {
	line, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(line)
	nLine, nErr := plotter.NewLine(wts)
	if nErr != nil {
		log.Fatal(nErr)
	}
	w.Add(nLine)
}

func pumpBalloon(volume float64, time int, weight float64, balloonVolume []float64, balloonWeight []float64) ([]float64, []float64, int) {
	for volume < 15000 {
		time += 1

		if time%3 == 0 {
			volume = volume + 1000 + 3000*rand.Float64()
			weight = volume * 0.001205
			balloonVolume = append(balloonVolume, volume)
			balloonWeight = append(balloonWeight, weight)
		} else {
			balloonVolume = append(balloonVolume, volume)
			balloonWeight = append(balloonWeight, weight)
		}

	}
	return balloonVolume, balloonWeight, time
}

func setPlotDetails(p *plot.Plot, w *plot.Plot) {
	p.Title.Text = "Simulation of the volume of an inflated balloon"
	p.X.Label.Text = "Time [s]"
	p.Y.Label.Text = "Balloon volume [mL]"
	p.X.Min = 0
	p.Y.Min = 0

	w.Title.Text = "Simulation of the weight of an inflated balloon"
	w.X.Label.Text = "Time [s]"
	w.Y.Label.Text = "internal weight of the balloon [g]"
	w.X.Min = 0
	w.Y.Min = 0
}

func addData(pts plotter.XYs, wts plotter.XYs, time int, balloonVolume []float64, balloonWeight []float64) {
	for i := 0; i < time; i++ {
		pts[i].X = float64(i)
		pts[i].Y = balloonVolume[i]
		wts[i].X = float64(i)
		wts[i].Y = balloonWeight[i]
	}
}

func savePlots(p *plot.Plot, w *plot.Plot) {
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "balloonVolume_simulation.png"); err != nil {
		fmt.Println("Error during saving the plot:", err)
	} else {
		fmt.Println("Plot has been saved to the file 'balloonVolume_simulation.png'.")
	}

	if err := w.Save(6*vg.Inch, 4*vg.Inch, "balloonWeight_simulation.png"); err != nil {
		fmt.Println("Error during saving the plot:", err)
	} else {
		fmt.Println("Plot has been saved to the file 'balloonWeight_simulation.png'.")
	}
}
