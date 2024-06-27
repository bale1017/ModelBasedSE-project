package main

import (
	"fmt"
)

// ========================
// EXAMPLE 1 IMPLEMENTATION
// ========================

type node[T any] struct {
	val  T
	next *node[T]
}

func mkNode[T any](v T) *node[T] {
	return &node[T]{val: v, next: nil}
}

func insert[T any](n *node[T], v T) {
	n.next = mkNode[T](v)
}

// n > 0
func mkList[T any](n int, v T) *node[T] {

	head := mkNode(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert(current, v)
		current = current.next

	}
	return head

}

func len[T any](n *node[T]) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}

	return i

}

func testNode() {
	n1 := mkList[int](10, 1)
	fmt.Printf("%d\n", len(n1))

	n2 := mkList[bool](10, true)
	fmt.Printf("%d\n", len(n2))

}

//
// Generic translation
//

type node_G struct {
	val  interface{}
	next *node_G
}

func mkNode_G(v interface{}) *node_G {
	return &node_G{val: v, next: nil}
}

func insert_G(n *node_G, v interface{}) {
	n.next = mkNode_G(v)
}

func mkList_G(n int, v interface{}) *node_G {

	head := mkNode_G(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert_G(current, v)
		current = current.next
	}
	return head
}

func len_G(n *node_G) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}

	return i
}

func testNode_generic() {
	n1 := mkList_G(10, 1)
	fmt.Printf("%d\n", len_G(n1))

	n2 := mkList_G(10, true)
	fmt.Printf("%d\n", len_G(n2))
}

//
// Monomorphization
//

type node_int struct {
	val  int
	next *node_int
}

func mkNode_int(v int) *node_int {
	return &node_int{val: v, next: nil}
}

func insert_int(n *node_int, v int) {
	n.next = mkNode_int(v)
}

func mkList_int(n int, v int) *node_int {
	head := mkNode_int(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert_int(current, v)
		current = current.next
	}
	return head
}

func len_int(n *node_int) int {
	i := 0
	for n != nil {
		i++
		n = n.next

	}
	return i
}

type node_bool struct {
	val  bool
	next *node_bool
}

func mkNode_bool(v bool) *node_bool {
	return &node_bool{val: v, next: nil}
}

func insert_bool(n *node_bool, v bool) {
	n.next = mkNode_bool(v)
}

func mkList_bool(n int, v bool) *node_bool {
	head := mkNode_bool(v)
	current := head
	for i := 0; i < n-1; i++ {
		insert_bool(current, v)
		current = current.next
	}
	return head
}

func len_bool(n *node_bool) int {
	i := 0
	for n != nil {
		i++
		n = n.next

	}
	return i
}

func testNode_monomorphization() {
	n1 := mkList_int(10, 1)
	fmt.Printf("%d\n", len_int(n1))

	n2 := mkList_bool(10, true)
	fmt.Printf("%d\n", len_bool(n2))
}

// ========================
// EXAMPLE 2 IMPLEMENTATION
// ========================

func sum[T int | float32](xs []T) T {
	var x T
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x
}

func sum_int(xs []int) int {
	var x int
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x
}

func sum_float32(xs []float32) float32 {
	var x float32
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x
}

func test_sum_monomorphization() {
	xs1 := []int{1, 1, 1}
	s1 := sum_int(xs1)
	fmt.Printf("%d\n", s1)

	xs2 := []float32{1.1, 1.1, 1.1}
	s2 := sum_float32(xs2)
	fmt.Printf("%f\n", s2)
}

// ========================
// EXAMPLE 3 IMPLEMENTATION
// ========================

func swap[T any](x *T, y *T) {
	tmp := *x
	*x = *y
	*y = tmp
}

func swap_int(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func swap_bool(x *bool, y *bool) {
	tmp := *x
	*x = *y
	*y = tmp
}

func swap_G(x, y *interface{}) {
	tmp := *x
	*x = *y
	*y = tmp
}

func test_swap_monomorphization() {
	x := 1
	y := 2
	fmt.Printf("x=%d, y=%d -> ", x, y)
	swap_int(&x, &y)
	fmt.Printf("x=%d, y=%d\n", x, y)

	a := true
	b := false
	fmt.Printf("a=%t, b=%t -> ", a, b)
	swap_bool(&a, &b)
	fmt.Printf("a=%t, b=%t\n", a, b)
}

func main() {
	fmt.Println("=========")
	fmt.Println("EXAMPLE 1")
	fmt.Println("=========")

	fmt.Println("")
	fmt.Println("-------------------")
	fmt.Println("Generic translation")
	fmt.Println("-------------------")
	testNode_generic()

	fmt.Println("")
	fmt.Println("----------------")
	fmt.Println("Monomorphization")
	fmt.Println("----------------")
	testNode_monomorphization()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("=========")
	fmt.Println("EXAMPLE 2")
	fmt.Println("=========")

	fmt.Println("")
	fmt.Println("-------------------")
	fmt.Println("Generic translation")
	fmt.Println("-------------------")
	fmt.Println("Not possible because we don't know the types.")
	fmt.Println("Adding two untyped operators is undefined.")

	fmt.Println("")
	fmt.Println("----------------")
	fmt.Println("Monomorphization")
	fmt.Println("----------------")
	test_sum_monomorphization()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("=========")
	fmt.Println("EXAMPLE 3")
	fmt.Println("=========")

	fmt.Println("")
	fmt.Println("-------------------")
	fmt.Println("Generic translation")
	fmt.Println("-------------------")
	fmt.Println("cannot use &x (value of type *int) as *interface{} value in argument to swap_G: *int does not implement *interface{} (type *interface{} is pointer to interface, not interface")

	fmt.Println("")
	fmt.Println("----------------")
	fmt.Println("Monomorphization")
	fmt.Println("----------------")
	test_swap_monomorphization()
}
