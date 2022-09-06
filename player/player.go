package player

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	defaultSpeed    = float32(3)
	defaultSrcRect  = rl.NewRectangle(0, 0, 48, 48)
	defaultDestRect = rl.NewRectangle(200, 200, 150, 150)
)

type Direction int

const (
	down Direction = iota
	up
	left
	right
)

type Player struct {
	sprite                                                rl.Texture2D
	speed                                                 float32
	srcRect                                               rl.Rectangle
	destRect                                              rl.Rectangle
	direction                                             Direction
	moving, movingUp, movingDown, movingRight, movingLeft bool
	frames                                                int
}

func NewPlayer(spriteFile string) *Player {
	return &Player{
		sprite:      rl.LoadTexture(spriteFile),
		speed:       defaultSpeed,
		srcRect:     defaultSrcRect,
		destRect:    defaultDestRect,
		direction:   down,
		moving:      false,
		movingUp:    false,
		movingDown:  false,
		movingRight: false,
		movingLeft:  false,
	}
}

func (p *Player) DestinationRectX() float32 {
	return p.destRect.X
}
func (p *Player) DestinationRectY() float32 {
	return p.destRect.Y
}
func (p *Player) DestinationRectWidth() float32 {
	return p.destRect.Width
}
func (p *Player) DestinationRectHeight() float32 {
	return p.destRect.Height
}

func (p *Player) HandleInput() {
	p.resetMovement()

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		p.moving = true
		p.movingUp = true
		p.direction = up
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		p.moving = true
		p.movingDown = true
		p.direction = down
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		p.moving = true
		p.movingLeft = true
		p.direction = left
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		p.moving = true
		p.movingRight = true
		p.direction = right
	}
}

func (p *Player) UpdatePosition(gameFrames int) {
	if !p.moving {
		// idle animation here
		p.incrementFramesIdling(gameFrames)
	} else {
		// increment player frames once every 8 gameFrames
		p.incrementFramesMoving(gameFrames)
	}
	// get the next animation image from sprite each time the frames increment
	p.updateSpriteImages()
	// update destinationRect coordinates
	p.updateDestCoordinates()
}

func (p *Player) Draw() {
	rl.DrawTexturePro(
		p.sprite,
		p.srcRect,
		p.destRect,
		rl.NewVector2(p.destRect.Width, p.destRect.Height),
		0,
		rl.White)
}

func (p *Player) Dispose() {
	rl.UnloadTexture(p.sprite)
}

func (p *Player) incrementFramesMoving(gameFrames int) {
	if gameFrames%8 == 1 {
		p.frames++
	}
	// set frames to 0 to loop back to default image
	if p.frames > 3 {
		p.frames = 0
	}
}

func (p *Player) incrementFramesIdling(gameFrames int) {
	if gameFrames%45 == 1 {
		p.frames++
	}
	// set frames to 0 to loop back to default image
	if p.frames > 1 {
		p.frames = 0
	}
}

// point the source rectangle to a new image in the sprite
func (p *Player) updateSpriteImages() {
	p.srcRect.X = p.srcRect.Width * float32(p.frames)
	p.srcRect.Y = p.srcRect.Height * float32(p.direction)
}

func (p *Player) updateDestCoordinates() {
	if p.movingUp {
		p.destRect.Y -= p.speed
	}
	if p.movingDown {
		p.destRect.Y += p.speed
	}
	if p.movingLeft {
		p.destRect.X -= p.speed
	}
	if p.movingRight {
		p.destRect.X += p.speed
	}
}

func (p *Player) resetMovement() {
	p.moving = false
	p.movingUp = false
	p.movingDown = false
	p.movingLeft = false
	p.movingRight = false
}
