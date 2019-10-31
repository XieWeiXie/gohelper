package gohelper

import (
	"math/rand"
	"time"
)

type RandomAction struct {
	generator *rand.Rand
}

func NewRandomAction() *RandomAction {
	return &RandomAction{generator: rand.New(rand.NewSource(time.Now().UnixNano()))}
}
func (R *RandomAction) GeneratorInt31n(n int32) int32 {
	return R.generator.Int31n(n)
}

func (R *RandomAction) GeneratorInt31() int32 {
	return R.generator.Int31()
}
func (R *RandomAction) GeneratorIntn(n int) int {
	return R.generator.Intn(n)
}

func (R *RandomAction) GeneratorInt() int {
	return R.generator.Int()
}
