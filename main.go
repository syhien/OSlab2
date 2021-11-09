package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
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
	elevator.Run()
	for true {
		select {
		case request <- requestOne:
			for true {
				elevatorStatus.lock.Lock()
				if elevatorStatus.position == 1 {
					elevatorStatus.lock.Unlock()
					break
				}
				elevatorStatus.position--
				elevatorStatus.lock.Unlock()
				time.Sleep(5 * time.Second)
			}
		case request <- requestTwo:
			for true {
				elevatorStatus.lock.Lock()
				if elevatorStatus.position == 2 {
					elevatorStatus.lock.Unlock()
					break
				}
				elevatorStatus.position = 2
				elevatorStatus.lock.Unlock()
				time.Sleep(5 * time.Second)
			}
		case request <- requestThree:
			for true {
				elevatorStatus.lock.Lock()
				if elevatorStatus.position == 3 {
					elevatorStatus.lock.Unlock()
					break
				}
				elevatorStatus.position++
				elevatorStatus.lock.Unlock()
				time.Sleep(5 * time.Second)
			}
		}
	}
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

	content := container.NewVBox(titleLabel, statusText)
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
	content := container.NewVBox(titleLabel, statusText)
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
	content := container.NewVBox(titleLabel, statusLabel)
	winThree.SetContent(content)
	winThree.Show()
}
