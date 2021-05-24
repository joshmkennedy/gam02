package main

import (
	"fmt"
	"reflect"

	"github.com/hajimehoshi/ebiten"
)

type Component interface {
	OnUpdate() error
	OnDraw(screen *ebiten.Image) error
	OnCollision(other *Object) error
}

type Vector struct {
	x, y float64
}
type Circle struct {
	center Vector
	radius float64
}
type Object struct {
	position      Vector
	width, height float64
	tag           string
	isActive      bool
	components    []Component
	collisions    []Circle
	texture       *ebiten.Image
}

var Objects []*Object

func addObject(o *Object) {
	Objects = append(Objects, o)
}

func (o *Object) Draw(screen *ebiten.Image) error {
	// temporary
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(o.width)/2, -float64(o.height)/2)
	op.GeoM.Translate(o.position.x, o.position.y)
	screen.DrawImage(o.texture, op)

	for _, component := range o.components {
		err := component.OnDraw(screen)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *Object) Update() error {
	for _, component := range o.components {
		err := component.OnUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *Object) collision(other *Object) error {
	for _, component := range o.components {
		err := component.OnCollision(other)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *Object) addComponent(new Component) {
	for _, existing := range o.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	o.components = append(o.components, new)
}

func (o *Object) getComponent(withType Component) Component {
	typ := reflect.TypeOf(withType)
	for _, comp := range o.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	panic(fmt.Sprintf("no component with type %v", reflect.TypeOf(withType)))
}
