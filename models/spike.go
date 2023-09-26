package models

type Spike struct {
    x       int
    y       int
    width   int
    height  int
    frameX  int
    frameY  int
    cyclesX int
    upY     int
    downY   int
    leftY   int
    rightY  int
    speed   int
    xMov    int
    yMov    int
}


func NewSpike(x int, y int, width int, height int, frameX int, frameY int, cyclesX int, upY int, downY int, leftY int, rightY int, speed int, xMov int, yMov int) *Spike{
	return (&Spike{
		x: x,
		y: y,
        width: width,
        height: height,
        frameX: frameX,
        frameY: frameY,
        cyclesX: cyclesX,
        upY: upY,
        downY: downY,
        leftY: leftY,
        rightY: rightY,
        speed: speed,
        xMov: xMov,
        yMov: yMov,
	})
}

func (v *Spike) SetX(x int){
    v.x = x
}

func(v *Spike) GetX() int {
    return v.x
}


func (v *Spike) SetY(y int){
    v.y = y
}

func(v *Spike) GetY() int {
    return v.y
}

func (v *Spike) SetWidth(width int){
    v.width = width
}

func(v *Spike) GetWidth() int {
    return v.width
}

func (v *Spike) SetHeight(height int){
    v.height = height
}

func(v *Spike) GetHeight() int {
    return v.height
}

func(v *Spike) SetFrameX(frameX int){
    v.frameX = frameX
}

func(v *Spike) GetFrameX() int {
    return v.frameX
}

func (v *Spike) SetFrameY(frameY int){
    v.frameY = frameY
}

func(v *Spike) GetFrameY() int {
    return v.frameY
}

func (v *Spike) SetCyclesX(cyclesX int){
    v.cyclesX = cyclesX
}

func(v *Spike) GetCyclesX() int {
    return v.cyclesX
}

func (v *Spike) SetUpY(upY int){
    v.upY = upY
}

func(v *Spike) GetUpY() int {
    return v.upY
}

func(v *Spike) SetDownY(downY int){
    v.downY = downY
}

func(v *Spike) GetDownY() int {
    return v.downY
}

func(v *Spike) SetLeftY(leftY int){
    v.leftY = leftY
}

func(v *Spike) GetLeftY() int {
    return v.leftY
}

func(v *Spike) SetRightY(rightY int){
    v.rightY = rightY
}

func(v *Spike) GetRightY() int {
    return v.rightY
}

func (v *Spike) SetSpeed(speed int){
    v.speed = speed
}

func(v *Spike) GetSpeed() int {
    return v.speed
}

func(v *Spike) SetXMov(xMov int){
    v.xMov = xMov
}

func(v *Spike) GetXMov() int {
    return v.xMov
}

func(v *Spike) SetYMov(yMov int){
    v.yMov = yMov
}

func(v *Spike) GetYMov() int {
    return v.yMov
}


func (s *Spike) CollidesWith(player *Tom) bool {
    spikeLeft := s.x
    spikeRight := s.x + s.width
    spikeTop := s.y
    spikeBottom := s.y + s.height

    playerLeft := player.x
    playerRight := player.x + player.width
    playerTop := player.y
    playerBottom := player.y + player.height

    if spikeLeft < playerRight && spikeRight > playerLeft && spikeTop < playerBottom && spikeBottom > playerTop {
        // Hay superposición, lo que significa que colisionaron
        return true
    }

    // No hay superposición, no colisionaron
    return false
}