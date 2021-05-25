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

func newEnemy(x, y float64) *Object {
	enemy := &Object{}
	enemy.position = Vector{
		x: x,
		y: y,
	}
	enemy.width = enemyWidth
	enemy.height = enemyHeight
	enemy.isActive = true
	enemy.texture = createImage("./assets/enemy.png")
	enemy.tag = "ENEMY"
	return enemy
}

//MOVEMENT
type HitBox struct {
	parent *Object
}

func newHitBox(parent *Object) *HitBox {
	return &HitBox{
		parent: parent,
	}
}

//OTHER ACTIONS

func (hb *HitBox) OnCollision(other *Object) error {
	hb.parent.isActive = false
	return nil
}

// Draws the Enemy at the current state
func (hb *HitBox) OnDraw(screen *ebiten.Image) error {
	return nil
}

//  Listens for state Changes
func (hb *HitBox) OnUpdate() error {
	if !hb.parent.isActive {
		return nil
	}
	return nil
}

var enemies []*Object

func initEnemies() {
	for i := 0; i <= enemyCount; i++ {
		enemy := newEnemy(float64(i)*enemyWidth*3+(enemyWidth*2), float64(i%2)*playerHeight+playerHeight)
		enemies = append(enemies, enemy)
		Objects = append(Objects, enemy)
	}
}
