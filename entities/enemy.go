package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	Pos   rl.Vector2
	Speed float32
	Size  float32
	HP    int32
	MaxHP int32
}

func NewEnemy(x, y float32) *Enemy {
	return &Enemy{
		Pos:   rl.NewVector2(x, y),
		Speed: 1.5,
		Size:  20,
		HP:    5,
		MaxHP: 5,
	}
}

func (e *Enemy) Update(player *Player) {
	dir := rl.Vector2Subtract(player.Pos, e.Pos)
	if rl.Vector2Length(dir) > 1 {
		dir = rl.Vector2Normalize(dir)
		e.Pos.X += dir.X * e.Speed
		e.Pos.Y += dir.Y * e.Speed
	}
}

func (e *Enemy) Draw() {
	rl.DrawRectangleV(e.Pos, rl.NewVector2(e.Size, e.Size), rl.Blue)
}

func (e *Enemy) Rect() rl.Rectangle {
	return rl.Rectangle{
		X:      e.Pos.X,
		Y:      e.Pos.Y,
		Width:  e.Size,
		Height: e.Size,
	}
}
