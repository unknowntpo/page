package domain

import (
	"math/rand"
	"time"
)

func GenerateDummyArticles(length int) []Article {
	out := make([]Article, 0, length)
	for i := 0; i < length; i++ {
		out = append(out, Article{Title: GenerateRandomString(10), Content: GenerateRandomString(10)})
	}
	return out
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
