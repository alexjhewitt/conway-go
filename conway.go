package main

import (
	"fmt"
	"math/rand"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func (u Universe) Alive(x, y int) bool {
	var newX int
	var newY int
	if x > width {
		newX = (x % width) - 1
	} else if x < 0 {
		newX = width + x
	} else {
		newX = x
	}
	if y > height {
		newY = (y % height) - 1
	} else if y < 0 {
		newY = height + y
	} else {
		newY = y
	}
	return u[newY][newX]
}

func NewUniverse() Universe {
	matrix := make(Universe, height)
	for i := range matrix {
		matrix[i] = make([]bool, width)
	}
	return matrix
}

func (u Universe) Seed() {
	for i := range u {
		for j := range u[i] {
			isAlive := rand.Float64() > .75
			if isAlive {
				u[i][j] = true
			}
		}
	}
}

func (u Universe) Show() {
	for i := range u {
		line := ""
		for j := range u[i] {
			if u[i][j] {
				line = line + "*"
			} else {
				line = line + " "
			}
		}
		fmt.Println(line)
	}
}

func main() {
	universe := NewUniverse()
	universe.Seed()
	universe.Show()
}
