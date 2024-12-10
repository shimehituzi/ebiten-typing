package main

import (
	"image"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const (
    screenWidth  = 640
    screenHeight = 480
)

type Game struct {
    text string
}

func (g *Game) Update() error {
    // 新たに押されたキーを取得
    pressedKeys := inpututil.AppendPressedKeys(nil)
    for _, k := range pressedKeys {
        if k == ebiten.KeyBackspace && len(g.text) > 0 {
            // バックスペースで最後の文字を削除
            g.text = g.text[:len(g.text)-1]
        } else if k == ebiten.KeyEnter {
            // エンターキーで改行を追加
            g.text += "\n"
        }
    }

    // 入力された文字を追加
    for _, char := range ebiten.InputChars() {
        g.text += string(char)
    }
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.White)
    drawText(screen, g.text, 20, 40, color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func drawText(screen *ebiten.Image, text string, x, y int, clr color.Color) {
    face := basicfont.Face7x13
    drawer := &font.Drawer{
        Dst:  screen,
        Src:  image.NewUniform(clr),
        Face: face,
        Dot:  fixed.P(x, y),
    }
    lines := strings.Split(text, "\n")
    for _, line := range lines {
        drawer.DrawString(line)
        drawer.Dot.Y += face.Metrics().Height
    }
}

func main() {
    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Text Input Example")
    if err := ebiten.RunGame(&Game{}); err != nil {
        panic(err)
    }
}
