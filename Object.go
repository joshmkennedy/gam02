package main

import "github.com/hajimehoshi/ebiten"

type Object struct {
	x, y          float64
	width, height float64
	tag           string
	isActive      bool
	components    []*Component
	texture       *ebiten.Image
}
