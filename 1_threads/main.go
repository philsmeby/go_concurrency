package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
)

var (
	green = color.RGBA{10, 255, 50, 255}
	boids [boidCount]*Boid
)

type BoidGame struct{}

func (g *BoidGame) Update() error {
	return nil
}

func (g *BoidGame) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.X+1), int(boid.position.Y), green)
		screen.Set(int(boid.position.X-1), int(boid.position.Y), green)
		screen.Set(int(boid.position.X), int(boid.position.Y-1), green)
		screen.Set(int(boid.position.X), int(boid.position.Y+1), green)
	}
}

func (g *BoidGame) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")
	if err := ebiten.RunGame(&BoidGame{}); err != nil {
		log.Fatal(err)
	}
}
