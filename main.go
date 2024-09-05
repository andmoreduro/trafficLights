package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

const windowWidth = 800
const windowHeight = 600

//type Grass struct{}

func createGrassObject() fyne.CanvasObject {
	background := canvas.NewRectangle(color.RGBA{R: 124, G: 200, B: 0, A: 255})
	background.SetMinSize(fyne.NewSize(windowWidth/2-windowHeight/10, windowHeight/3))
	return background
}

//type Road struct{}

func createRoadObject(main bool) fyne.CanvasObject {
	background := canvas.NewRectangle(color.RGBA{R: 18, G: 10, B: 5, A: 255})
	line := canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	if main {
		background.SetMinSize(fyne.NewSize(windowWidth+10, windowHeight/5))
		line.SetMinSize(fyne.NewSize(windowWidth+10, windowHeight/100))
	} else {
		background.SetMinSize(fyne.NewSize(windowHeight/5, windowHeight/3))
		line.SetMinSize(fyne.NewSize(windowHeight/100, windowHeight/3))
	}
	return container.NewCenter(background, line)
}

type Semaphore struct {
	redLight      *canvas.Rectangle
	yellowLight   *canvas.Rectangle
	greenLight    *canvas.Rectangle
	activeLight   string
	previousLight string
}

func (semaphore *Semaphore) init() {
	semaphore.redLight = canvas.NewRectangle(color.RGBA{R: 255, A: 255})
	semaphore.redLight.SetMinSize(fyne.NewSize(20, 20))
	semaphore.yellowLight = canvas.NewRectangle(color.RGBA{R: 255, G: 255, A: 60})
	semaphore.yellowLight.SetMinSize(fyne.NewSize(20, 20))
	semaphore.greenLight = canvas.NewRectangle(color.RGBA{G: 255, A: 60})
	semaphore.greenLight.SetMinSize(fyne.NewSize(20, 20))
	semaphore.activeLight = "red"
	semaphore.previousLight = "yellow"
}

// Inicializa el semáforo y retorna el objeto que lo representa gráficamente
func createSemaphoreObject(semaphore *Semaphore, vertical bool, inverted bool) fyne.CanvasObject {
	background := canvas.NewRectangle(color.RGBA{R: 100, G: 100, B: 100, A: 255})
	if vertical {
		background.SetMinSize(fyne.NewSize(20, 70))
	} else {
		background.SetMinSize(fyne.NewSize(70, 20))
	}
	var lightContainer *fyne.Container
	if vertical {
		lightContainer = container.NewVBox()
	} else {
		lightContainer = container.NewHBox()
	}
	if inverted {
		lightContainer.Add(semaphore.greenLight)
		lightContainer.Add(semaphore.yellowLight)
		lightContainer.Add(semaphore.redLight)
	} else {
		lightContainer.Add(semaphore.redLight)
		lightContainer.Add(semaphore.yellowLight)
		lightContainer.Add(semaphore.greenLight)
	}
	return container.NewCenter(
		background,
		lightContainer,
	)
}

// Cambia la luz del semáforo
func (semaphore *Semaphore) switchLight() {
	if semaphore.activeLight == "red" {
		semaphore.redLight = canvas.NewRectangle(color.RGBA{R: 255, A: 60})
		semaphore.yellowLight = canvas.NewRectangle(color.RGBA{R: 255, G: 255, A: 255})
		semaphore.activeLight = "yellow"
		semaphore.previousLight = "red"
	} else if semaphore.activeLight == "green" {
		semaphore.greenLight = canvas.NewRectangle(color.RGBA{G: 255, A: 60})
		semaphore.yellowLight = canvas.NewRectangle(color.RGBA{R: 255, G: 255, A: 255})
		semaphore.activeLight = "yellow"
		semaphore.previousLight = "green"
	} else if semaphore.previousLight == "red" {
		semaphore.yellowLight = canvas.NewRectangle(color.RGBA{R: 255, G: 255, A: 60})
		semaphore.greenLight = canvas.NewRectangle(color.RGBA{G: 255, A: 255})
		semaphore.activeLight = "green"
		semaphore.previousLight = "yellow"
	} else {
		semaphore.yellowLight = canvas.NewRectangle(color.RGBA{R: 255, G: 255, A: 60})
		semaphore.redLight = canvas.NewRectangle(color.RGBA{R: 255, A: 255})
		semaphore.activeLight = "red"
		semaphore.previousLight = "yellow"
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Simulación de Cruce con Semáforo")
	myWindow.Resize(fyne.NewSize(windowWidth, windowHeight))

	topRightDownSemaphore := Semaphore{}
	topRightDownSemaphore.init()
	topLeftRightSemaphore := Semaphore{}
	topLeftRightSemaphore.init()
	bottomRightLeftSemaphore := Semaphore{}
	bottomRightLeftSemaphore.init()
	bottomLeftTopSemaphore := Semaphore{}
	bottomLeftTopSemaphore.init()

	// El contenido de la aplicación
	content := container.NewCenter(
		container.NewVBox(
			container.NewHBox(
				createGrassObject(),
				createRoadObject(false),
				createGrassObject(),
			),
			container.NewHBox(
				createRoadObject(true),
			),
			container.NewHBox(
				createGrassObject(),
				createRoadObject(false),
				createGrassObject(),
			),
		),
		container.NewWithoutLayout(
			createSemaphoreObject(&topRightDownSemaphore, false, false),
			createSemaphoreObject(&topLeftRightSemaphore, true, true),
			createSemaphoreObject(&bottomRightLeftSemaphore, true, false),
			createSemaphoreObject(&bottomLeftTopSemaphore, false, true),
		),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
