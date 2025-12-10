package world

import (
	"CalmMenace/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Room struct {
	Enemies []*entities.Enemy
}

func NewRoom() *Room {
	return &Room{
		Enemies: []*entities.Enemy{
			entities.NewEnemy(100, 100),
			entities.NewEnemy(600, 300),
		},
	}
}

func (r *Room) Update(player *entities.Player) {
	for _, e := range r.Enemies {
		e.Update(player)
	}
}

func (r *Room) Draw() {
	for _, e := range r.Enemies {
		e.Draw()
		drawHealthBarEnemy(e)
	}
}

func (r *Room) RemoveDeadEnemies() {
	var alive []*entities.Enemy
	for _, e := range r.Enemies {
		if e.HP > 0 {
			alive = append(alive, e)
		}
	}
	r.Enemies = alive
}

func drawHealthBarEnemy(e *entities.Enemy) {
	barWidth := e.Size
	barHeight := float32(4)

	x := e.Pos.X
	y := e.Pos.Y - barHeight - 2

	pct := float32(e.HP) / float32(e.MaxHP)

	rl.DrawRectangle(int32(x), int32(y), int32(barWidth), int32(barHeight), rl.DarkGray)
	rl.DrawRectangle(int32(x), int32(y), int32(barWidth*pct), int32(barHeight), rl.Green)
}
