package main

import (
	"fmt"
	"game/models"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"math"
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
	spikeSprites := load("./assets/spike-sprite.png")

	now := time.Now().UnixMilli()
	game := models.NewGame(
		564,
		314,
		10,
		now,
		10,
	)

	// Ajusta el intervalo de actualización de la pantalla
	fpsInterval := time.Second / time.Duration(game.GetFps())

	player := models.NewTom(100, 200, 40, 72, 0, 0, 4, 3, 0, 1, 2, 25, 0, 0)
	jerry := models.NewJerry(200, 200, 40, 72, 0, 0, 4, 3, 0, 1, 2, 25, 0, 0)
	spike := models.NewSpike(300, 100, 40, 72, 0, 0, 4, 3, 0, 1, 2, 40, 0, 0)

	img := canvas.NewImageFromImage(background)
	img.FillMode = canvas.ImageFillOriginal

	sprite := image.NewRGBA(background.Bounds())

	playerImg := canvas.NewRasterFromImage(sprite)
	jerryImg := canvas.NewRasterFromImage(sprite)
	spikeImg := canvas.NewRasterFromImage(sprite)
	spriteSize := image.Pt(player.GetWidth(), player.GetHeight())

	c := container.New(layout.NewBorderLayout(nil, nil, nil, nil), img, playerImg, jerryImg, spikeImg)

	// Variables para el cronómetro y el marcador
	var (
		startTime  time.Time
		score      int
		timerLabel = canvas.NewText("", color.White)
		scoreLabel = canvas.NewText("Score: 0", color.White)
	)

	// Establece el estilo de texto para el cronómetro como negritas
	timerLabel.TextStyle = fyne.TextStyle{Bold: true}
	scoreLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Variable para rastrear si el juego ha sido ganado
	var hasWon bool
	var winTime time.Time
	_ = winTime

// Gorotuine que controla el juego
go func() {
		startTime = time.Now()
		for {
			time.Sleep(time.Millisecond)

			now := time.Now().UnixMilli()
			elapsed := now - game.GetThen()

			if elapsed > fpsInterval.Milliseconds() {
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

					if score == 15 && !hasWon {
						hasWon = true
						winTime = time.Now()
					}
				}
				// Dibuja el sprite de spike
				spriteDPSpike := image.Pt(spike.GetWidth()*spike.GetFrameX(), spike.GetHeight()*spike.GetFrameY())
				srSpike := image.Rectangle{spriteDPSpike, spriteDPSpike.Add(spriteSize)}

				dpSpike := image.Pt(spike.GetX(), spike.GetY())
				rSpike := image.Rectangle{dpSpike, dpSpike.Add(spriteSize)}

				draw.Draw(sprite, rSpike, spikeSprites, srSpike.Min, draw.Over)
				spikeImg = canvas.NewRasterFromImage(sprite)

				if r.Overlaps(rSpike) {
					// Si el jugador colisiona con Spike, reduce un punto del marcador
					if score > 0 {
						score--
						scoreLabel.Text = "Score: " + strconv.Itoa(score)
					}
				
					// Luego, puedes ajustar la posición de Spike como prefieras
					// Por ejemplo, puedes moverlo a una posición aleatoria
					randX := rand.Intn(400) + 100
					randY := rand.Intn(121) + 100
					spike.SetX(randX)
					spike.SetY(randY)
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

// Goroutine para mover a Jerry 
go func() {
    // Temporizador para cambiar el sprite de Jerry
    spriteTicker := time.NewTicker(200 * time.Millisecond)
    defer spriteTicker.Stop()

    for {
		select {
		case <-spriteTicker.C:
			// Cambia el sprite de Jerry aquí
			// Ajusta la dirección de animación según sea necesario
			// Actualiza la imagen de Jerry en la pantalla
	
			// Por ejemplo, puedes hacer que Jerry tenga una animación básica de caminar
			if jerry.GetXMov() < 0 {
				// Movimiento hacia la izquierda, ajusta los sprites para la izquierda
				jerry.SetFrameY(jerry.GetLeftY())
			} else if jerry.GetXMov() > 0 {
				// Movimiento hacia la derecha, ajusta los sprites para la derecha
				jerry.SetFrameY(jerry.GetRightY())
			} else if jerry.GetYMov() < 0 {
				// Movimiento hacia arriba, ajusta los sprites para arriba
				jerry.SetFrameY(jerry.GetUpY())
			} else if jerry.GetYMov() > 0 {
				// Movimiento hacia abajo, ajusta los sprites para abajo
				jerry.SetFrameY(jerry.GetDownY())
			} else {
				// Jerry no está en movimiento, establece la animación de caminar hacia abajo
				jerry.SetFrameY(jerry.GetDownY())
			}
			
			jerry.SetFrameX((jerry.GetFrameX() + 1) % jerry.GetCyclesX())
		}

        // Lógica para el movimiento automático de Jerry
        // Esto puede ser similar a la lógica que ya tienes para mover a Jerry hacia Tom

        // Obtiene la posición actual de Tom
        tomX, tomY := player.GetX(), player.GetY()

        // Calcula la diferencia en coordenadas entre Jerry y Tom
        deltaX := jerry.GetX() - tomX
        deltaY := jerry.GetY() - tomY

        // Calcula la nueva posición de Jerry en dirección opuesta a Tom, pero un poco más alejada
        // y ajusta la dirección en función de la distancia a Tom
        randX := jerry.GetX()
        randY := jerry.GetY()

        // Ajusta la dirección horizontal
        if deltaX != 0 {
            if deltaX > 0 {
                // Tom está a la izquierda de Jerry, ajusta los sprites para la izquierda
                jerry.SetFrameY(jerry.GetRightY())
            } else {
                // Tom está a la derecha de Jerry, ajusta los sprites para la derecha
                jerry.SetFrameY(jerry.GetLeftY())
            }
            randX += 2 * deltaX
        }else if deltaY != 0 {
            randY += 2 * deltaY
        }

        // Verifica si la nueva posición está dentro de los límites del juego y ajusta si es necesario
        if randX < 10 {
            randX = 10
        }
        if randX > 510-jerry.GetWidth() {
            randX = 510 - jerry.GetWidth()
        }
        if randY < 100 {
            randY = 100
        }
        if randY > 240-jerry.GetHeight() {
            randY = 240 - jerry.GetHeight()
        }

        // Calcula la cantidad de pasos necesarios para llegar a la nueva posición
        steps := 35
        stepX := float64(randX-jerry.GetX()) / float64(steps)
        stepY := float64(randY-jerry.GetY()) / float64(steps)

        for i := 0; i < steps; i++ {
            // Realiza un paso hacia la nueva posición
            jerry.SetX(jerry.GetX() + int(stepX))
            jerry.SetY(jerry.GetY() + int(stepY))
            // Verifica si la nueva posición está dentro de los límites del juego y ajusta si es necesario

            if jerry.GetX() < 10 {
                jerry.SetX(10)
            }
            if jerry.GetX() > 510-jerry.GetWidth() {
                jerry.SetX(510 - jerry.GetWidth())
            }
            if jerry.GetY() < 100 {
                jerry.SetY(100)
            }
            if jerry.GetY() > 240-jerry.GetHeight() {
                jerry.SetY(240 - jerry.GetHeight())
            }
            // Actualiza la pantalla después de cada paso
            c.Refresh()

            // Pausa para hacer que el movimiento sea visible
            time.Sleep(time.Millisecond * 10)
        }
    }
}()

//Goroutine para mover a Spike
go func() {
    // Temporizador para cambiar el sprite de Spike
    spriteTicker := time.NewTicker(200 * time.Millisecond)
    defer spriteTicker.Stop()

    for {
        select {
        case <-spriteTicker.C:
            // Obtiene la posición actual de Tom (jugador)
            tomX, tomY := player.GetX(), player.GetY()

            // Calcula la diferencia en coordenadas entre Spike y Tom
            deltaX := tomX - spike.GetX()
            deltaY := tomY - spike.GetY()

            // Calcula la distancia entre Spike y Tom
            distance := math.Sqrt(float64(deltaX*deltaX + deltaY*deltaY))

            // Define la velocidad de Spike (ajusta este valor según sea necesario)
            spikeSpeed := 2

            // Calcula la nueva posición de Spike hacia Tom
            if distance > 0 {
                // Calcula las cantidades de movimiento en las direcciones x e y
                moveX := spikeSpeed * deltaX / int(distance)
                moveY := spikeSpeed * deltaY / int(distance)

                // Actualiza la posición de Spike hacia Tom
                spike.SetX(spike.GetX() + moveX)
                spike.SetY(spike.GetY() + moveY)
            }

            // Verifica si la nueva posición de Spike colisiona con el jugador
            if spike.CollidesWith(player) {
                // Aquí puedes reducir el puntaje del jugador y realizar otras acciones
                // Por ejemplo, decrementa el puntaje y actualiza la etiqueta de puntaje
                score--
                scoreLabel.Text = "Score: " + strconv.Itoa(score)
                
                // Vuelve a posicionar a Spike en un lugar aleatorio después de la colisión
                randX := rand.Intn(400) + 100
                randY := rand.Intn(121) + 100
                spike.SetX(randX)
                spike.SetY(randY)
            }

            // Actualiza la pantalla después de cada movimiento de Spike
            c.Refresh()
        }
    }
}()

//Goroutine para contador
go func() {
		for !hasWon {
			time.Sleep(time.Millisecond)
		}

		elapsedDuration := time.Since(startTime)
		elapsedSeconds := int(elapsedDuration.Seconds())

		// Formatea la duración como "Xs"
		elapsedTimeString := strconv.Itoa(elapsedSeconds) + "s"

		wonLabel := canvas.NewText("Has ganado con un tiempo de "+elapsedTimeString, color.White)
		wonLabel.TextSize = 30
		wonLabel.TextStyle = fyne.TextStyle{Bold: true}

		// Crea un contenedor para mostrar el mensaje y el tiempo
		winContainer := container.NewVBox(wonLabel)
		winContainer.Layout = layout.NewCenterLayout()

		// Actualiza el contenido de la ventana para mostrar el mensaje y el tiempo
		w.SetContent(winContainer)
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
