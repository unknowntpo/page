package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type PageUsecase interface {
	GetPage(ctx context.Context, pageKey PageKey) (Page, error)
	GetHead(ctx context.Context, listKey ListKey) (PageKey, error)
	SetPage(ctx context.Context, page Page) error
}

type PageAPI interface {
	GetPage(ctx context.Context, pageKey PageKey) (Page, error)
	GetHead(ctx context.Context, listKey ListKey) (PageKey, error)
	SetPage(ctx context.Context, page Page) error
}

type PageRepo interface {
	GetPage(ctx context.Context, pageKey PageKey) (Page, error)
	GetHead(ctx context.Context, listKey ListKey) (PageKey, error)
	SetPage(ctx context.Context, listkey ListKey, page Page) error
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

func GeneratePageKey() PageKey {
	return PageKey(uuid.NewString())
}

func GenerateListKeyByUserID(listKey ListKey, userID int64) ListKey {
	return ListKey(fmt.Sprintf("%s:%d", listKey, userID))
}

const (
	PersonalBoardKey ListKey       = "personal"
	DefaultPageTTL   time.Duration = 24 * time.Hour
)
