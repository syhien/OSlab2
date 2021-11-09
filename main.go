package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type ElelvatorStatus struct {
	lock     sync.RWMutex
	position int
	isUp     bool
	isMoving bool
}

func main() {
	rand.Seed(time.Now().Unix())
	elevatorStatus := ElelvatorStatus{lock: sync.RWMutex{}, position: rand.Intn(3) + 1, isUp: true, isMoving: rand.Float64() > 0.5}
	elevator := app.New()
	requestOne := make(chan int)
	defer close(requestOne)
	requestTwo := make(chan int)
	defer close(requestTwo)
	requestThree := make(chan int)
	defer close(requestThree)
	go floodOne(elevator, elevatorStatus, requestOne)
	go floodTwo(elevator, elevatorStatus, requestTwo)
	go floodThree(elevator, elevatorStatus, requestThree)
	go func() {
		for true {
			select {
			case request := <-requestOne:
				fmt.Println(request)
			case request := <-requestTwo:
				fmt.Println(request)
			case request := <-requestThree:
				fmt.Println(request)
			}
		}
	}()
	elevator.Run()
}

func floodOne(elevator fyne.App, elevatorStatus ElelvatorStatus, requestChan chan int) {
	winOne := elevator.NewWindow("1")
	winOne.Resize(fyne.NewSize(200, 200))
	titleLabel := widget.NewLabel("Flood 1 Panel")
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}
	statusText := widget.NewLabel("")
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
	goto2Button := widget.NewButton("2", func() {
		dialog.NewInformation("Flood 1", "2 pressed", winOne).Show()
		requestChan <- 2
	})
	goto3Button := widget.NewButton("3", func() {
		dialog.NewInformation("Flood 1", "3 pressed", winOne).Show()
		requestChan <- 3
	})
	content := container.NewVBox(titleLabel, statusText, goto2Button, goto3Button)
	winOne.SetContent(content)
	winOne.Show()
}

func floodTwo(elevator fyne.App, elevatorStatus ElelvatorStatus, requestChan chan int) {
	winTwo := elevator.NewWindow("2")
	winTwo.Resize(fyne.NewSize(200, 200))
	titleLabel := widget.NewLabel("Flood 2 Panel")
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}
	statusText := widget.NewLabel("")
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
	goto1Button := widget.NewButton("1", func() {
		dialog.NewInformation("Flood 2", "1 pressed", winTwo).Show()
		requestChan <- 1
	})
	goto3Button := widget.NewButton("3", func() {
		dialog.NewInformation("Flood 2", "3 pressed", winTwo).Show()
		requestChan <- 3
	})
	content := container.NewVBox(titleLabel, statusText, goto1Button, goto3Button)
	winTwo.SetContent(content)
	winTwo.Show()
}

func floodThree(elevator fyne.App, elevatorStatus ElelvatorStatus, requestChan chan int) {
	winThree := elevator.NewWindow("3")
	winThree.Resize(fyne.NewSize(200, 200))
	titleLabel := widget.NewLabel("Flood 3 Panel")
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}
	statusLabel := widget.NewLabel("")
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
			statusLabel.Text = tmpLabel
		}
	}()
	goto1Button := widget.NewButton("1", func() {
		dialog.NewInformation("Flood 3", "1 pressed", winThree).Show()
		requestChan <- 1
	})
	goto2Button := widget.NewButton("2", func() {
		dialog.NewInformation("Flood 3", "2 pressed", winThree).Show()
		requestChan <- 2
	})
	content := container.NewVBox(titleLabel, statusLabel, goto1Button, goto2Button)
	winThree.SetContent(content)
	winThree.Show()
}
