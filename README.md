# 电梯模拟程序

## 简述

运用多线程，模拟一个三层楼房中的电梯运行

## 运行

### 从code运行

```shell
$ go get fyne.io/fyne/v2
$ go run main.go
```

### 二进制文件运行

#### Windows & Mac OS

在[Github Release](https://github.com/syhien/OSlab2/releases/latest)中下载对应二进制文件运行

*程序在开发过程中仓库为Private状态，本report提交后将转为Public*

#### Linux各发行版

```shell
$ go get fyne.io/fyne/v2
$ go get fyne.io/fyne/v2/cmd/fyne
$ fyne package -os linux
$ ./OSlab2
```

## 开发环境

- Go 1.17.3
- Fyne 2.1.1 - Go module
- Fyne Cross 1.1.3 - Go module

## 实现细节

### 选择Go和Fyne的原因

因为这次实验不限制编程的语言，所以我选择了我一直有在了解但缺少实践经验的Go。

Go基于C诞生，继承了C的高性能，并且在语言的层面上支持协程（多线程），有优异的并发性能，在配置很一般的服务器上也能轻松掌控百万级别的并发数。我个人认为Go的协程是（我所接触过的语言中）使用多线程最方便的语言，使用`go`关键字如`go func_name()`即启动了一个协程。线程之间可通过共享变量、信道`chan`来通信。

因为这次实验不限制GUI CLI，所以我选择使用Fyne这个简单易用的第三方Go GUI库实现界面。

考虑到需要同时展示多个视角，并且还需要进行一些交互，我觉得GUI显然更有优势一些，没做太多考虑就放弃了CLI

### 电梯状态

使用一个包含读写锁的结构`ElevatorStatus`来共享电梯状态：

```go
type ElelvatorStatus struct {
	lock     sync.RWMutex
	position int
	isUp     bool
	isMoving bool
}
```

`sync.RWMutex`是Go的读写锁，和一般的锁差别在于：读与读之间不互斥，意味着可以有复数个线程同时读取。当然，读与写仍然是互斥的。

下面这段在视图中更新电梯状态的代码是一个使用锁的例子：

```go
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
```

### 主线程

程序的入口是`func main()`：

```go
func main() {
	rand.Seed(time.Now().Unix())
	elevatorStatus := ElelvatorStatus{lock: sync.RWMutex{}, position: rand.Intn(3) + 1, isUp: false, isMoving: false} // 随机初始化电梯停在某一层
	elevator := app.New()
	requestOne := make(chan int)
	defer close(requestOne) // 用defer关键字保证退出函数前一定关闭信道
	requestTwo := make(chan int)
	defer close(requestTwo)
	requestThree := make(chan int)
	defer close(requestThree)
	go floodOne(elevator, &elevatorStatus, requestOne)
	go floodTwo(elevator, &elevatorStatus, requestTwo)
	go floodThree(elevator, &elevatorStatus, requestThree)
	go elevatorController(&elevatorStatus, requestOne, requestTwo, requestThree)
	elevator.Run()
}
```

`main()`中初始化了一个`elevatorStatus`作为一台全新的电梯。

接着，开启三个`int`类型的信道`chan int`用于各楼层的电梯面板与电梯中控交互。

然后，`go`三条协程用于控制三个楼层的电梯面板，在`floodXXX`内各自新建窗口、接受按键请求、更新电梯状态等；`go`一个`elevatorController`作为电梯中控

最后运行GUI类`elevator`

### 电梯中控

电梯中控要做的事情无非就是：

1. 等待某个楼层发出使用电梯的请求
2. 将电梯移动到sender所在的楼层
3. 将电梯移动到指定的楼层
4. 返回1

```go
func elevatorController(elevatorStatus *ElelvatorStatus, requestOne chan int, requestTwo chan int, requestThree chan int) {
	for true {

    }
}
```

`elevatorController`上来就是一个`for true`的没有退出条件的循环，表示会不断处理电梯请求

然后使用了`Go`中搭配协程使用的`select - case`语句：

```go
select {
	case request := <-requestOne:
		// ...
	case request := <-requestTwo:
		// ...
	case request := <-requestThree:
		// ...
}
```

意思是，`select`会尝试从所有case的信道中取出数据并执行相应代码。如果没有任何一个case的信道有数据可取出，就阻塞直到任何一个有；如果有多个信道有数据，任选一个。

以1楼的电梯的请求为例：

```go
case request := <-requestOne:
    senderPosition := 1
    elevatorStatus.lock.Lock()
    if elevatorStatus.position != senderPosition {
        elevatorStatus.isMoving = true
        elevatorStatus.lock.Unlock()
        for true {
            elevatorStatus.lock.Lock()
            if elevatorStatus.position == senderPosition {
                elevatorStatus.isMoving = false
                elevatorStatus.lock.Unlock()
                break
            }
            if elevatorStatus.position > senderPosition {
                elevatorStatus.isUp = false
                elevatorStatus.lock.Unlock()
                time.Sleep(2 * time.Second)
                elevatorStatus.lock.Lock()
                elevatorStatus.position--
                elevatorStatus.lock.Unlock()
            }
            if elevatorStatus.position < senderPosition {
                elevatorStatus.isUp = true
                elevatorStatus.lock.Unlock()
                time.Sleep(2 * time.Second)
                elevatorStatus.lock.Lock()
                elevatorStatus.position++
                elevatorStatus.lock.Unlock()
            }
        }
    } else {
        elevatorStatus.lock.Unlock()
    }
    time.Sleep(3 * time.Second)
    for true {
        elevatorStatus.lock.Lock()
        elevatorStatus.isMoving = true
        if elevatorStatus.position == request {
            elevatorStatus.isMoving = false
            elevatorStatus.lock.Unlock()
            break
        }
        if elevatorStatus.position > request {
            elevatorStatus.isUp = false
            elevatorStatus.lock.Unlock()
            time.Sleep(2 * time.Second)
            elevatorStatus.lock.Lock()
            elevatorStatus.position--
            elevatorStatus.lock.Unlock()
        }
        if elevatorStatus.position < request {
            elevatorStatus.isUp = true
            elevatorStatus.lock.Unlock()
            time.Sleep(2 * time.Second)
            elevatorStatus.lock.Lock()
            elevatorStatus.position++
            elevatorStatus.lock.Unlock()
        }
    }
```

中控需要写数据，所以全程都用的**写锁**

先判断电梯的`positon`和`senderPosition`是否一致，如果不一致则移动电梯。表示移动电梯的延时过程中，中控**没有上锁**，目的是为了各楼层的电梯面板能够用读锁来获取电梯状态（运行中、运行的方向）

然后延时3秒，模拟乘客上电梯的过程

最后将电梯移动向目的地。和上一个移动过程相同，移动中不会上锁，目的是让各楼层电梯面板能获取到电梯状态并更新显示

### 楼层电梯面板

以2层的面板为例：

```go
func floodTwo(elevator fyne.App, elevatorStatus *ElelvatorStatus, requestChan chan int) {
	winTwo := elevator.NewWindow("2")
	winTwo.Resize(fyne.NewSize(200, 200))
	icon, _ := fyne.LoadResourceFromPath("Icon.png")
	winTwo.SetIcon(icon)
	titleLabel := widget.NewLabel("Flood 2 Panel")
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}
	statusLable := widget.NewLabel("")
	go func() {
		for true {
			time.Sleep(time.Duration(1) * time.Second)
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
			statusLable.SetText(tmpLabel)
			statusLable.Refresh()
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
	content := container.NewVBox(titleLabel, statusLable, goto1Button, goto3Button)
	winTwo.SetContent(content)
	winTwo.Show()
}
```

首先进行了一些图形界面和控件的绘制

然后`go`了一个协程用于更新电梯状态。该函数每1s尝试一次用**读锁**以获取电梯信息并更新到视图中。因为使用的是读锁，所以各个楼层可以同时上读锁而无需阻塞等待

创建了两个按钮，分别表示前往1楼和3楼，闭包函数的功能是弹出一个窗口、向信道中送入表示目的地的数据

得益于Go的信道设计，没有缓冲区的信道中，发送者在送入数据后会阻塞直到有线程将数据接受。所以如果在电梯移动过程中，有线程产生新的请求，该请求会使得楼层面板阻塞（更新视图的线程是另一个线程所以不会被阻塞，视图可以正常更新），等到当前请求执行完毕、轮到自己时才接触阻塞。

