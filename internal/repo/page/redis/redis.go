package redis

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/unknowntpo/page/domain"

	"github.com/redis/go-redis/v9"
)

type pageRepoImpl struct {
	// any fields needed for implementation
	client *redis.Client
}

func NewPageRepo(c *redis.Client) domain.PageRepo {
	return &pageRepoImpl{client: c}
}

func (r *pageRepoImpl) GetPage(pageKey domain.PageKey) (domain.Page, error) {
	// implementation

	// if page doesn't exist, return PageDoesNotExist error
	return domain.Page{}, nil
}

func (r *pageRepoImpl) GetHead(ctx context.Context, listKey domain.ListKey) (domain.PageKey, error) {
	// implementation

	// return the first element of the page range from now() - 1 Day to now()
	return domain.PageKey("asdf"), nil
}

func (r *pageRepoImpl) SetPage(
	ctx context.Context,
	userID int64,
	listKey domain.ListKey,
	p domain.Page,
) error {
	// implementation
	listKeyByUser := domain.GenerateListKeyByUserID(listKey, userID)
	pageKey := domain.GeneratePageKey()
	member := redis.Z{
		Score:  float64(time.Now().UnixNano()),
		Member: pageKey,
	}
	if err := r.client.ZAdd(ctx, string(listKeyByUser), member).Err(); err != nil {
		return errors.Wrap(err, "failed on r.client.ZAdd for listKeyByUser(%s)", listKeyByUser)
	}
	// Get Latest Key, then Set page.Next to this PageKey,
	// Finally set it to list:<ListKey> and Hash
	if err := r.client.ZMScore(ctx, string(pageKey), p, domain.DefaultPageTTL).Err(); err != nil {
		return errors.Wrap(err, "failed on r.client.ZAdd for listKeyByUser(%s)", listKeyByUser)
	}

	if err := r.client.Set(ctx, string(pageKey), p, domain.DefaultPageTTL).Err(); err != nil {
		return errors.Wrap(err, "failed on r.client.ZAdd for listKeyByUser(%s)", listKeyByUser)
	}
	return nil
}
