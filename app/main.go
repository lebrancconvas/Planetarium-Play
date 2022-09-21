package main

import (
	_ "image/png"
	"bytes"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

var (
	err error
	bg *ebiten.Image 
	logo *ebiten.Image 
	startButton *ebiten.Image 
	audioContext *audio.Context 
	// sound *audio.Player
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

	startButton, _, err = ebitenutil.NewImageFromFile("../assets/img/1.png");
	if err != nil {
		log.Fatal(err); 
	}

	sound, err := readAudio(audioContext, "../assets/audio/bgm/star_wars.mp3");
	if err != nil {
		log.Fatal(err); 
	}
}

type Game struct{
	
}

func readAudio(context *audio.Context, asset []byte) (*audio.Player, error) {
	mp3Decoded, err := mp3.Decode(context, bytes.NewReader(asset))
	if err != nil {
		return nil, err
	}
	player, err := audio.NewPlayer(context, mp3Decoded)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (g *Game) Update() error {
	// audioContext.NewPlayer();
	return nil;
}

func (g *Game) Draw(screen *ebiten.Image) {
	logoOp := &ebiten.DrawImageOptions{};
	logoOp.GeoM.Scale(0.75, 0.75);
	logoOp.GeoM.Translate(140, -30);

	startOp := &ebiten.DrawImageOptions{};
	startOp.GeoM.Scale(0.75, 0.75);
	startOp.GeoM.Translate(140, 100);

	screen.DrawImage(bg, nil);
	screen.DrawImage(logo, logoOp); 
	screen.DrawImage(startButton, startOp); 
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