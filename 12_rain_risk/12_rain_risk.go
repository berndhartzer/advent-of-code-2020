package rain_risk

import (
	"strconv"
	"strings"
)

const (
	north   byte = 78
	south   byte = 83
	east    byte = 69
	west    byte = 87
	left    byte = 76
	right   byte = 82
	forward byte = 70
)

type Ferry struct {
	x, y, dist, orientation int
}

func (f *Ferry) MoveNorth(n int) {
	f.y -= n
}

func (f *Ferry) MoveSouth(n int) {
	f.y += n
}

func (f *Ferry) MoveEast(n int) {
	f.x += n
}

func (f *Ferry) MoveWest(n int) {
	f.x -= n
}

func (f *Ferry) MoveLeft(n int) {
	f.orientation -= n
	if f.orientation < 0 {
		f.orientation = 360 - -f.orientation
	}
}

func (f *Ferry) MoveRight(n int) {
	f.orientation += n
	if f.orientation > 360 {
		f.orientation = 0 + (f.orientation - 360)
	}
}

func (f *Ferry) MoveForward(num int) {
	switch {
	case f.orientation <= 45:
		f.MoveNorth(num)
	case f.orientation <= 135:
		f.MoveSouth(num)
	case f.orientation <= 225:
		f.MoveEast(num)
	case f.orientation <= 315:
		f.MoveWest(num)
	default:
		f.MoveNorth(num)
	}
}

func (f *Ferry) Move(in string) {
	move := in[0]

	trim := strings.TrimLeft(in, "NSEWLRF")
	num, _ := strconv.Atoi(trim)

	switch move {
	case north:
		f.MoveNorth(num)
	case south:
		f.MoveSouth(num)
	case east:
		f.MoveEast(num)
	case west:
		f.MoveWest(num)
	case left:
		f.MoveLeft(num)
	case right:
		f.MoveRight(num)
	case forward:
		f.MoveForward(num)
	}
}

func (f *Ferry) Manhattan() int {
	xAbs := f.x
	yAbs := f.y

	if xAbs < 0 {
		xAbs = -xAbs
	}
	if yAbs < 0 {
		yAbs = -yAbs
	}

	return xAbs + yAbs
}

func RainRiskPartOne(instructions []string) int {
	ferry := &Ferry{
		x:           0,
		y:           0,
		dist:        0,
		orientation: 90,
	}

	for _, instruction := range instructions {
		ferry.Move(instruction)
	}

	return ferry.Manhattan()
}
