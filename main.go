package main

import (
    "fmt"
    "image"
    "image/draw"
    "image/png"
    "os"
    "time"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
	"game/models"
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
    playerSprites := load("./assets/tom-sprite.png")

    now := time.Now().UnixMilli()
    game := models.NewGame(
        800,
        500,
        10,
        now,
        4,
	)

    fpsInterval := int64(1000 / game.GetFps())

    player := models.NewPlayer(100, 200, 40, 72, 0, 0, 4, 3, 0, 1, 2, 9, 0, 0)

    img := canvas.NewImageFromImage(background)
    img.FillMode = canvas.ImageFillOriginal

    sprite := image.NewRGBA(background.Bounds())

    playerImg := canvas.NewRasterFromImage(sprite)
    spriteSize := image.Pt(player.GetWidth(), player.GetHeight())

    c := container.New(layout.NewMaxLayout(), img, playerImg)
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
            if player.GetX() < int(game.GetWidth())-player.GetWidth()-game.GetMargin() {
                player.SetXMov(player.GetSpeed())
            }
            player.SetFrameY(player.GetRightY())
        }
    })

    go func() {

        for {
            time.Sleep(time.Millisecond)

            now := time.Now().UnixMilli()
            elapsed := now - game.GetThen()

            if elapsed > fpsInterval {
                game.SetThen(now)

                spriteDP := image.Pt(player.GetWidth()*player.GetFrameX(), player.GetHeight()*player.GetFrameY())
                sr := image.Rectangle{spriteDP, spriteDP.Add(spriteSize)}

                dp := image.Pt(player.GetX(), player.GetY())
                r := image.Rectangle{dp, dp.Add(spriteSize)}

                draw.Draw(sprite, sprite.Bounds(), image.Transparent, image.ZP, draw.Src)
                draw.Draw(sprite, r, playerSprites, sr.Min, draw.Src)
                playerImg = canvas.NewRasterFromImage(sprite)

                if player.GetXMov() != 0 || player.GetYMov() != 0 {
                    player.SetX(player.GetX() + player.GetXMov())
					player.SetY(player.GetY() + player.GetYMov())
                    player.SetFrameX((player.GetFrameX() + 1) % player.GetCyclesX())
                    player.SetXMov(0)
                    player.SetYMov(0)
                } else {
                    player.SetFrameX(0)
                }

                c.Refresh()

            }
        }

    }()

    w.CenterOnScreen()
    w.ShowAndRun()
}