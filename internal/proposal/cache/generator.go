package cache

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"
const length = 2

type Generator struct {
	cacheSize int
}

func NewGenerator(cacheSize int) *Generator {
	rand.Seed(time.Now().UnixNano())
	return &Generator{cacheSize: cacheSize}
}

func (g Generator) GenName() string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (g Generator) RandomId() int {
	return rand.Intn(g.cacheSize)
}
