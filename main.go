package main

import (
	"fmt"
	"game/models"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func load(filePath string) image.Image {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		fmt.Println("Cannot read file:", err)
	}

	imgData, err := png.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}
	return imgData.(image.Image)
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Tom and Jerry game")

	background := load("./assets/tom-jerry.png")
	tomSprites := load("./assets/tom-sprite.png")
	jerrySprites := load("./assets/jerry-sprite.png")

	now := time.Now().UnixMilli()
	game := models.NewGame(
		800,
		500,
		10,
		now,
		1,
	)

	fpsInterval := int64(1000 / game.GetFps())

	player := models.NewTom(100, 200, 40, 72, 0, 0, 4, 3, 0, 1, 2, 25, 0, 0)
	jerry := models.NewJerry(200, 200, 40, 72, 0, 0, 4, 3, 0, 1, 2, 25, 0, 0)

	img := canvas.NewImageFromImage(background)
	img.FillMode = canvas.ImageFillOriginal

	sprite := image.NewRGBA(background.Bounds())

	playerImg := canvas.NewRasterFromImage(sprite)
	jerryImg := canvas.NewRasterFromImage(sprite)
	spriteSize := image.Pt(player.GetWidth(), player.GetHeight())

	c := container.New(layout.NewBorderLayout(nil, nil, nil, nil), img, playerImg, jerryImg)

	// Variables para el cronómetro y el marcador
	var (
		startTime   time.Time
		score       int
		timerLabel  = canvas.NewText("", color.White)
		scoreLabel  = canvas.NewText("Score: 0", color.White)
	)

	// Establece el estilo de texto para el cronómetro como negritas
	timerLabel.TextStyle = fyne.TextStyle{Bold: true}
	scoreLabel.TextStyle = fyne.TextStyle{Bold: true}

	go func() {
		startTime = time.Now()

		for {
			time.Sleep(time.Millisecond)

			now := time.Now().UnixMilli()
			elapsed := now - game.GetThen()

			if elapsed > fpsInterval {
				game.SetThen(now)

				draw.Draw(sprite, sprite.Bounds(), image.Transparent, image.ZP, draw.Src)

				spriteDP := image.Pt(player.GetWidth()*player.GetFrameX(), player.GetHeight()*player.GetFrameY())
				sr := image.Rectangle{spriteDP, spriteDP.Add(spriteSize)}

				dp := image.Pt(player.GetX(), player.GetY())
				r := image.Rectangle{dp, dp.Add(spriteSize)}

				draw.Draw(sprite, r, tomSprites, sr.Min, draw.Over)
				playerImg = canvas.NewRasterFromImage(sprite)

				spriteDPJerry := image.Pt(jerry.GetWidth()*jerry.GetFrameX(), jerry.GetHeight()*jerry.GetFrameY())
				srJerry := image.Rectangle{spriteDPJerry, spriteDPJerry.Add(spriteSize)}

				dpJerry := image.Pt(jerry.GetX(), jerry.GetY())
				rJerry := image.Rectangle{dpJerry, dpJerry.Add(spriteSize)}

				if !r.Overlaps(rJerry) {
					draw.Draw(sprite, rJerry, jerrySprites, srJerry.Min, draw.Over)
					jerryImg = canvas.NewRasterFromImage(sprite)
				} else {
					randX := rand.Intn(400) + 100
					randY := rand.Intn(121) + 100
					jerry.SetX(randX)
					jerry.SetY(randY)

					// Incrementa el marcador cuando atrapas a Jerry
					score++
					scoreLabel.Text = "Score: " + strconv.Itoa(score)
				}

				if player.GetXMov() != 0 || player.GetYMov() != 0 {
					player.SetX(player.GetX() + player.GetXMov())
					player.SetY(player.GetY() + player.GetYMov())
					player.SetFrameX((player.GetFrameX() + 1) % player.GetCyclesX())
					player.SetXMov(0)
					player.SetYMov(0)
				} else {
					player.SetFrameX(0)
				}

				// Actualiza el cronómetro
				elapsedTime := time.Since(startTime)
				timerLabel.Text = "Time: " + elapsedTime.Round(time.Second).String()

				c.Refresh()
			}
		}
	}()

	infoContainer := container.NewVBox(timerLabel, scoreLabel)
	c.Add(infoContainer)

	w.SetContent(c)

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyDown:
			if player.GetY() < int(game.GetHeight())-player.GetHeight()-game.GetMargin() {
				player.SetYMov(player.GetSpeed())
			}
			player.SetFrameY(player.GetDownY())
		case fyne.KeyUp:
			if player.GetY() > 100 {
				player.SetYMov(-(player.GetSpeed()))
			}
			player.SetFrameY(player.GetUpY())
		case fyne.KeyLeft:
			if player.GetX() > game.GetMargin() {
				player.SetXMov(-(player.GetSpeed()))
			}
			player.SetFrameY(player.GetLeftY())
		case fyne.KeyRight:
			if player.GetX() < int(game.GetWidth())-(player.GetWidth())-game.GetMargin() {
				player.SetXMov(player.GetSpeed())
			}
			player.SetFrameY(player.GetRightY())
		}
	})

	w.CenterOnScreen()
	w.ShowAndRun()
}
