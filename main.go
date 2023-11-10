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

	rand.Seed(time.Now().UnixNano())

	// latexBalon := Balon{
	//Volume: 15 ,
	//Material: "Latex" ,
	//}

	// Inicjalizacja balona
	var radius float64 = 200
	var balloon []float64
	balloon = append(balloon, 200)
	var time int = 0

	// Tworzenie wykresu
	p := plot.New()

	// Symulacja pompowania balona
	for radius < 15000 {
		// Symulacja rozciągliwości balona
		time += 1
		if time%3 == 0 {
			radius = radius + 1000 + 3000*rand.Float64()
			balloon = append(balloon, radius)
			fmt.Println("if")
		} else {
			balloon = append(balloon, radius)
			fmt.Println("else")
		}
		fmt.Println(balloon[time])
		fmt.Println(time)

	}

	// Przygotowanie danych
	pts := make(plotter.XYs, time)
	for i := 0; i < time; i++ {
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
	p.X.Label.Text = "Czas [s]"
	p.Y.Label.Text = "Objętość balona [mL]"
	p.X.Min = 0
	p.Y.Min = 0

	// Zapisanie wykresu do pliku
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "balloon_simulation.png"); err != nil {
		fmt.Println("Błąd podczas zapisywania wykresu:", err)
	} else {
		fmt.Println("Wykres został zapisany do pliku 'balloon_simulation.png'.")
	}
}
