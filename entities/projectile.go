package entities

import (
	"CalmMenace/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Projectile struct {
	Pos    rl.Vector2
	Dir    rl.Vector2
	Speed  float32
	Size   float32
	Damage int32
	Active bool
}

func NewProjectile(start rl.Vector2, dir rl.Vector2) *Projectile {
	if rl.Vector2Length(dir) == 0 {
		dir = rl.NewVector2(0, -1)
	}

	return &Projectile{
		Pos:    start,
		Dir:    rl.Vector2Normalize(dir),
		Speed:  8,
		Size:   6,
		Damage: 1,
		Active: true,
	}
}

func (p *Projectile) Update() {
	if !p.Active {
		return
	}

	p.Pos.X += p.Dir.X * p.Speed
	p.Pos.Y += p.Dir.Y * p.Speed

	if p.Pos.X < 0 || p.Pos.X > config.ScreenWidth ||
		p.Pos.Y < 0 || p.Pos.Y > config.ScreenHeight {
		p.Active = false
	}
}

func (p *Projectile) Draw() {
	rl.DrawRectangleV(p.Pos, rl.NewVector2(p.Size, p.Size), rl.Black)
}

func (p *Projectile) Rect() rl.Rectangle {
	return rl.Rectangle{
		X:      p.Pos.X,
		Y:      p.Pos.Y,
		Width:  p.Size,
		Height: p.Size,
	}
}
