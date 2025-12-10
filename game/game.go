package game

import (
	"CalmMenace/entities"
	"CalmMenace/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var fireCooldown float32 = 0

type Game struct {
	Player      *entities.Player
	Enemies     []*entities.Enemy
	Room        *world.Room
	Projectiles []*entities.Projectile
}

func New() *Game {
	return &Game{
		Player: entities.NewPlayer(),
		Room:   world.NewRoom(),
	}
}

func (g *Game) Update() {
	g.Player.Update()
	g.Room.Update(g.Player)

	fireCooldown -= rl.GetFrameTime()

	dir := rl.Vector2{X: 0, Y: 0}

	// FIRE ONCE APPROACH
	if rl.IsKeyPressed(rl.KeyLeft) {
		dir.X = -1
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		dir.X = 1
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		dir.Y = -1
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		dir.Y = 1
	}

	// Fire once per key press
	if (dir.X != 0 || dir.Y != 0) && fireCooldown <= 0 {
		p := entities.NewProjectile(g.Player.Pos, dir)
		g.Projectiles = append(g.Projectiles, p)
		fireCooldown = PlayerProjectileSpeed
	}

	// CONTINUOUS FIRE APPROACH
	holdDir := rl.Vector2{X: 0, Y: 0}

	if rl.IsKeyDown(rl.KeyLeft) {
		holdDir.X = -1
	}
	if rl.IsKeyDown(rl.KeyRight) {
		holdDir.X = 1
	}
	if rl.IsKeyDown(rl.KeyUp) {
		holdDir.Y = -1
	}
	if rl.IsKeyDown(rl.KeyDown) {
		holdDir.Y = 1
	}

	if (holdDir.X != 0 || holdDir.Y != 0) && fireCooldown <= 0 {
		p := entities.NewProjectile(g.Player.Pos, holdDir)
		g.Projectiles = append(g.Projectiles, p)
		fireCooldown = PlayerProjectileSpeed
	}

	for i := range g.Projectiles {
		if g.Projectiles[i].Active {
			g.Projectiles[i].Update()
		}
	}
	g.UpdateProjectiles()

	g.Room.RemoveDeadEnemies()
	g.RemoveInactiveProjectiles()
}

func (g *Game) Draw() {
	g.Room.Draw()
	g.Player.Draw()
	drawPlayerHealth(g.Player)

	for _, pr := range g.Projectiles {
		pr.Draw()
	}
}

func (g *Game) RemoveInactiveProjectiles() {
	var active []*entities.Projectile
	for _, p := range g.Projectiles {
		if p.Active {
			active = append(active, p)
		}
	}
	g.Projectiles = active
}

func (g *Game) UpdateProjectiles() {
	for i := range g.Projectiles {
		p := &g.Projectiles[i]

		if !(*p).Active {
			continue
		}

		g.handleCollision(*p)
	}
}

func (g *Game) handleCollision(p *entities.Projectile) {
	for i := range g.Room.Enemies {
		enemy := &g.Room.Enemies[i]

		if (*enemy).HP <= 0 {
			continue
		}

		if !rl.CheckCollisionRecs(p.Rect(), (*enemy).Rect()) {
			continue
		}

		applyDamage(*enemy, p)
		break // projectile hits only one enemy
	}
}

func applyDamage(enemy *entities.Enemy, p *entities.Projectile) {
	enemy.HP -= p.Damage
	p.Active = false
}

func drawPlayerHealth(p *entities.Player) {
	barWidth := float32(200)
	barHeight := float32(20)

	x := float32(20)
	y := float32(20)

	pct := float32(p.HP) / float32(p.MaxHP)

	rl.DrawRectangle(int32(x), int32(y), int32(barWidth), int32(barHeight), rl.DarkGray)
	rl.DrawRectangle(int32(x), int32(y), int32(barWidth*pct), int32(barHeight), rl.Green)
}
