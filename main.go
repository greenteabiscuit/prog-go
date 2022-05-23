package main

import (
	"fmt"
	"math/rand"
)

type G func(f float64)

func main() {
	f := func(g G) {
		g(1.0)
	}
	var g G
	g = func(f float64) {
		fmt.Println(f)
	}
	f(g)

	frg := FloatRandomGenerator{
		a: 1,
		b: 2,
		n: 3,
	}

	for i := 0; i < 3; i++ {
		fmt.Println(frg.next())
	}
}

type FloatRandomGenerator struct {
	a int
	b int
	n int
}

func (f *FloatRandomGenerator) next() float64 {
	if f.n > 0 {
		return float64(f.a) + rand.Float64()*float64(f.b-f.a)
	}
	f.n -= 1
	return -1.0
}

type flFunc func(fl float64) float64

func minimize(f flFunc, frg FloatRandomGenerator) {
	for {
		x := frg.next()
		if x == -1.0 {
			break
		}
		_ = f(x)

	}
}
