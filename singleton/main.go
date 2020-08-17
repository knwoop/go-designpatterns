package main

import (
	"fmt"
	"sync"
	"time"
)

var instance *TokyoTower
var once sync.Once

type TokyoTower struct {
	color string
}

func (h *TokyoTower) String() string {
	return "Tokyo tower " + h.color
}

//GetInstance returns singleton instance
func GetInstance(color string) *TokyoTower {
	once.Do(func() {
		fmt.Println("new instance")
		time.Sleep(1 * time.Second)
		instance = &TokyoTower{
			color: color,
		}
	})
	return instance
}

func main() {
	ch := make(chan interface{})
	go run(ch, "Red")
	go run(ch, "Blue")
	go run(ch, "Green")
	<-ch
	<-ch
	<-ch

	tower := GetInstance("Yellow")
	fmt.Println(tower)
}

func run(ch chan<- interface{}, color string) {
	tower := GetInstance(color)
	fmt.Println(tower)
	ch <- tower //blocking
}
