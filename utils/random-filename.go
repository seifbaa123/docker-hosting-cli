package utils

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const filenameLength = 16

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomFilename() string {
	b := make([]byte, filenameLength)
	for i := range b {
		b[i] = letterBytes[rng.Intn(len(letterBytes))]
	}

	return string(b)
}
