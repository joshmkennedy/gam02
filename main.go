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
)

// Game implements ebiten.Game interface.
type Game struct {
	count  int
	player *Player
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	//Updates Player Object
	g.player.Update()

	//Updates all Bullets
	for _, bul := range bulletPool {
		bul.Update()
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	//Drawing player Object
	g.player.Draw(screen)

	// Draws all the bullets
	for _, bul := range bulletPool {
		bul.Draw(screen)
	}

}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func NewGame() *Game {
	// Create all objects that wee need

	initBulletPool() // this just creates a global slice of Bullets

	return &Game{
		player: newPlayer(),
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
