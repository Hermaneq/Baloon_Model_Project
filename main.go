package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Initializing data
	var radius float64 = 6
	var surface float64 = 4 * radius * radius * math.Pi
	var startSurface float64 = surface
	var tension float64 = (0.5 / startSurface) * (1 - (math.Pow(startSurface, 3) / (math.Pow(surface, 3)))) * 1000
	var volume float64 = 200
	var weight float64 = 0.241
	var balloonTension []float64
	var balloonVolume []float64
	var balloonWeight []float64
	var balloonRadius []float64
	var balloonSurface []float64
	balloonTension = append(balloonTension, tension)
	balloonSurface = append(balloonSurface, surface)
	balloonVolume = append(balloonVolume, 200)
	balloonWeight = append(balloonWeight, 0)
	balloonRadius = append(balloonRadius, radius)
	rand.Seed(time.Now().UnixNano())
	var time int = 0

	// Creating plots
	p := plot.New()
	w := plot.New()
	d := plot.New()
	q := plot.New()
	t := plot.New()

	// Balloon pumping simulation
	balloonTension, balloonVolume, balloonWeight, balloonRadius, balloonSurface, time = pumpBalloon(tension, startSurface, volume, time, weight, radius, surface, balloonSurface, balloonVolume, balloonWeight, balloonRadius, balloonTension)

	// Creating plots data
	pts := make(plotter.XYs, time)
	wts := make(plotter.XYs, time)
	dts := make(plotter.XYs, time)
	qts := make(plotter.XYs, time)
	tts := make(plotter.XYs, time)

	addData(pts, wts, dts, qts, tts, time, balloonVolume, balloonWeight, balloonRadius, balloonSurface, balloonTension)
	createLines(pts, p, wts, w, dts, d, qts, q, tts, t)
	setPlotDetails(p, w, d, q, t)

	// Saving plots to the files

	savePlots(p, w, d, q, t)

}

func createLines(pts plotter.XYs, p *plot.Plot, wts plotter.XYs, w *plot.Plot, dts plotter.XYs, d *plot.Plot, qts plotter.XYs, q *plot.Plot, tts plotter.XYs, t *plot.Plot) {
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
	dLine, dErr := plotter.NewLine(dts)
	if dErr != nil {
		log.Fatal(dErr)
	}
	d.Add(dLine)
	qLine, qErr := plotter.NewLine(qts)
	if qErr != nil {
		log.Fatal(qErr)
	}
	q.Add(qLine)
	tLine, tErr := plotter.NewLine(tts)
	if tErr != nil {
		log.Fatal(tErr)
	}
	t.Add(tLine)
}

func pumpBalloon(tension float64, startSurface float64, volume float64, time int, weight float64, radius float64, surface float64, ballonSurface []float64, balloonVolume []float64, balloonWeight []float64, balloonRadius []float64, balloonTension []float64) ([]float64, []float64, []float64, []float64, []float64, int) {
	for volume < 15000 {
		time += 1

		if time%3 == 0 {
			volume = volume + 1000 + 3000*rand.Float64()
			weight = volume * 0.001205
			number := ((3.0 / 4.0) * volume) / math.Pi
			radius = math.Pow(number, 1.0/3.0)
			surface = 4 * radius * radius * math.Pi
			tension = (1 / startSurface) * (1 - (math.Pow(startSurface, 3) / (math.Pow(surface, 3)))) * 1000

			balloonVolume = append(balloonVolume, volume)
			balloonWeight = append(balloonWeight, weight)
			balloonRadius = append(balloonRadius, radius)
			ballonSurface = append(ballonSurface, surface)
			balloonTension = append(balloonTension, tension)
		} else {
			balloonVolume = append(balloonVolume, volume)
			balloonWeight = append(balloonWeight, weight)
			balloonRadius = append(balloonRadius, radius)
			ballonSurface = append(ballonSurface, surface)
			balloonTension = append(balloonTension, tension)
		}

	}
	for i := 0; i < 4; i++ {
		balloonVolume = append(balloonVolume, 0)
		balloonWeight = append(balloonWeight, 0)
		balloonRadius = append(balloonRadius, 0)
		ballonSurface = append(ballonSurface, 0)
		balloonTension = append(balloonTension, 0)
		time += 1
	}
	return balloonTension, balloonVolume, balloonWeight, balloonRadius, ballonSurface, time
}

func setPlotDetails(p *plot.Plot, w *plot.Plot, d *plot.Plot, q *plot.Plot, t *plot.Plot) {
	p.Title.Text = "Simulation of the volume of an inflated balloon"
	p.X.Label.Text = "Time [s]"
	p.Y.Label.Text = "Balloon volume [mL]"
	p.X.Min = 0
	p.Y.Min = 0

	w.Title.Text = "Simulation of the weight of an inflated balloon"
	w.X.Label.Text = "Time [s]"
	w.Y.Label.Text = "Internal weight of the balloon [g]"
	w.X.Min = 0
	w.Y.Min = 0

	d.Title.Text = "Simulation of the radius of an inflated balloon"
	d.X.Label.Text = "Time [s]"
	d.Y.Label.Text = "Balloon radius [cm]"
	d.X.Min = 0
	d.Y.Min = 0

	q.Title.Text = "Simulation of the surface area of an inflated balloon"
	q.X.Label.Text = "Time [s]"
	q.Y.Label.Text = "Internal weight of the balloon [cm^2]"
	q.X.Min = 0
	q.Y.Min = 0

	t.Title.Text = "Simulation of the surface tension of an inflated balloon"
	t.X.Label.Text = "Time [s]"
	t.Y.Label.Text = "Surface tension of the balloon"
	t.X.Min = 0
	t.Y.Min = 0
}

func addData(pts plotter.XYs, wts plotter.XYs, dts plotter.XYs, qts plotter.XYs, tts plotter.XYs, time int, balloonVolume []float64, balloonWeight []float64, balloonRadius []float64, balloonSurface []float64, balloonTension []float64) {
	for i := 0; i < time; i++ {
		pts[i].X = float64(i)
		pts[i].Y = balloonVolume[i]
		wts[i].X = float64(i)
		wts[i].Y = balloonWeight[i]
		dts[i].X = float64(i)
		dts[i].Y = balloonRadius[i]
		qts[i].X = float64(i)
		qts[i].Y = balloonSurface[i]
		tts[i].X = float64(i)
		tts[i].Y = balloonTension[i]
	}
}

func savePlots(p *plot.Plot, w *plot.Plot, d *plot.Plot, q *plot.Plot, t *plot.Plot) {
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
	if err := d.Save(6*vg.Inch, 4*vg.Inch, "balloonRadius_simulation.png"); err != nil {
		fmt.Println("Error during saving the plot:", err)
	} else {
		fmt.Println("Plot has been saved to the file 'balloonRadius_simulation.png'.")
	}

	if err := q.Save(6*vg.Inch, 4*vg.Inch, "balloonSurface_simulation.png"); err != nil {
		fmt.Println("Error during saving the plot:", err)
	} else {
		fmt.Println("Plot has been saved to the file 'balloonSurface_simulation.png'.")
	}

	if err := t.Save(6*vg.Inch, 4*vg.Inch, "balloonSurfaceTension_simulation.png"); err != nil {
		fmt.Println("Error during saving the plot:", err)
	} else {
		fmt.Println("Plot has been saved to the file 'balloonSurfaceTension_simulation.png'.")
	}
}
