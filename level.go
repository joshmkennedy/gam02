package main

import "github.com/hajimehoshi/ebiten"

type Level struct {
	activeState string
	objects     []*Object
}

func newLevel() *Level {
	return &Level{}
}

func (l *Level) Update() {

}

func (l *Level) Draw(screen *ebiten.Image) {

}
