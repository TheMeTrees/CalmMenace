package world

import (
	"CalmMenace/entities"
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
