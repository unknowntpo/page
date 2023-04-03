package redis

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/unknowntpo/page/internal/domain"

	"github.com/redis/go-redis/v9"
)

type pageRepoImpl struct {
	// any fields needed for implementation
	client *redis.Client
}

func NewPageRepo(c *redis.Client) domain.PageRepo {
	return &pageRepoImpl{client: c}
}

func (r *pageRepoImpl) GetPage(ctx context.Context, pageKey domain.PageKey) (domain.Page, error) {
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
	// member := redis.Z{
	// 	Score:  float64(time.Now().Unix()),
	// 	Member: pageKey,
	// }
	// if err := r.client.ZAdd(ctx, string(listKeyByUser), member).Err(); err != nil {
	// 	return errors.Wrap(err, "failed on r.client.ZAdd for listKeyByUser(%s)", listKeyByUser)
	// }
	// Get Latest Key, then Set page.Next to this PageKey,
	// Finally set it to list:<ListKey> and Hash
	// if err := r.client.ZMScore(ctx, string(pageKey), p, domain.DefaultPageTTL).Err(); err != nil {
	// 	return errors.Wrap(err, "failed on r.client.ZAdd for listKeyByUser(%s)", listKeyByUser)
	// }

	// if err := r.client.Set(ctx, string(pageKey), p, domain.DefaultPageTTL).Err(); err != nil {
	// 	return errors.Wrap(err, "failed on r.client.ZAdd for listKeyByUser(%s)", listKeyByUser)
	// }
	// return nil

	// pageList
	// pageMeta: stores key, head, tail, nextCandidate
	// pageKey -> data

	// Watch the sorted set
	if err := r.client.Watch(ctx, func(tx *redis.Tx) error {
		_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			list, err := addPageToList(pipe, listKey)
			if err != nil {
				return err
			}

			next, err := list.GetNextCandidate()
			if err != nil {
				return err
			}

			p.NextPage = next

			if err := list.SetPage(p); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
		return nil

		// if _, err := tx.TxPipelined(ctx, func(pipe redis.Pipeline) error {
		// 	_ = pipe

		// 	list, err := createListIfNotExist(pipe)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	next, err := list.GetNextCandidate()
		// 	if err != nil {
		// 		return err
		// 	}

		// 	p.NextPage = next

		// 	if err := list.SetPage(p); err != nil {
		// 		return err
		// 	}

		// 	return nil
		// }); err != nil {
		// 	return err
		// }
	}, string(listKeyByUser), string(pageKey)); err != nil {
		return err
	}
	return nil
}

func addPageToList(ctx context.Context, pipe redis.Pipeliner, pageKey domain.PageKey, listKeyByUser domain.ListKey) (list, error) {
	member := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: pageKey,
	}
	if err := pipe.ZAdd(ctx, string(listKeyByUser), member).Err(); err != nil {
		return list{}, errors.Wrap(err, "failed on r.client.ZAdd for listKeyByUser(%s)", listKeyByUser)
	}
}

// list represent the list in redis,
type list struct {
	Key           domain.ListKey
	Head          domain.PageKey
	Tail          domain.PageKey
	NextCandidate domain.PageKey
}

func (l *list) SetPage(pipe *redis.Pipeliner, pageKey domain.PageKey) error {
	//
	return nil
}

//
// list of pages in this list
// type pageList struct{}
// type pageMeta struct{}
