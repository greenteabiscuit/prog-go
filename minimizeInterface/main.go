package main

import "math/rand"

type flFunc func(fl float64) float64

func main() {
	var cubicFunc flFunc

	cubicFunc = func(x float64) float64 {
		return x * (x - 1) * (x - 2)
	}

	frg := &FloatRandomGenerator{
		a: 0,
		b: 2,
		n: 7,
	}

	x, f := minimizeUsingInterface(cubicFunc, frg)
	println(x, f)
}

const minimumFloat = -1000000

type FloatRandomGenerator struct {
	a int
	b int
	n int
}

func (f *FloatRandomGenerator) Next() float64 {
	if f.n > 0 {
		f.n -= 1
		return float64(f.a) + rand.Float64()*float64(f.b-f.a)
	}
	return minimumFloat // 他思いつかない
}

type floatGenerator interface {
	Next() float64
}

// frg の型をfloatGeneratorに変更

func minimizeUsingInterface(f flFunc, frg floatGenerator) (x, res float64) {
	var minF, minX float64
	minF = 100.0 // ここもテキトーに大きい値
	for {
		x := frg.Next()
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
