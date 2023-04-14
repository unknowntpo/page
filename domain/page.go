package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/unknowntpo/page/pkg/errors"

	"github.com/oklog/ulid/v2"
	"golang.org/x/exp/rand"
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

// getRedisPartitionKey generates given redis partition key by
// userID, listKey, and have format {<listKey>:<userID>}
// NOTE: This format can not be changed, or every key will be migrated to other Redis Node.
func getRedisPartitionKey(listKey ListKey, userID int64) string {
	return fmt.Sprintf("{%s:%d}", listKey, userID)
}

// GeneratePageKeyByListKeyUserID generates given redis key by
// userID, listKey, and have format page:{<listKey>:<userID>}:<ulid>
func GeneratePageKeyByListKeyUserID(listKey ListKey, userID int64, now time.Time) PageKey {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(uint64(now.UnixNano()))), 0)
	ms := ulid.Timestamp(now)
	ulid, err := ulid.New(ms, entropy)
	if err != nil {
		panic(err)
	}
	return PageKey(fmt.Sprintf("page:%s:%s", getRedisPartitionKey(listKey, userID), ulid.String()))
}

// GenerateListKeyByUserID generates pageList (sorted set) redis key by
// listKey, userID, and have format pageList:{<listKey>:<userID>}
func GenerateListKeyByUserID(listKey ListKey, userID int64) ListKey {
	return ListKey(fmt.Sprintf("pageList:%s", getRedisPartitionKey(listKey, userID)))
}

// GenerateListKeyByUserID generates pageList (sorted set) redis key by
// listKey, userID, and have format pageList:{<listKey>:<userID>}
func GenerateListMetaKeyByUserID(listKey ListKey, userID int64) PageMetaKey {
	return PageMetaKey(fmt.Sprintf("listMeta:%s", getRedisPartitionKey(listKey, userID)))
}

const (
	PersonalBoardKey ListKey       = "personal"
	DefaultPageTTL   time.Duration = 24 * time.Hour
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
