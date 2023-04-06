package page

import (
	"context"

	"github.com/unknowntpo/page/domain"
	"github.com/unknowntpo/page/pkg/errors"
)

type pageUsecaseImpl struct {
	repo domain.PageRepo
}

func NewPageUsecase(r domain.PageRepo) domain.PageUsecase {
	return &pageUsecaseImpl{repo: r}
}

func (u *pageUsecaseImpl) NewList(ctx context.Context, userID int64, listKey domain.ListKey) error {
	if err := u.repo.NewList(ctx, userID, listKey); err != nil {
		switch {
		case errors.KindIs(err, errors.ResourceAlreadyExist):
			return errors.Wrap(errors.ResourceAlreadyExist, "resource already exist", err)
		default:
			return errors.Wrap(errors.Internal, "failed on u.repo.NewList", err)
		}
	}
	return nil
}

func (u *pageUsecaseImpl) GetPage(ctx context.Context, pageKey domain.PageKey) (domain.Page, error) {
	p, err := u.repo.GetPage(ctx, pageKey)
	if err != nil {
		switch {
		case errors.KindIs(err, errors.ResourceNotFound):
			return domain.Page{}, errors.Wrap(errors.ResourceNotFound, "resource not found", err)
		default:
			return domain.Page{}, errors.Wrap(errors.Internal, "failed on u.repo.GetPage", err)
		}
	}
	return p, nil
}

func (u *pageUsecaseImpl) GetHead(ctx context.Context, userID int64, listKey domain.ListKey) (domain.PageKey, error) {
	headKey, err := u.repo.GetHead(ctx, userID, listKey)
	if err != nil {
		switch {
		// case errors.KindIs(err, errors.ResourceNotFound):
		//	return errors.Wrap(errors.ResourceNotFound, "resource not found", err)
		default:
			return "", errors.Wrap(errors.Internal, "failed on u.repo.GetHead", err)
		}
	}
	return headKey, nil
}

func (u *pageUsecaseImpl) SetPage(ctx context.Context, userID int64, listKey domain.ListKey, page domain.Page) (domain.PageKey, error) {
	pageKey, err := u.repo.SetPage(ctx, userID, listKey, page)
	if err != nil {
		switch {
		case errors.KindIs(err, errors.ResourceNotFound):
			return "", errors.Wrap(errors.ResourceNotFound, "resource not found", err)
		default:
			return "", errors.Wrap(errors.Internal, "failed on u.repo.SetPage", err)
		}
	}
	return pageKey, nil
}
