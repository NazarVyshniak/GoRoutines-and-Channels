package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3
var eatTime = 1 * time.Second
var think = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")

	dine()

}

func dine() {

}
