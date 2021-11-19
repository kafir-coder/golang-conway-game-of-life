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
