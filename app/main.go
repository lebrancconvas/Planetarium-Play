package main

import (
	_ "image/png"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	err error
	bg *ebiten.Image 
	logo *ebiten.Image 
)

func init() {
	bg, _, err = ebitenutil.NewImageFromFile("../assets/img/space_bg07.png");
	if err != nil {
		log.Fatal(err); 
	}

	logo, _, err = ebitenutil.NewImageFromFile("../assets/img/logo.png");
	if err != nil {
		log.Fatal(err); 
	}
}

type Game struct{

}

func (g *Game) Update() error {
	return nil;
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil);
	screen.DrawImage(logo, nil);
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480;
}

func main() {
	ebiten.SetWindowSize(640, 480);
	ebiten.SetWindowTitle("Planetarium");
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err);
	}
}