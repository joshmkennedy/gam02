package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type KeyboardShooter struct {
	parent   *Object
	cooldown time.Duration
	lastShot time.Time
}

func newKeyboardShooter(o *Object, cooldown time.Duration) *KeyboardShooter {
	return &KeyboardShooter{
		parent:   o,
		cooldown: cooldown,
	}
}

func (shooter *KeyboardShooter) OnUpdate() error {

	return nil
}
func (shooter *KeyboardShooter) OnDraw(screen *ebiten.Image) error {
	return nil
}
func (shooter *KeyboardShooter) OnCollision(other *Object) error {

	return nil
}

func (shooter *KeyboardShooter) Shoot() error {
	fmt.Println("Pew Pew")
	// 	// TODO this will be handles by Ship Object / Gun Object
	if time.Since(shooter.lastShot) >= shooter.cooldown {
		//firing 2 bullets
		if bul, ok := bulletFromPool(); ok {
			bul.Fire(shooter.parent.position.x, shooter.parent.position.y)
			fmt.Println(bul)

		}
		if bul, ok := bulletFromPool(); ok {
			bul.Fire(shooter.parent.position.x-25, shooter.parent.position.y)
			fmt.Println(bul)
		}
		shooter.lastShot = time.Now()
	}
	return nil
}
