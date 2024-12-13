package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	color2 "image/color"
	"math"
	"os"
)

const mapWidth = 24
const mapHeight = 24
const screenWidth = 1000
const screenHeight = 800
const framerate = 120

var world = [mapWidth][mapHeight]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 2, 0, 2, 0, 2, 0, 2, 0, 0, 0, 3, 2, 3, 1, 2, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
	{1, 0, 0, 0, 2, 0, 2, 0, 2, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 2, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 0, 0, 0, 5, 2, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 4, 0, 0, 0, 0, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 0, 4, 4, 4, 4, 4, 4, 3, 2, 1, 0, 0, 0, 0, 1, 2, 3, 4, 5, 0, 0, 1},
	{1, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

func main() {
	os.Exit(run())
}

func run() int {

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf(os.Stderr.Name(), "Failed to create window: %s\n", err)
		return 1
	}

	window, err := sdl.CreateWindow("go_raycasting", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf(os.Stderr.Name(), "Failed to create window: %s\n", err)
		return 1
	}
	defer func(window *sdl.Window) {
		err := window.Destroy()
		if err != nil {

		}
	}(window)

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	// game state

	var wPressed, sPressed, aPressed, dPressed bool

	var time /* current frame */, oldTime uint64 /* previous frame */

	// player position vector
	pPos := Vector2d[float64]{22.0, 12.0}
	// player view direction vector
	pDir := Vector2d[float64]{-1.0, 0}

	// every point in the direction of view is a sum of pPos and k*pDir where k is some real number

	// player camera plane vector
	// since this is a 2d rendering, this vector is not a real plane but a line
	// this one needs to be perpendicular to the pDir, otherwise all the graphics appear skewed
	cPlane := Vector2d[float64]{0.0, 0.66}

	running := true
	for running {

		// for every vertical stripe of the screen
		for x := 0; x < screenWidth; x++ {

			// the x-coordinate on the camera plane to the current x of the screen
			cameraX := 2*float64(x)/float64(screenWidth) - 1 // screenLeft = -1 >= cameraX <= 1 = screenRight

			rayDir := rayDirection(cameraX, pDir, cPlane)

			// since the map is stored as a 2-dimensional array, the corresponding position must be a vector of integers
			mapPos := Vector2d[int]{int(pPos.X), int(pPos.Y)}

			// distance the ray has to travel to go from one x-side to the next x-side or from one y-side to the next
			rayDistDelta := rayDistanceDelta(rayDir)

			//length of the ray from pPos to next x or y-side
			raySideDist := Vector2d[float64]{0, 0}

			// the DDA always jumps one square each loop, either in positive or negative x- or y-direction (1 or -1)
			var stepX, stepY int

			// calculate step and initial sideDist
			if rayDir.X < 0 {
				stepX = -1
				raySideDist.X = (pPos.X - float64(mapPos.X)) * rayDistDelta.X
			} else {
				stepX = 1
				raySideDist.X = (float64(mapPos.X) + 1.0 - pPos.X) * rayDistDelta.X
			}
			if rayDir.Y < 0 {
				stepY = -1
				raySideDist.Y = (pPos.Y - float64(mapPos.Y)) * rayDistDelta.Y
			} else {
				stepY = 1
				raySideDist.Y = (float64(mapPos.Y) + 1.0 - pPos.Y) * rayDistDelta.Y
			}

			hit := false
			side := false

			// actual DDA
			for hit == false {
				// jump to next map square, either in x-direction, or in y-direction
				if raySideDist.X < raySideDist.Y {
					raySideDist.X += rayDistDelta.X
					mapPos.X += stepX
					side = false
				} else {
					raySideDist.Y += rayDistDelta.Y
					mapPos.Y += stepY
					side = true
				}

				// Check if ray has hit a wall
				if world[mapPos.X][mapPos.Y] > 0 {
					hit = true
				}
			}

			// distance projected on camera direction. This is the shortest distance from the point where the wall is
			// perpendicular to the camera plane
			var cameraWallDist float64
			if side == false {
				cameraWallDist = raySideDist.X - rayDistDelta.X
			} else {
				cameraWallDist = raySideDist.Y - rayDistDelta.Y
			}

			// height of line to draw on screen
			lineHeight := int(float64(screenHeight) / cameraWallDist)

			// lowest and highest pixel to fill in current stripe
			drawStart := -lineHeight/2 + screenHeight/2
			if drawStart < 0 {
				drawStart = 0
			}
			drawEnd := lineHeight/2 + screenHeight/2
			if drawEnd >= screenHeight {
				drawEnd = screenHeight - 1
			}

			var color color2.RGBA
			switch world[mapPos.X][mapPos.Y] {
			case 1:
				color = color2.RGBA{R: 255, G: 0, B: 0, A: 255}
			case 2:
				color = color2.RGBA{R: 0, G: 255, B: 0, A: 255}
			case 3:
				color = color2.RGBA{R: 0, G: 0, B: 255, A: 255}
			case 4:
				color = color2.RGBA{R: 255, G: 255, B: 255, A: 255}
			default:
				color = color2.RGBA{R: 255, G: 255, B: 200, A: 255}
			}

			//give x and y sides different brightness
			if side == true {
				color = darkened(color)
			}

			for i := drawEnd; i >= drawStart; i-- {
				surface.Set(x, i, color)
			}
		}

		oldTime = time
		time = sdl.GetTicks64()
		var frameTime float64 = float64(time-oldTime) / 1000 //frameTime is the time this frame has taken, in seconds
		fmt.Println(1 / frameTime)                           //FPS counter

		window.UpdateSurface()

		sdl.Delay(uint32(1000/framerate - frameTime))

		surface.FillRect(nil, 0)

		//speed modifiers
		moveSpeed := float64(frameTime) * 5.0 //the constant value is in squares/second
		rotSpeed := float64(frameTime) * 3.0

		// input loop
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false

			case *sdl.KeyboardEvent:
				{
					k := event.(*sdl.KeyboardEvent)
					if k.Keysym.Scancode == sdl.SCANCODE_W {
						if k.State == sdl.PRESSED {
							wPressed = true
						} else {
							wPressed = false
						}
					}
					if k.Keysym.Scancode == sdl.SCANCODE_S {
						if k.State == sdl.PRESSED {
							sPressed = true
						} else {
							sPressed = false
						}
					}
					if k.Keysym.Scancode == sdl.SCANCODE_A {
						if k.State == sdl.PRESSED {
							aPressed = true
						} else {
							aPressed = false
						}
					}
					if k.Keysym.Scancode == sdl.SCANCODE_D {
						if k.State == sdl.PRESSED {
							dPressed = true
						} else {
							dPressed = false
						}
					}
				}
			}
		}

		if wPressed {
			if world[int(pPos.X+pDir.X*moveSpeed)][int(pPos.Y)] == 0 {
				pPos.X += pDir.X * moveSpeed
			}
			if world[int(pPos.X)][int(pPos.Y+pDir.Y*moveSpeed)] == 0 {
				pPos.Y += pDir.Y * moveSpeed
			}
		}
		if sPressed {
			if world[int(pPos.X-pDir.X*moveSpeed)][int(pPos.Y)] == 0 {
				pPos.X -= pDir.X * moveSpeed
			}
			if world[int(pPos.X)][int(pPos.Y-pDir.Y*moveSpeed)] == 0 {
				pPos.Y -= pDir.Y * moveSpeed
			}
		}
		if aPressed {
			//both camera direction and camera plane must be rotated
			oldDirX := pDir.X
			pDir.X = pDir.X*math.Cos(rotSpeed) - pDir.Y*math.Sin(rotSpeed)
			pDir.Y = oldDirX*math.Sin(rotSpeed) + pDir.Y*math.Cos(rotSpeed)
			oldPlaneX := cPlane.X
			cPlane.X = cPlane.X*math.Cos(rotSpeed) - cPlane.Y*math.Sin(rotSpeed)
			cPlane.Y = oldPlaneX*math.Sin(rotSpeed) + cPlane.Y*math.Cos(rotSpeed)
		}
		if dPressed {
			oldDirX := pDir.X
			pDir.X = pDir.X*math.Cos(-rotSpeed) - pDir.Y*math.Sin(-rotSpeed)
			pDir.Y = oldDirX*math.Sin(-rotSpeed) + pDir.Y*math.Cos(-rotSpeed)
			oldPlaneX := cPlane.X
			cPlane.X = cPlane.X*math.Cos(-rotSpeed) - cPlane.Y*math.Sin(-rotSpeed)
			cPlane.Y = oldPlaneX*math.Sin(-rotSpeed) + cPlane.Y*math.Cos(-rotSpeed)
		}

	}

	return 0
}

func rayDirection(cameraX float64, pDir, cPlane Vector2d[float64]) Vector2d[float64] {
	return Vector2d[float64]{pDir.X + cPlane.X*cameraX, pDir.Y + cPlane.Y*cameraX}
}

// relative length of ray from one x or y-side to next x or y-side
// for the DDA algorithm just the ratio between delX and delY matters, so instead of including the |rayDir| just 1 is used
func rayDistanceDelta(rayDir Vector2d[float64]) Vector2d[float64] {
	var deltaDistX float64
	if rayDir.X == 0 {
		deltaDistX = 1e30
	} else {
		deltaDistX = math.Abs(1 / rayDir.X)
	}
	var deltaDistY float64
	if rayDir.Y == 0 {
		deltaDistY = 1e30
	} else {
		deltaDistY = math.Abs(1 / rayDir.Y)
	}
	return Vector2d[float64]{deltaDistX, deltaDistY}
}

func darkened(c color2.RGBA) color2.RGBA {
	c.R = c.R / 2
	c.G = c.G / 2
	c.B = c.B / 2
	return c
}
