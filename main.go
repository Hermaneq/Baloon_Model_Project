package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var avogadra float64 = math.Pow(10, 23) * 6.03

const weightUnit string = "gram"
const molVolume float64 = 22.4
const volumeUnit string = "Liter"
const temperature int = 0

// average human exhaust is 500ml
const exhaust int = 500
const exhaustPerMinute int = 15

type Balon struct {
	Volume   int
	Material string
}

func main() {

	// latexBalon := Balon{
	//Volume: 15 ,
	//Material: "Latex" ,
	//}

	// Parametry symulacji
	numSteps := 100
	inflationRate := 0.1 // Współczynnik rozciągliwości balona

	// Inicjalizacja balona
	radius := 1.0
	balloon := make([]float64, numSteps)
	balloon[0] = radius

	// Tworzenie wykresu
	p := plot.New()

	// Symulacja pompowania balona
	for step := 1; step < numSteps; step++ {
		// Symulacja rozciągliwości balona
		radius = radius * (1 + inflationRate*rand.Float64())
		balloon[step] = radius
	}

	// Przygotowanie danych
	pts := make(plotter.XYs, numSteps)
	for i := 0; i < numSteps; i++ {
		pts[i].X = float64(i)
		pts[i].Y = balloon[i]
	}

	// Dodawanie danych do wykresu
	line, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(line)

	// Ustawienia wykresu
	p.Title.Text = "Symulacja pompowania balona"
	p.X.Label.Text = "Kroki"
	p.Y.Label.Text = "Promień balona"
	p.X.Min = 0
	p.Y.Min = 0

	// Zapisanie wykresu do pliku
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "balloon_simulation.png"); err != nil {
		fmt.Println("Błąd podczas zapisywania wykresu:", err)
	} else {
		fmt.Println("Wykres został zapisany do pliku 'balloon_simulation.png'.")
	}
}
