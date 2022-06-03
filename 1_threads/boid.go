package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (b *Boid) moveOne() {
	b.position = b.position.Add(b.velocity)
	next := b.position.Add(b.velocity)
	if next.X >= screenWidth || next.X < 0 {
		b.velocity = Vector2D{-b.velocity.X, b.velocity.Y}
	}

	if next.Y >= screenHeight || next.Y < 0 {
		b.velocity = Vector2D{b.velocity.X, -b.velocity.Y}
	}
}

func createBoid(id int) {
	b := Boid{
		position: Vector2D{
			X: rand.Float64() * screenWidth,
			Y: rand.Float64() * screenHeight,
		},
		velocity: Vector2D{
			X: (rand.Float64() * 2) - 1.0,
			Y: (rand.Float64() * 2) - 1.0,
		},
		id: id,
	}
	boids[id] = &b
	go b.start()
}
