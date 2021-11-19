package main

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
