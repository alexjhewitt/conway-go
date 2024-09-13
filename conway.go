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

func (u Universe) Neighbors(x, y int) int {
	count := 0
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i == x && j == y {
				continue
			} else if u.Alive(j, i) {
				fmt.Println(j, i)
				count += 1
			}
		}
	}
	return count
}

func (u Universe) Next(x, y int) bool {
	isAlive := u[y][x]
	neighbors := u.Neighbors(x, y)
	switch {
	case isAlive && neighbors < 2:
		return false
	case isAlive && (neighbors == 2 || neighbors == 3):
		return true
	case isAlive && neighbors > 3:
		return false
	case !isAlive && neighbors == 3:
		return true
	default:
		return false
	}
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

func NewUniverse() Universe {
	matrix := make(Universe, height)
	for i := range matrix {
		matrix[i] = make([]bool, width)
	}
	return matrix
}

func main() {
	universe := NewUniverse()
	universe.Seed()
	universe.Show()
	fmt.Println(universe.Neighbors(0, 0))
}
