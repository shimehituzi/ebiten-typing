package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestGame_Update(t *testing.T) {
	g := &Game{}
	err := g.Update()
	if err != nil {
		t.Errorf("Update() error = %v", err)
	}
}

func TestDraw(t *testing.T) {
	g := &Game{}
	screen := ebiten.NewImage(200, 200)

	g.Draw(screen)
}
func TestGame_Layout(t *testing.T) {
	g := &Game{}
	w, h := g.Layout(640, 480)
	if w != 320 || h != 240 {
		t.Errorf("Layout() = %v, %v, want %v, %v", w, h, 320, 240)
	}
}
