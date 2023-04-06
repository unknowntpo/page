package domain

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/exp/rand"

	"github.com/oklog/ulid/v2"
)

type PageUsecase interface {
	GetPage(ctx context.Context, pageKey PageKey) (Page, error)
	GetHead(ctx context.Context, listKey ListKey) (PageKey, error)
	SetPage(ctx context.Context, userID int64, listKey ListKey, page Page) (PageKey, error)
	NewList(ctx context.Context, userID int64, listKey ListKey) error
}

type PageRepo interface {
	GetPage(ctx context.Context, pageKey PageKey) (Page, error)
	GetHead(ctx context.Context, userID int64, listKey ListKey) (PageKey, error)
	SetPage(ctx context.Context, userID int64, listkey ListKey, page Page) (PageKey, error)
	NewList(ctx context.Context, userID int64, listKey ListKey) error
}

type Page struct {
	Key      PageKey
	Content  string
	NextPage PageKey
}

func (p *Page) SetContent(c string) {
	p.Content = c
}

type PageKey string
type PageMetaKey string
type ListKey string

type Article struct {
	Title   string
	Content string
}

func GeneratePageKey() PageKey {
	now := time.Now()
	entropy := rand.New(rand.NewSource(uint64(now.UnixNano())))
	ms := ulid.Timestamp(now)
	ulid, err := ulid.New(ms, entropy)
	if err != nil {
		panic(err)
	}
	return PageKey("page:" + ulid.String())
}

func BuildRedisPageKeyStr(pageKey PageKey) string {
	return "page:" + string(pageKey)
}

func GenerateListKeyByUserID(listKey ListKey, userID int64) ListKey {
	return ListKey(fmt.Sprintf("pageList:%s:%d", listKey, userID))
}

func GenerateListMetaKeyByUserID(listKey ListKey, userID int64) PageMetaKey {
	return PageMetaKey(fmt.Sprintf("listMeta:%s:%d", listKey, userID))
}

const (
	PersonalBoardKey ListKey       = "personal"
	DefaultPageTTL   time.Duration = 24 * time.Hour
)
