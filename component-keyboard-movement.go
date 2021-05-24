package main

import "github.com/hajimehoshi/ebiten"

type KeyboardMover struct {
	parent *Object
	speed  float64
}

func newKeyboardMover(o *Object, speed float64) *KeyboardMover {
	return &KeyboardMover{
		parent: o,
		speed:  speed,
	}
}

func (mover *KeyboardMover) OnDraw(screen *ebiten.Image) error {
	return nil
}
func (mover *KeyboardMover) OnUpdate() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		mover.MoveLeft(mover.parent)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		mover.MoveRight(mover.parent)
	}

	return nil
}

func (mover *KeyboardMover) OnCollision(o *Object) error {
	return nil
}

func (mover *KeyboardMover) MoveLeft(parent *Object) {
	if parent.position.x+float64(parent.width) <= 0 {
		parent.position.x = windowWidth
	} else {
		parent.position.x -= mover.speed
	}
}
func (mover *KeyboardMover) MoveRight(parent *Object) {
	if parent.position.x > windowWidth {
		parent.position.x = 0
	} else {
		parent.position.x += mover.speed
	}
}
