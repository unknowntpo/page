package mock

import (
	"math/rand"
	"time"

	"github.com/unknowntpo/page/internal/domain"
)

func GenerateDummyArticles(length int) []domain.Article {
	out := make([]domain.Article, 0, length)
	for i := 0; i < length; i++ {
		out = append(out, domain.Article{Title: GenerateRandomString(10), Content: GenerateRandomString(10)})
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
