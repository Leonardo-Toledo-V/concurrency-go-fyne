package models

type Tom struct {
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


func NewTom(x int, y int, width int, height int, frameX int, frameY int, cyclesX int, upY int, downY int, leftY int, rightY int, speed int, xMov int, yMov int) *Tom{
	return (&Tom{
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

func (v *Tom) SetX(x int){
    v.x = x
}

func(v *Tom) GetX() int {
    return v.x
}


func (v *Tom) SetY(y int){
    v.y = y
}

func(v *Tom) GetY() int {
    return v.y
}

func (v *Tom) SetWidth(width int){
    v.width = width
}

func(v *Tom) GetWidth() int {
    return v.width
}

func (v *Tom) SetHeight(height int){
    v.height = height
}

func(v *Tom) GetHeight() int {
    return v.height
}

func(v *Tom) SetFrameX(frameX int){
    v.frameX = frameX
}

func(v *Tom) GetFrameX() int {
    return v.frameX
}

func (v *Tom) SetFrameY(frameY int){
    v.frameY = frameY
}

func(v *Tom) GetFrameY() int {
    return v.frameY
}

func (v *Tom) SetCyclesX(cyclesX int){
    v.cyclesX = cyclesX
}

func(v *Tom) GetCyclesX() int {
    return v.cyclesX
}

func (v *Tom) SetUpY(upY int){
    v.upY = upY
}

func(v *Tom) GetUpY() int {
    return v.upY
}

func(v *Tom) SetDownY(downY int){
    v.downY = downY
}

func(v *Tom) GetDownY() int {
    return v.downY
}

func(v *Tom) SetLeftY(leftY int){
    v.leftY = leftY
}

func(v *Tom) GetLeftY() int {
    return v.leftY
}

func(v *Tom) SetRightY(rightY int){
    v.rightY = rightY
}

func(v *Tom) GetRightY() int {
    return v.rightY
}

func (v *Tom) SetSpeed(speed int){
    v.speed = speed
}

func(v *Tom) GetSpeed() int {
    return v.speed
}

func(v *Tom) SetXMov(xMov int){
    v.xMov = xMov
}

func(v *Tom) GetXMov() int {
    return v.xMov
}

func(v *Tom) SetYMov(yMov int){
    v.yMov = yMov
}

func(v *Tom) GetYMov() int {
    return v.yMov
}


