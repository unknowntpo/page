package domain

import (
	"context"
	"encoding/json"
	"time"

	"github.com/unknowntpo/page/pkg/errors"
)

type PageUsecase interface {
	GetPage(ctx context.Context, userID int64, listKey ListKey, pageKey PageKey) (Page, error)
	GetHead(ctx context.Context, userID int64, listKey ListKey) (PageKey, error)
	SetPage(ctx context.Context, userID int64, listKey ListKey, page Page) (PageKey, error)
	NewList(ctx context.Context, userID int64, listKey ListKey) error
}

type PageRepo interface {
	GetPage(ctx context.Context, uesrID int64, listKey ListKey, pageKey PageKey) (Page, error)
	GetHead(ctx context.Context, userID int64, listKey ListKey) (PageKey, error)
	SetPage(ctx context.Context, userID int64, listkey ListKey, page Page) (PageKey, error)
	NewList(ctx context.Context, userID int64, listKey ListKey) error
}

type Page struct {
	Key     PageKey `json:"key"`
	Content string  `json:"content"`
	Next    PageKey `json:"next"`
}

func (p *Page) GetJSONContent() string {
	b, err := json.Marshal(p.Content)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (p *Page) Marshal() []byte {
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return b
}

func (p *Page) String() string {
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(b)
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

const (
	DefaultPageTTL time.Duration = 24 * time.Hour
)

// Errors defined by our domain
var (
	ErrInvalidUserID     = errors.New(errors.BadRequest, "userID should be greater than 0")
	ErrInvalidListKey    = errors.New(errors.BadRequest, "listKey can not be empty")
	ErrListAlreadyExists = errors.New(errors.ResourceAlreadyExist, "list already exist")
	ErrListNotFound      = errors.New(errors.ResourceNotFound, "list not found")
	ErrInternal          = errors.New(errors.Internal, "Internal Server Error")
	ErrPageNotFound      = errors.New(errors.ResourceNotFound, "page not found")
)
