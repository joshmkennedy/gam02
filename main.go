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
	// Write your game's logical update.
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.MoveLeft()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.MoveRight()
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.player.Shoot()
	}

	for _, bul := range bulletPool {
		bul.Update()
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(g.player.width)/2, -float64(g.player.height)/2)
	op.GeoM.Translate(g.player.x, g.player.y)
	img := createImage("./assets/player.png")
	screen.DrawImage(img, op)

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
	initBulletPool()
	return &Game{
		player: &Player{
			x:      float64(windowWidth)/2 - float64(playerWidth)/2,
			y:      float64(windowHeight - playerHeight),
			width:  playerWidth,
			height: playerHeight,
		},
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
