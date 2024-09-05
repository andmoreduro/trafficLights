package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"runtime"
	"time"
)

const windowWidth = 800
const windowHeight = 600
const semaphoreLength = 70
const semaphoreGirth = 20

type Semaphore struct {
	redLight      canvas.Rectangle
	yellowLight   canvas.Rectangle
	greenLight    canvas.Rectangle
	vertical      bool
	inverted      bool
	activeLight   string
	previousLight string
}

// Inicializa el semáforo y guarda el objeto que lo representa gráficamente
func (semaphore *Semaphore) init(vertical bool, inverted bool) {
	semaphore.redLight = *canvas.NewRectangle(color.RGBA{R: 255, A: 255})
	semaphore.redLight.SetMinSize(fyne.NewSize(semaphoreGirth, semaphoreGirth))
	semaphore.yellowLight = *canvas.NewRectangle(color.RGBA{R: 255, G: 255, A: 60})
	semaphore.yellowLight.SetMinSize(fyne.NewSize(semaphoreGirth, semaphoreGirth))
	semaphore.greenLight = *canvas.NewRectangle(color.RGBA{G: 255, A: 60})
	semaphore.greenLight.SetMinSize(fyne.NewSize(semaphoreGirth, semaphoreGirth))
	semaphore.vertical = vertical
	semaphore.inverted = inverted
	semaphore.activeLight = "red"
	semaphore.previousLight = "yellow"
}

func (semaphore *Semaphore) getObject() fyne.CanvasObject {
	background := canvas.NewRectangle(color.RGBA{R: 100, G: 100, B: 100, A: 255})
	if semaphore.vertical {
		background.SetMinSize(fyne.NewSize(semaphoreGirth, semaphoreLength))
	} else {
		background.SetMinSize(fyne.NewSize(semaphoreLength, semaphoreGirth))
	}
	var lightContainer *fyne.Container
	if semaphore.vertical {
		lightContainer = container.NewVBox()
	} else {
		lightContainer = container.NewHBox()
	}
	if semaphore.inverted {
		lightContainer.Add(&semaphore.greenLight)
		lightContainer.Add(&semaphore.yellowLight)
		lightContainer.Add(&semaphore.redLight)
	} else {
		lightContainer.Add(&semaphore.redLight)
		lightContainer.Add(&semaphore.yellowLight)
		lightContainer.Add(&semaphore.greenLight)
	}
	result := container.NewCenter(
		background,
		lightContainer,
	)
	return result
}

// Cambia la luz del semáforo
func (semaphore *Semaphore) switchLight() {
	if semaphore.activeLight == "red" {
		semaphore.redLight.FillColor = color.RGBA{R: 255, A: 60}
		semaphore.yellowLight.FillColor = color.RGBA{R: 255, G: 255, A: 255}
		semaphore.activeLight = "yellow"
		semaphore.previousLight = "red"
	} else if semaphore.activeLight == "green" {
		semaphore.greenLight.FillColor = color.RGBA{G: 255, A: 60}
		semaphore.yellowLight.FillColor = color.RGBA{R: 255, G: 255, A: 255}
		semaphore.activeLight = "yellow"
		semaphore.previousLight = "green"
	} else if semaphore.previousLight == "red" {
		semaphore.yellowLight.FillColor = color.RGBA{R: 255, G: 255, A: 60}
		semaphore.greenLight.FillColor = color.RGBA{G: 255, A: 255}
		semaphore.activeLight = "green"
		semaphore.previousLight = "yellow"
	} else {
		semaphore.yellowLight.FillColor = color.RGBA{R: 255, G: 255, A: 60}
		semaphore.redLight.FillColor = color.RGBA{R: 255, A: 255}
		semaphore.activeLight = "red"
		semaphore.previousLight = "yellow"
	}
}

