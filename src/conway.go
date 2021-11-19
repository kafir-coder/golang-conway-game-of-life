package main

import (
	"fmt"
	"math/rand"
)

/*

	Author: Caio Tony
	Email: caiotony16@gmail.com
	Country: Angola
	Github: kafir-coder

*/

type Universe [][]bool

const (
	width  = 10
	height = 10
)

/*
	Time Complexity: O(n)
*/
func NewUniverse() Universe {
	universe := make([][]bool, width)
	for k, _ := range universe {
		universe[k] = make([]bool, height)
	}
	return universe
}

/*
	Time Complexity: O(n^2)
*/
func (u Universe) Show() {
	for k, _ := range u {
		for _, v1 := range u[k] {
			if v1 == false {
				fmt.Printf("⛔")
			} else {
				fmt.Printf("♎")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

/*
	Time Complexity: O(n)
*/
func (u Universe) Seed() {

	universe_size := width * height
	universe_size_genesis_population := (universe_size * 25) / 100

	counter := 0
	for counter < universe_size_genesis_population {
		x_axis_random := rand.Intn(width - 1)
		y_axis_random := rand.Intn(height - 1)
		u[x_axis_random][y_axis_random] = true
		counter++
	}

}

/*
	Time Complexity: O(1)
*/
func wrapUp(x, y int) (int, int) {
	if x < 0 {
		x = width - 1
	}
	if y < 0 {
		y = height - 1
	}

	if x >= width {
		x = 0
	}
	if y >= height {
		y = 0
	}
	return x, y
}

/*
	Time Complexity: O(1)
*/
func (u Universe) Alive(x, y int) bool {
	x, y = wrapUp(x, y)
	return u[x][y]
}

/*
	Time Complexity: O(n^2)
*/
func (u Universe) Neighbors(x, y int) int {

	counter := 0
	for i := 0; i < 3; i++ {
		for n := 0; n < 3; n++ {
			x_axis := x + 1 - i
			y_axis := y + 1 - n
			l, c := wrapUp(x_axis, y_axis)
			if u.Alive(l, c) {
				counter++
			}
		}
	}
	if u.Alive(x, y) {
		counter--
	}
	return counter
}

/*
	Time Complexity: O(1)
*/
func (u Universe) Next(x, y int) bool {
	neighbors := u.Neighbors(x, y)
	var willLive bool
	switch u.Alive(x, y) {
	case true:
		if neighbors < 2 {
			willLive = false
		} else if neighbors == 2 || neighbors == 3 {
			willLive = true
		} else if neighbors > 3 {
			willLive = false
		}
	case false:
		if neighbors == 3 {
			willLive = true
		}
	}
	return willLive
}

/*
	Time Complexity: O(n^2)
*/
func (u Universe) Step(future Universe) {
	for k, _ := range u {
		for k1, _ := range u[k] {
			future[k][k1] = u.Next(k, k1)
		}
	}
}

func main() {
	var universe Universe = NewUniverse()
	var next_universe Universe = NewUniverse()
	universe.Seed()
	i := 0
	for i < 101 {
		universe.Show()
		universe.Step(next_universe)
		next_universe, universe = universe, next_universe
		i++
	}

}
