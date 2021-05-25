package main

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	PAUSED  = "PAUSED"
	LOST    = "LOST"
	PLAYING = "PLAYING"
	WON     = "WON"
)

type Level struct {
	activeState string
	objects     []*Object
}

func newLevel() *Level {

	addAllObjects()

	return &Level{
		activeState: PLAYING,
		objects:     Objects,
	}
}

func (l *Level) Update() {

	for _, object := range Objects {
		object.Update()
	}
}

func (l *Level) Draw(screen *ebiten.Image) {

	for _, object := range Objects {
		object.Draw(screen)
	}
}

//Adds the level specific objects to the Objects Global variable
func addAllObjects() {
	addObject(newPlayer())

	// enemies
	initEnemies()

	//bullets
	initBulletPool()

}
