package utils

import (
	"math/rand"
	"time"
)

type Strings []string

func (slice Strings) Include(value string) bool {
	for _, item := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func (slice Strings) Unique() []string {
	unique := Strings{}
	for _, v := range slice {
		if !unique.Include(v) {
			unique = append(unique, v)
		}
	}

	return unique
}

func (slice Strings) Random() string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	n := rand.Int() % len(slice)

	return slice[n]
}
