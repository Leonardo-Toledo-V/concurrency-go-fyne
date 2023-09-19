package models

type Jerry struct {
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

func NewJerry(x int, y int, width int, height int, frameX int, frameY int, cyclesX int, upY int, downY int, leftY int, rightY int, speed int, xMov int, yMov int) *Jerry{
	return (&Jerry{
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


func (v *Jerry) SetX(x int){
    v.x = x
}

func(v *Jerry) GetX() int {
    return v.x
}


func (v *Jerry) SetY(y int){
    v.y = y
}

func(v *Jerry) GetY() int {
    return v.y
}

func (v *Jerry) SetWidth(width int){
    v.width = width
}

func(v *Jerry) GetWidth() int {
    return v.width
}

func (v *Jerry) SetHeight(height int){
    v.height = height
}

func(v *Jerry) GetHeight() int {
    return v.height
}

func(v *Jerry) SetFrameX(frameX int){
    v.frameX = frameX
}

func(v *Jerry) GetFrameX() int {
    return v.frameX
}

func (v *Jerry) SetFrameY(frameY int){
    v.frameY = frameY
}

func(v *Jerry) GetFrameY() int {
    return v.frameY
}

func (v *Jerry) SetCyclesX(cyclesX int){
    v.cyclesX = cyclesX
}

func(v *Jerry) GetCyclesX() int {
    return v.cyclesX
}

func (v *Jerry) SetUpY(upY int){
    v.upY = upY
}

func(v *Jerry) GetUpY() int {
    return v.upY
}

func(v *Jerry) SetDownY(downY int){
    v.downY = downY
}

func(v *Jerry) GetDownY() int {
    return v.downY
}

func(v *Jerry) SetLeftY(leftY int){
    v.leftY = leftY
}

func(v *Jerry) GetLeftY() int {
    return v.leftY
}

func(v *Jerry) SetRightY(rightY int){
    v.rightY = rightY
}

func(v *Jerry) GetRightY() int {
    return v.rightY
}

func (v *Jerry) SetSpeed(speed int){
    v.speed = speed
}

func(v *Jerry) GetSpeed() int {
    return v.speed
}

func(v *Jerry) SetXMov(xMov int){
    v.xMov = xMov
}

func(v *Jerry) GetXMov() int {
    return v.xMov
}

func(v *Jerry) SetYMov(yMov int){
    v.yMov = yMov
}

func(v *Jerry) GetYMov() int {
    return v.yMov
}


