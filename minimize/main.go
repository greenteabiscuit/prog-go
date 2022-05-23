package main

import (
	"fmt"
	"math/rand"
)

type flFunc func(fl float64) float64

const minimumFloat = -1000000

func main() {
	var cubicFunc flFunc

	cubicFunc = func(x float64) float64 {
		return x * (x - 1) * (x - 2)
	}

	frg := FloatRandomGenerator{
		a: 0,
		b: 2,
		n: 7,
	}

	x, f := minimize(cubicFunc, frg)
	fmt.Println(x, f)
}

func minimize(f flFunc, frg FloatRandomGenerator) (x, res float64) {
	var minF, minX float64
	minF = 100.0 // ここもテキトーに大きい値
	for {
		x := frg.next()
		if x == minimumFloat {
			break
		}
		res := f(x)
		if res < minF {
			minF = res
			minX = x
		}
	}

	return minX, minF
}

type FloatRandomGenerator struct {
	a int
	b int
	n int
}

func (f *FloatRandomGenerator) next() float64 {
	if f.n > 0 {
		f.n -= 1
		return float64(f.a) + rand.Float64()*float64(f.b-f.a)
	}
	return minimumFloat // 他思いつかない
}
