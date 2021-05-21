package main

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	enemyWidth  = 25
	enemyHeight = 25
	enemySpeed  = 7.8

	enemyCount = 6
)

//the enemy
type Enemy struct {
	x, y          float64
	width, height float64
	active        bool
}

func newEnemy(x, y float64) *Enemy {
	return &Enemy{
		x:      x,
		y:      y,
		width:  enemyWidth,
		height: enemyHeight,
		active: true,
	}
}

//MOVEMENT

//OTHER ACTIONS

func (e *Enemy) Collide() {
	e.active = false
}

// Draws the Enemy at the current state
func (e *Enemy) Draw(screen *ebiten.Image) {
	if !e.active {
		return
	}

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(e.width)/2, float64(e.height)/2)
	op.GeoM.Translate(e.x, e.y)

	img := createImage("./assets/enemy.png")
	screen.DrawImage(img, op)
}

//  Listens for state Changes
func (e *Enemy) Update() {
	if !e.active {
		return
	}

}

var enemies []*Enemy

func initEnemies() {
	for i := 0; i <= enemyCount; i++ {
		enemy := newEnemy(float64(i)*enemyWidth+(windowWidth/3), float64(i%2)*playerHeight)
		enemies = append(enemies, enemy)
	}
}