func createGrassObject(semaphoreObject fyne.CanvasObject, top bool, right bool) fyne.CanvasObject {
	background := canvas.NewRectangle(color.RGBA{R: 124, G: 200, B: 0, A: 255})
	const grassWidth = windowWidth/2 - windowHeight/10
	const grassHeight = windowHeight / 3
	background.SetMinSize(fyne.NewSize(grassWidth, grassHeight))
	if semaphoreObject == nil {
		return background
	}
	var semaphoreObjectX float32
	var semaphoreObjectY float32
	if top {
		if right {
			semaphoreObjectX = semaphoreLength/2 - grassWidth/3
			semaphoreObjectY = semaphoreGirth/2 + 3*grassHeight/8
		} else {
			semaphoreObjectX = semaphoreGirth/2 + 109*grassWidth/256
			semaphoreObjectY = semaphoreLength/2 + 15*grassHeight/64
		}
	} else {
		if right {
			semaphoreObjectX = semaphoreGirth/2 - 109*grassWidth/256
			semaphoreObjectY = semaphoreLength/2 - 15*grassHeight/64
		} else {
			semaphoreObjectX = semaphoreLength/2 + grassWidth/3
			semaphoreObjectY = semaphoreGirth/2 - 3*grassHeight/8
		}
	}
	semaphoreObject.Move(fyne.NewPos(semaphoreObjectX, semaphoreObjectY))
	return container.NewCenter(
		background,
		container.NewWithoutLayout(semaphoreObject),
	)
}

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

func main() {
	runtime.GOMAXPROCS(4)

	myApp := app.New()
	myWindow := myApp.NewWindow("Simulación de Cruce con Semáforo")
	myWindow.Resize(fyne.NewSize(windowWidth, windowHeight))

	// Arriba a la izquierda
	tlSemaphore := Semaphore{}
	tlSemaphore.init(true, true)
	tlSemaphoreObject := tlSemaphore.getObject()
	// Arriba a la derecha
	trSemaphore := Semaphore{}
	trSemaphore.init(false, false)
	trSemaphoreObject := trSemaphore.getObject()
	// Abajo a la izquierda
	blSemaphore := Semaphore{}
	blSemaphore.init(false, true)
	blSemaphoreObject := blSemaphore.getObject()
	// Abajo a la derecha
	brSemaphore := Semaphore{}
	brSemaphore.init(true, false)
	brSemaphoreObject := brSemaphore.getObject()

	// El contenido de la aplicación
	content := container.NewCenter(
		container.NewVBox(
			container.NewHBox(
				createGrassObject(tlSemaphoreObject, true, false),
				createRoadObject(false),
				createGrassObject(trSemaphoreObject, true, true),
			),
			container.NewHBox(
				createRoadObject(true),
			),
			container.NewHBox(
				createGrassObject(blSemaphoreObject, false, false),
				createRoadObject(false),
				createGrassObject(brSemaphoreObject, false, true),
			),
		),
	)

	// Todas las cosas que funcionaran el paralelo
	// Semáforo arriba a la izquierda
	go func() {
		time.Sleep(10 * time.Second)
		for {
			tlSemaphore.switchLight()
			tlSemaphoreObject.Refresh()
			time.Sleep(1 * time.Second)
			tlSemaphore.switchLight()
			tlSemaphoreObject.Refresh()
			time.Sleep(9 * time.Second)
		}
	}()

	// Semáforo arriba a la derecha
	go func() {
		for {
			trSemaphore.switchLight()
			trSemaphoreObject.Refresh()
			time.Sleep(1 * time.Second)
			trSemaphore.switchLight()
			trSemaphoreObject.Refresh()
			time.Sleep(9 * time.Second)
		}
	}()

	// Semáforo abajo a la izquierda
	go func() {
		for {
			blSemaphore.switchLight()
			blSemaphoreObject.Refresh()
			time.Sleep(1 * time.Second)
			blSemaphore.switchLight()
			blSemaphoreObject.Refresh()
			time.Sleep(9 * time.Second)
		}
	}()

	// Semáforo abajo a la derecha
	go func() {
		time.Sleep(10 * time.Second)
		for {
			brSemaphore.switchLight()
			brSemaphoreObject.Refresh()
			time.Sleep(1 * time.Second)
			brSemaphore.switchLight()
			brSemaphoreObject.Refresh()
			time.Sleep(9 * time.Second)
		}
	}()

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
