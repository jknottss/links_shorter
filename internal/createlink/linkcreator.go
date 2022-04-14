package createlink

import (
	"math/rand"
	"time"
)

const allowChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func CreateLink() string {
	rand.Seed(time.Now().UnixNano())
	shortLink := make([]byte, 10)
	for i := range shortLink {
		shortLink[i] = allowChars[rand.Intn(len(allowChars))]
	}
	return string(shortLink)
}
