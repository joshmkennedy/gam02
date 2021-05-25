package main

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	bulletSpeed float64 = 5
	bulletCount         = 20
)

func newBullet() *Object {
	bullet := &Object{}
	bullet.isActive = false
	bullet.position.y = windowHeight - 100
	bullet.texture = createImage("./assets/bullet.png")
	bullet.tag = "BULLET"
	bullet.width = 4
	bullet.height = 7

	bulletMover := newBulletMover(bullet)
	bullet.addComponent(bulletMover)
	return bullet

}

type BulletMover struct {
	parent      *Object
	bulletSpeed float64
}

func newBulletMover(parent *Object) *BulletMover {
	return &BulletMover{
		parent:      parent,
		bulletSpeed: bulletSpeed,
	}
}

// ACTIONS

func (b *BulletMover) Deactivate() {
	b.parent.isActive = false
	b.parent.position.y = windowHeight - 100
}

func (b *BulletMover) Fire(x float64) {
	b.parent.isActive = true
	b.parent.position.x = x
}

// BULLET POOL FUNCTIONS

var bulletPool []*Object

func initBulletPool() {
	for i := 0; i < bulletCount; i++ {
		bul := newBullet()
		bulletPool = append(bulletPool, bul)
		Objects = append(Objects, bul)
	}
}

func bulletFromPool() (*Object, bool) {
	for _, bul := range bulletPool {
		if !bul.isActive {
			return bul, true
		}
	}

	return nil, false
}

// Draws Current State to Game
func (b *BulletMover) OnDraw(screen *ebiten.Image) error {
	return nil
}
func (b *BulletMover) OnCollision(other *Object) error {
	return nil
}

// Updates State
func (b *BulletMover) OnUpdate() error {
	if !b.parent.isActive {
		return nil
	}
	b.parent.position.y -= bulletSpeed
	if b.parent.position.y <= 0 || b.parent.position.y >= windowHeight {
		b.Deactivate()
	}
	return nil
}
