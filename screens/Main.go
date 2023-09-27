package screens

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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
	backgroundImage := canvas.NewImageFromFile("assets/screen.png")
	backgroundImage.Resize(fyne.NewSize(1280, 720))
	backgroundImage.Move(fyne.NewPos(0, 0))


	start := widget.NewButton("Start", m.StartGame)

	start.Resize(fyne.NewSize(160, 40));
	start.Move(fyne.NewPos(180, 550));

	exit := widget.NewButton("Exit", m.ExitGame)
	exit.Resize(fyne.NewSize(160, 40));
	exit.Move(fyne.NewPos(180, 600));

	m.window.SetContent(container.NewWithoutLayout(backgroundImage, start, exit))
	m.window.Resize(fyne.NewSize(1280, 720))
	m.window.SetFixedSize(true)
}

func (m *MainScene) ExitGame() {
	m.window.Close()
}

func (m *MainScene) StartGame() {
	NewGame(m.window)
}
