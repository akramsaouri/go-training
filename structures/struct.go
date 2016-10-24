package main

import (
	"fmt"
	"strconv"
)

type Coordonee struct {
	X, Y int
}

func (c *Coordonee) toString() string {
	return "X : " + strconv.Itoa(c.X) + " Y: " + strconv.Itoa(c.Y)
}

func (c *Coordonee) setX(x int) {
	c.X = x
}

func (c *Coordonee) setY(y int) {
	c.Y = y
}

func main() {
	s := Coordonee{1, 2}
	fmt.Println(s.toString())
	s.setX(10)
	fmt.Println(s.toString())
	s.setY(-5)
	fmt.Println(s.toString())
}
