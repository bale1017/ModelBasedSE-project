package main

import (
	"fmt"
	"time"
)

type rectangle struct {
	length int
	width  int
}

type square struct {
	length int
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (s square) area() int {
	return s.length * s.length
}

func (r *rectangle) scale(x int) {
	r.length = r.length * x
	r.width = r.width * x
}

func (s *square) scale(x int) {
	s.length = s.length * x
}

type shape interface {
	area() int
}

func area_Rec(r rectangle) int {
	return r.length * r.width
}

func area_Sq(s square) int {
	return s.length * s.length
}

//
// RT
//

func area_Lookup(x interface{}) int {
	var y int

	switch v := x.(type) {
	case square:
		y = area_Sq(v)
	case rectangle:
		y = area_Rec(v)
	}
	return y

}

func sumArea_Lookup(x, y interface{}) int {
	return area_Lookup(x) + area_Lookup(y)
}

//
// DT
//

type shape_Value struct {
	val  interface{}
	area func(interface{}) int
}

func sumArea_Dict(x, y shape_Value) int {
	return x.area(x.val) + y.area(y.val)
}

func main() {
	fmt.Println("Hello, 世界")

	var r rectangle = rectangle{1, 2}
	var s square = square{3}

	//
	// (1.)
	//

	N := 1000000000

	var t int64 = 0
	start := time.Now()
	fmt.Printf("Start @ %s\n", start)
	for i := 0; i < N; i++ {
		sumArea_Lookup(r, s)
	}
	end := time.Now()
	fmt.Printf("End @ %s\n", end)
	diff := time.Since(start)
	fmt.Printf("diff=%d\n", diff)
	milli := diff.Nanoseconds()
	fmt.Printf("milli=%d\n", milli)
	var m float64 = float64(t) / float64(N)
	fmt.Printf("Mean time after %d rounds: %f", N, m)

	//
	// DT
	//
	area_Rec_Wrapper := func(v interface{}) int {
		return area_Rec(v.(rectangle))

	}

	area_Sq_Wrapper := func(v interface{}) int {
		return area_Sq(v.(square))

	}

	rDictShape := shape_Value{r, area_Rec_Wrapper}

	sDictShape := shape_Value{s, area_Sq_Wrapper}

	sumArea_Dict(rDictShape, sDictShape)
}
