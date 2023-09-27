package screens;

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MainScene struct {
	window fyne.Window
}

func NewMainScene(window fyne.Window) *MainScene {
	MainScene := &MainScene{
		window: window,
	}
	MainScene.InitApp()
	return MainScene
}

func (m *MainScene) InitApp() {
	m.DrawSceneMenu()
}

func (m *MainScene) DrawSceneMenu() {
	titleImg := canvas.NewImageFromFile("assets/logo.png")
	titleImg.Resize(fyne.NewSize(300, 270))
	titleImg.Move(fyne.NewPos(50, 10))
	titleContainer := container.NewWithoutLayout(titleImg)

	start := widget.NewButton("Start", m.StartGame)

	exit := widget.NewButton("Exit", m.ExitGame)

	container_center := container.NewVBox(
		titleContainer,
		layout.NewSpacer(),
		start,
		exit,
	)

	m.window.SetContent(container_center)
	m.window.Resize(fyne.NewSize(400, 500))
	m.window.SetFixedSize(true)
}

func (m *MainScene) ExitGame() {
	m.window.Close()
}


func (m *MainScene) StartGame() {
	NewGame(m.window)
}