package main

import (
	p "github.com/adrian-petrov/go-game/player"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
	playerSprite = "assets/characters/BasicCharakterSpritesheet.png"
)

var (
	running          = true
	backgroundColour = rl.NewColor(147, 211, 196, 255)

	grassSprite rl.Texture2D

	player *p.Player

	gameFrames int

	musicPaused bool
	music       rl.Music

	cam rl.Camera2D
)

func main() {
	for running {
		input()
		update()
		render()
	}
	quit()
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Sproutlings")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	player = p.NewPlayer(playerSprite)

	// TODO: to be refactored
	initTextures()
	initCamera()
	initMusic()
}

func input() {
	player.HandleInput()

	if rl.IsKeyPressed(rl.KeyP) {
		musicPaused = !musicPaused
	}
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColour)
	rl.BeginMode2D(cam)
	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func update() {
	running = !rl.WindowShouldClose()
	updateMusic()
	updateCamera()
	player.UpdatePosition(gameFrames)
	gameFrames++
}

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	player.Draw()
}

func quit() {
	rl.UnloadTexture(grassSprite)
	player.Dispose()
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func initTextures() {
	grassSprite = rl.LoadTexture("assets/tilesets/Grass.png")
}

func initMusic() {
	rl.InitAudioDevice()
	music = rl.LoadMusicStream("assets/music/AverysFarmLoopable.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)
}

func updateMusic() {
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}
}

func initCamera() {
	x, y, width, height :=
		player.DestinationRectX(),
		player.DestinationRectY(),
		player.DestinationRectWidth(),
		player.DestinationRectHeight()

	cam = rl.NewCamera2D(
		rl.NewVector2(
			float32(screenWidth/2),
			float32(screenHeight/2)),
		rl.NewVector2(
			float32(x-(width/2)),
			float32(y-(height/2))),
		0,
		1.0)
}

func updateCamera() {
	x, y, width, height :=
		player.DestinationRectX(),
		player.DestinationRectY(),
		player.DestinationRectWidth(),
		player.DestinationRectHeight()

	cam.Target = rl.NewVector2(
		float32(x-(width/2)),
		float32(y-(height/2)))
}
