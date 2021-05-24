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

	return player
}

// TODO GUN COMPONENT
// //OTHER ACTIONS
// func (p *Player) Shoot() {
// 	fmt.Println("Pew Pew")
// 	// TODO this will be handles by Ship Object / Gun Object
// 	if time.Since(p.lastShot) >= playerShotCoolDown {
// 		//firing 2 bullets
// 		if bul, ok := bulletFromPool(); ok {
// 			bul.Fire(p.x, p.y)
// 			fmt.Println(bul)

// 		}
// 		if bul, ok := bulletFromPool(); ok {
// 			bul.Fire(p.x-25, p.y)
// 			fmt.Println(bul)
// 		}
// 		p.lastShot = time.Now()
// 	}

// }

// // Draws the Player at the current state
// TODO PLAYER SKIN COMPONENT
// func (p *Player) Draw(screen *ebiten.Image) {
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(-float64(p.width)/2, -float64(p.height)/2)
// 	op.GeoM.Translate(p.x, p.y)
// 	img := createImage("./assets/player.png")
// 	screen.DrawImage(img, op)
// }

//TODO  PlayerKeyboardController Component
// //  Listens for state Changes
// func (p *Player) Update() {
// 	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
// 		p.MoveLeft()
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
// 		p.MoveRight()
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeySpace) {
// 		p.Shoot()
// 	}
// }
