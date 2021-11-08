package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"strconv"
	"sync"
)

type ElelvatorStatus struct {
	lock     sync.RWMutex
	position int
	isUp     bool
	isMoving bool
}

func main() {
	elevatorStatus := ElelvatorStatus{lock: sync.RWMutex{}, position: 1, isUp: true, isMoving: true}
	elevator := app.New()
	go floodOne(elevator, elevatorStatus)
	go floodTwo(elevator, elevatorStatus)
	go floodThree(elevator, elevatorStatus)
	elevator.Run()
}

func floodOne(elevator fyne.App, elevatorStatus ElelvatorStatus) {
	winOne := elevator.NewWindow("1")
	winOne.Resize(fyne.NewSize(200, 200))
	titleText := canvas.NewText("Flood 1 Panel", color.Black)
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	statusText := canvas.NewText("", color.Black)
	go func() {
		for true {
			tmpLabel := "> "
			elevatorStatus.lock.RLock()
			tmpLabel += strconv.Itoa(elevatorStatus.position) + " "
			if elevatorStatus.isMoving {
				if elevatorStatus.isUp {
					tmpLabel += "UP"
				} else {
					tmpLabel += "DOWN"
				}
			} else {
				tmpLabel += ""
			}
			elevatorStatus.lock.RUnlock()
			statusText.Text = tmpLabel
		}
	}()
	content := container.NewVBox(titleText, statusText)
	winOne.SetContent(content)
	winOne.Show()
}

func floodTwo(elevator fyne.App, elevatorStatus ElelvatorStatus) {
	winTwo := elevator.NewWindow("2")
	winTwo.Resize(fyne.NewSize(200, 200))
	titleText := canvas.NewText("Flood 2 Panel", color.Black)
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	statusText := canvas.NewText("", color.Black)
	go func() {
		for true {
			tmpLabel := "> "
			elevatorStatus.lock.RLock()
			tmpLabel += strconv.Itoa(elevatorStatus.position) + " "
			if elevatorStatus.isMoving {
				if elevatorStatus.isUp {
					tmpLabel += "UP"
				} else {
					tmpLabel += "DOWN"
				}
			} else {
				tmpLabel += ""
			}
			elevatorStatus.lock.RUnlock()
			statusText.Text = tmpLabel
		}
	}()
	content := container.NewVBox(titleText, statusText)
	winTwo.SetContent(content)
	winTwo.Show()
}

func floodThree(elevator fyne.App, elevatorStatus ElelvatorStatus) {
	winThree := elevator.NewWindow("3")
	winThree.Resize(fyne.NewSize(200, 200))
	titleText := canvas.NewText("Flood 3 Panel", color.Black)
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	statusText := canvas.NewText("", color.Black)
	go func() {
		for true {
			tmpLabel := "> "
			elevatorStatus.lock.RLock()
			tmpLabel += strconv.Itoa(elevatorStatus.position) + " "
			if elevatorStatus.isMoving {
				if elevatorStatus.isUp {
					tmpLabel += "UP"
				} else {
					tmpLabel += "DOWN"
				}
			} else {
				tmpLabel += ""
			}
			elevatorStatus.lock.RUnlock()
			statusText.Text = tmpLabel
		}
	}()
	content := container.NewVBox(titleText, statusText)
	winThree.SetContent(content)
	winThree.Show()
}
