package models

type Game struct {
    canvasWidth  float32
    canvasHeight float32
    fps          int
    then         int64
    margin       int
}

func NewGame(canvasWidth float32, canvasHeight float32, fps int, then int64, margin int) *Game {
    return(&Game{
        canvasWidth: canvasWidth,
        canvasHeight: canvasHeight,
        fps: fps,
        then: then,
        margin: margin,
    })
}


func (v *Game) SetWidth(canvasWidth float32){
	v.canvasWidth = canvasWidth
}

func (v *Game) GetWidth() float32 {
	return v.canvasWidth
}

func(v *Game) SetHeight(canvasHeight float32) {
	v.canvasHeight = canvasHeight
}

func(v *Game) GetHeight() float32 {
	return v.canvasHeight
}

func(v *Game) SetFps(fps int) {
	v.fps = fps
}

func(v *Game) GetFps() int {
	return v.fps
}

func(v *Game) SetThen(then int64) {
	v.then = then
}

func(v *Game) GetThen() int64 {
	return v.then
}

func(v *Game) SetMargin(margin int) {
	v.margin = margin
}

func(v *Game) GetMargin() int {
	return v.margin
}
