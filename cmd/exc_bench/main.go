package main

import "math/rand"

var (
	randomNoException   *rand.Rand
	randomWithException *rand.Rand
)

func init() {
	randomNoException = rand.New(rand.NewSource(420))
	randomWithException = rand.New(rand.NewSource(420))
}

func noException() bool {
	return randomNoException.Float64() > 0.5
}

func withException() bool {
	if randomWithException.Float64() > 0.5 {
		return true
	}

	defer func() bool {
		recover()
		return false
	}()

	panic("Random issue")
}
