package domain

type PageUsecase interface {
}

type PageRepo interface {
}

type Page struct {
	Key      PageKey
	Articles []Article
	NextPage PageKey
}

type PageKey string

type Article struct {
	Title   string
	Content string
}
