package lpass

import (
	"fmt"
	"math/rand"
	"time"
)

const randomHashLength = 6

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Generator struct {
	KeyName string
}

func (g Generator) Name() string {
	rand.Seed(time.Now().UnixNano())

	return fmt.Sprintf("%s-%s", g.KeyName, g.randSeq(randomHashLength))
}

func (g Generator) randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
