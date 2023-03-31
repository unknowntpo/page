package domain

type PageUsecase interface {
	GetPage(pageKey PageKey) (Page, error)
	GetHead(listKey ListKey) (PageKey, error)
	SetPage(Page) error
}

type PageAPI interface {
	GetPage(pageKey PageKey) (Page, error)
	GetHead(listKey ListKey) (PageKey, error)
	SetPage(Page) error
}

type PageRepo interface {
	GetPage(pageKey PageKey) (Page, error)
	GetHead(listKey ListKey) (PageKey, error)
	SetPage(Page) error
}

type Page struct {
	Key      PageKey
	Articles []Article
	NextPage PageKey
}

type PageKey string
type ListKey string

type Article struct {
	Title   string
	Content string
}
