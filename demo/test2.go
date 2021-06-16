package main

import "fmt"

type a struct {
	b int
	c int
}

func (a *a) getSum(d int) int {
	return d + a.b + a.c
}

func (a *a) getMul(d int) int {
	return d * a.b * a.c
}

func main() {
	var e a
	e.b = 1
	e.c = 2
	r1 := e.getSum(2)
	r2 := e.getMul(3)
	fmt.Println(r1, r2)
}
