package main

import "fmt"

// type Transform1 struct {
// }

// func (t Transform1) Apply(in int) float32 {
// 	return float32(in) + 1.2
// }

// type Transform2 struct {
// }

// func (t Transform2) Apply(in float32) float64 {
// 	return float64(in + 1.1)
// }

// type Transform3 struct {
// }

// func (t Transform3) Apply(in float64) int {
// 	return int(100 * (in + 1.1))
// }

// type Transform4 struct {
// }

// func (t Transform4) Apply(in int) []int {
// 	a := []int{}
// 	for ii := 0; ii < in; ii++ {
// 		a = append(a, in/(ii+1))
// 	}
// 	return a
// }

func Transform1(in int) float32 {
	return float32(in) + 1.2
}

func Transform2(in float32) float64 {
	return float64(in + 1.1)
}
func Transform3(in float64) int {
	return int(100 * (in + 1.1))
}

func Transform4(in int) []int {
	a := []int{}
	for ii := 0; ii < in; ii++ {
		a = append(a, in/(ii+1))
	}
	return a
}

func main() {

	// t1 := Transform1{}
	// t2 := Transform2{}
	// t3 := Transform3{}
	// t4 := Transform4{}
	var input int = 1

	p2 := NewTransform(Transform1)
	p3 := AddTransform(p2, Transform2) //.AddTransform(Transform3)
	p4 := AddTransform(p3, Transform3)
	p5 := AddTransform(p4, Transform4)

	ret := p5.Apply(input)

	// p3 := AddTransform(t2.Apply)
	// ret := p2.Apply(a)

	fmt.Println(ret)
}

// interface with type parameter
type TransformStep[T1 any, T2 any] interface {
	Apply(T1) T2
}

type TransformFunc[T1 any, T2 any] func(T1) T2

// type AddFunc[T1 any, T2 any, T3 any] func(TransformFunc[T2, T3]) Step[T1, T3]

type Step[T1 any, T2 any] struct {
	// Apply TransformStep[T1, T2]
	Apply TransformFunc[T1, T2]
	// AddTransform AddFunc[T1,T2, T3]
}

func AddTransform[T1 any, T2 any, T3 any](s Step[T1, T2], t TransformFunc[T2, T3]) Step[T1, T3] {
	var apply = func(a T1) T3 {
		return t(s.Apply(a))
	}

	next := Step[T1, T3]{}
	next.Apply = apply
	return next
}

// func (s Step[T1, T2]) Apply(a T1) T2 {
// 	return s.step(a)
// }

func NewTransform[T1 any, T2 any](t TransformFunc[T1, T2]) Step[T1, T2] {
	s := Step[T1, T2]{
		Apply: t,
	}

	return s
}
