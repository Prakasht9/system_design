package main

import "fmt"

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning Tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning Tv Off")
}
