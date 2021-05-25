package main

import (
	"time"
)

const (
	playerWidth        = 60
	playerHeight       = 47
	playerSpeed        = 7.8
	playerShotCoolDown = time.Millisecond * 250
)

//the player
// type Player struct {
// 	x, y          float64
// 	width, height float64
// 	lastShot      time.Time
// }

func newPlayer() *Object {
	player := &Object{}
	player.position = Vector{
		x: float64(windowWidth)/2 - float64(playerWidth)/2,
		y: float64(windowHeight - playerHeight),
	}
	player.width = playerWidth
	player.height = playerHeight

	player.tag = "PLAYER"

	player.isActive = true
	player.texture = createImage("./assets/player.png")

	keyboardMover := newKeyboardMover(player, playerSpeed)
	player.addComponent(keyboardMover)

	keyboardShooter := newKeyboardShooter(player, playerShotCoolDown)
	player.addComponent((keyboardShooter))

	return player
}
