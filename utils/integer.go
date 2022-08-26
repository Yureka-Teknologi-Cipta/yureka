package utils

import (
	"math/rand"
	"time"
)

type Ints []int
type Ints64 []int64

func (slice Ints) Include(value int) bool {
	for _, item := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func (slice Ints64) Include(value int64) bool {
	for _, item := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func (slice Ints) Unique() []int {
	unique := Ints{}
	for _, v := range slice {
		if !unique.Include(v) {
			unique = append(unique, v)
		}
	}

	return unique
}

func (slice Ints64) Unique() []int64 {
	unique := Ints64{}
	for _, v := range slice {
		if !unique.Include(v) {
			unique = append(unique, v)
		}
	}

	return unique
}

func (slice Ints) Random() int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	n := rand.Int() % len(slice)

	return slice[n]
}

func (slice Ints64) Random() int64 {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	n := rand.Int() % len(slice)

	return slice[n]
}
