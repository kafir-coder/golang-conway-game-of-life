package main

import "fmt"

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
