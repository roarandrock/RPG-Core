package combat

import (
	"math/rand"
	"time"
)

//Brng returns a single value for given range
func Brng(r int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	m := r1.Intn(r) + 1
	return m
}
