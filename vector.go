package main

type Vector2d[T ~int | ~uint | ~float32 | ~int64 | ~uint64 | ~float64] struct {
	X, Y T
}
