package main

import (
	"fmt"
	"time"
)

const (
	playerWidth        = 60
	playerHeight       = 47
	playerSpeed        = 7.8
	playerShotCoolDown = time.Millisecond * 250
)

//the player
type Player struct {
	x, y          float64
	width, height float64
	lastShot      time.Time
}

func (p *Player) MoveLeft() {
	if p.x+float64(p.width) <= 0 {
		p.x = windowWidth
	} else {
		p.x -= playerSpeed
	}
}
func (p *Player) MoveRight() {
	if p.x > windowWidth {
		p.x = 0
	} else {
		p.x += playerSpeed
	}
}

func (p *Player) Shoot() {
	fmt.Println("Pew Pew")
	if time.Since(p.lastShot) >= playerShotCoolDown {
		//firing 2 bullets
		if bul, ok := bulletFromPool(); ok {
			bul.Fire(p.x, p.y)
			fmt.Println(bul)

		}
		if bul, ok := bulletFromPool(); ok {
			bul.Fire(p.x-25, p.y)
			fmt.Println(bul)
		}
		p.lastShot = time.Now()
	}

}
