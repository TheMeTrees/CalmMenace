package entities

import (
	"CalmMenace/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Pos   rl.Vector2
	Speed float32
	Size  float32
	Dir   rl.Vector2
	HP    int32
	MaxHP int32
}

func NewPlayer() *Player {
	return &Player{
		Pos:   rl.NewVector2(config.ScreenHeight/2, config.ScreenHeight/2),
		Speed: 4,
		Size:  20,
		HP:    6,
		MaxHP: 6,
	}
}

func (p *Player) Update() {
	p.Dir = rl.NewVector2(0, 0)

	if rl.IsKeyDown(rl.KeyW) {
		p.Pos.Y -= p.Speed
		p.Dir.Y = -1
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Pos.Y += p.Speed
		p.Dir.Y = 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.Pos.X -= p.Speed
		p.Dir.X = -1
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.Pos.X += p.Speed
		p.Dir.X = 1
	}

	if p.Pos.X < 0 {
		p.Pos.X = 0
	}
	if p.Pos.X+p.Size > config.ScreenWidth {
		p.Pos.X = config.ScreenWidth - p.Size
	}
	if p.Pos.Y < 0 {
		p.Pos.Y = 0
	}
	if p.Pos.Y+p.Size > config.ScreenHeight {
		p.Pos.Y = config.ScreenHeight - p.Size
	}
}

func (p *Player) Draw() {
	rl.DrawRectangleV(p.Pos, rl.NewVector2(p.Size, p.Size), rl.Red)
}
