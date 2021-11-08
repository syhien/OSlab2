package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
	"sync"
)

type ElelvatorStatus struct {
	lock     sync.RWMutex
	position int
	isUp     bool
}

func main() {
	elevatorStatus := ElelvatorStatus{lock: sync.RWMutex{}, position: 1, isUp: true}
	elevator := app.New()
	go floodOne(elevator, elevatorStatus)
	go floodTwo(elevator, elevatorStatus)
	go floodThree(elevator, elevatorStatus)
	elevator.Run()
}

func floodOne(elevator fyne.App, elevatorStatus ElelvatorStatus) {
	winOne := elevator.NewWindow("Flood 1")
	winOne.Resize(fyne.NewSize(200, 200))
	statusText := canvas.NewText("", color.Black)
	go func() {
		for true {
			var tmpLabel string
			elevatorStatus.lock.RLock()
			tmpLabel = strconv.Itoa(elevatorStatus.position) + " "
			if elevatorStatus.isUp {
				tmpLabel += "UP"
			} else {
				tmpLabel += "DOWN"
			}
			elevatorStatus.lock.RUnlock()
			statusText.Text = tmpLabel
		}
	}()
	content := container.NewVBox(canvas.NewText("Flood 1 Panel", color.Black), statusText)
	winOne.SetContent(content)
	winOne.Show()
}

func floodTwo(elevator fyne.App, elevatorStatus ElelvatorStatus) {
	winTwo := elevator.NewWindow("Flood 2")
	winTwo.SetContent(widget.NewLabel("2层电梯面板"))
	winTwo.Show()
}

func floodThree(elevator fyne.App, elevatorStatus ElelvatorStatus) {
	winThree := elevator.NewWindow("Flood 3 Controller")
	winThree.SetContent(widget.NewLabel("3层电梯面板"))
	winThree.Show()
}
