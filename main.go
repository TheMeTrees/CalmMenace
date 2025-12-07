package main

import (
	"CalmMenace/config"
	"CalmMenace/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Isaac-like")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	g := game.New()

	for !rl.WindowShouldClose() {
		g.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		g.Draw()
		rl.EndDrawing()
	}
}
