package models

type Player struct {
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


func NewPlayer(x int, y int, width int, height int, frameX int, frameY int, cyclesX int, upY int, downY int, leftY int, rightY int, speed int, xMov int, yMov int) *Player{
	return (&Player{
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

func (v *Player) SetX(x int){
    v.x = x
}

func(v *Player) GetX() int {
    return v.x
}


func (v *Player) SetY(y int){
    v.y = y
}

func(v *Player) GetY() int {
    return v.y
}

func (v *Player) SetWidth(width int){
    v.width = width
}

func(v *Player) GetWidth() int {
    return v.width
}

func (v *Player) SetHeight(height int){
    v.height = height
}

func(v *Player) GetHeight() int {
    return v.height
}

func(v *Player) SetFrameX(frameX int){
    v.frameX = frameX
}

func(v *Player) GetFrameX() int {
    return v.frameX
}

func (v *Player) SetFrameY(frameY int){
    v.frameY = frameY
}

func(v *Player) GetFrameY() int {
    return v.frameY
}

func (v *Player) SetCyclesX(cyclesX int){
    v.cyclesX = cyclesX
}

func(v *Player) GetCyclesX() int {
    return v.cyclesX
}

func (v *Player) SetUpY(upY int){
    v.upY = upY
}

func(v *Player) GetUpY() int {
    return v.upY
}

func(v *Player) SetDownY(downY int){
    v.downY = downY
}

func(v *Player) GetDownY() int {
    return v.downY
}

func(v *Player) SetLeftY(leftY int){
    v.leftY = leftY
}

func(v *Player) GetLeftY() int {
    return v.leftY
}

func(v *Player) SetRightY(rightY int){
    v.rightY = rightY
}

func(v *Player) GetRightY() int {
    return v.rightY
}

func (v *Player) SetSpeed(speed int){
    v.speed = speed
}

func(v *Player) GetSpeed() int {
    return v.speed
}

func(v *Player) SetXMov(xMov int){
    v.xMov = xMov
}

func(v *Player) GetXMov() int {
    return v.xMov
}

func(v *Player) SetYMov(yMov int){
    v.yMov = yMov
}

func(v *Player) GetYMov() int {
    return v.yMov
}


