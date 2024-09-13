package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func (u Universe) Alive(x, y int) bool {
	var newX int
	var newY int
	newX = (x + width) % width
	newY = (y + height) % height
	return u[newY][newX]
}

func (u Universe) Neighbors(x, y int) int {
	count := 0
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i == x && j == y {
				continue
			} else if u.Alive(j, i) {
				count += 1
			}
		}
	}
	return count
}

func (u Universe) Next(x, y int) bool {
	neighbors := u.Neighbors(x, y)
	return neighbors == 3 || neighbors == 2 && u.Alive(x, y)
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

func Step(a, b Universe) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Next(j, i)
		}
	}
}

func NewUniverse() Universe {
	matrix := make(Universe, height)
	for i := range matrix {
		matrix[i] = make([]bool, width)
	}
	return matrix
}

func main() {
	universe := NewUniverse()
	nextStage := NewUniverse()
	universe.Seed()
	for {
		Step(universe, nextStage)
		fmt.Print("\033[H")
		universe.Show()
		time.Sleep(time.Second / 30)
		universe, nextStage = nextStage, universe
	}
}
