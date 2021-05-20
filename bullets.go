package main

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	bulletSpeed float64 = 5
	bulletCount         = 20
)

type Bullet struct {
	x, y   float64
	active bool
	op     *ebiten.DrawImageOptions
}

func newBullet() *Bullet {
	return &Bullet{
		active: false,
		y:      100,
		op:     &ebiten.DrawImageOptions{},
	}
}

// ACTIONS

func (b *Bullet) Deactivate() {
	b.active = false
	b.y = 100

}

func (b *Bullet) Fire(x, y float64) {
	b.active = true
	b.x = x
}

// BULLET POOL FUNCTIONS

var bulletPool []*Bullet

func initBulletPool() {
	for i := 0; i < bulletCount; i++ {
		bul := newBullet()
		bulletPool = append(bulletPool, bul)
	}
}

func bulletFromPool() (*Bullet, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}

// Draws Current State to Game
func (b *Bullet) Draw(screen *ebiten.Image) {
	if !b.active {
		return
	}
	bop := &ebiten.DrawImageOptions{}

	bop.GeoM.Translate(4, 7)                      //sets origin center
	bop.GeoM.Translate(b.x, (windowHeight - b.y)) //move by this amount
	bimg := createImage("./assets/bullet.png")
	screen.DrawImage(bimg, bop)
}

// Updates State
func (b *Bullet) Update() {
	if !b.active {
		return
	}
	b.y += bulletSpeed
	if b.y <= 0 || b.y >= windowHeight {
		b.Deactivate()
	}
}
