package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

type Semaphore struct {
	color         color.RGBA
	previousColor color.RGBA
}

func (semaphore *Semaphore) init() {
	semaphore.color = color.RGBA{R: 255}
	semaphore.previousColor = color.RGBA{R: 255, G: 255}
}

func (semaphore *Semaphore) switchLight() {
	redLightColor := color.RGBA{R: 255, A: 255}
	yellowLightColor := color.RGBA{R: 255, G: 255, A: 255}
	greenLightColor := color.RGBA{G: 255, A: 255}
	if semaphore.color == redLightColor {
		semaphore.color = yellowLightColor
	} else if semaphore.color == greenLightColor {
		semaphore.color = yellowLightColor
	} else if semaphore.previousColor == greenLightColor {
		semaphore.color = redLightColor
	} else {
		semaphore.color = greenLightColor
	}
	semaphore.previousColor = semaphore.color
}

func main() {
	const width = 800
	const height = 600

	myApp := app.New()
	myWindow := myApp.NewWindow("Simulación de Cruce con Semáforo")
	myWindow.Resize(fyne.NewSize(width, height))

	leftGrass := canvas.NewRectangle(color.RGBA{R: 124, G: 200, B: 0, A: 255})
	leftGrass.SetMinSize(fyne.NewSize(width/2-height/10, height/3))

	rightGrass := canvas.NewRectangle(color.RGBA{R: 124, G: 200, B: 0, A: 255})
	rightGrass.SetMinSize(fyne.NewSize(width/2-height/10, height/3))

	secondaryRoad := canvas.NewRectangle(color.RGBA{R: 18, G: 10, B: 5, A: 255})
	secondaryRoad.SetMinSize(fyne.NewSize(height/5, height/3))

	secondaryRoadLine := canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	secondaryRoadLine.SetMinSize(fyne.NewSize(height/100, height/3))

	mainRoad := canvas.NewRectangle(color.RGBA{R: 18, G: 10, B: 5, A: 255})
	mainRoad.SetMinSize(fyne.NewSize(width+10, height/5))

	mainRoadLine := canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	mainRoadLine.SetMinSize(fyne.NewSize(width+10, height/100))

	semaphoreBackground := canvas.NewRectangle(color.RGBA{R: 100, G: 100, B: 100, A: 255})
	semaphoreBackground.SetMinSize(fyne.NewSize(70, 20))

	redSemaphoreLight := canvas.NewRectangle(color.RGBA{R: 255, A: 255})
	redSemaphoreLight.SetMinSize(fyne.NewSize(20, 20))

	yellowSemaphoreLight := canvas.NewRectangle(color.RGBA{R: 255, G: 255, A: 60})
	yellowSemaphoreLight.SetMinSize(fyne.NewSize(20, 20))

	greenSemaphoreLight := canvas.NewRectangle(color.RGBA{G: 255, A: 60})
	greenSemaphoreLight.SetMinSize(fyne.NewSize(20, 20))

	topRightDownSemaphore := container.NewCenter(
		semaphoreBackground,
		container.NewHBox(
			redSemaphoreLight,
			yellowSemaphoreLight,
			greenSemaphoreLight,
		),
	)

	topRightLeftSemaphore := container.NewCenter(
		semaphoreBackground,
		container.NewHBox(
			redSemaphoreLight,
			yellowSemaphoreLight,
			greenSemaphoreLight,
		),
	)

	// El contenido de la aplicación
	content := container.NewCenter(
		container.NewVBox(
			container.NewHBox(
				leftGrass,
				container.NewCenter(secondaryRoad, secondaryRoadLine),
				rightGrass,
			),
			container.NewHBox(
				container.NewCenter(mainRoad, mainRoadLine),
			),
			container.NewHBox(
				leftGrass,
				container.NewCenter(secondaryRoad, secondaryRoadLine),
				rightGrass,
			),
		),
		topRightDownSemaphore,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
