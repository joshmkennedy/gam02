package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	windowWidth  = 560
	windowHeight = 740
	LEVEL        = "LEVEL"
	MENU         = "MENU"
)

// Game implements ebiten.Game interface.
type Game struct {
	activeScreen string
	level        *Level
	// menu *Menu //TODO
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if g.activeScreen == LEVEL {
		g.level.Update()
	}
	//TODO menu update

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	if g.activeScreen == LEVEL {
		g.level.Draw(screen)
	}
	//TODO menu draw
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func NewGame() *Game {
	return &Game{
		activeScreen: LEVEL,
		level:        newLevel(),
	}
}

func main() {

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Your game's title")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(NewGame()); err != nil {
		fmt.Println("cant RunGame")
	}
}

// TODO: Move to utils
func createImage(path string) *ebiten.Image {
	imageFromFile, err := os.Open(path)
	if err != nil {
		fmt.Println("ERROR: opening ", path)
	}
	defer imageFromFile.Close()

	pngImage, err := png.Decode(imageFromFile)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(pngImage)
}
