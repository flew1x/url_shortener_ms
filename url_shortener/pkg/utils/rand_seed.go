package utils

import (
	"math/rand"
	"time"
)

func SeededRand() *rand.Rand {
	return rand.New(
		rand.NewSource(
			time.Now().UnixNano(),
		),
	)
}
