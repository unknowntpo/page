package page

import (
	"context"

	"github.com/unknowntpo/page/domain"
)

type pageUsecaseImpl struct {
	// any fields needed for implementation
}

func NewPageUsecase() domain.PageUsecase {
	return &pageUsecaseImpl{}
}

func (u *pageUsecaseImpl) GetPage(ctx context.Context, pageKey domain.PageKey) (domain.Page, error) {
	// implementation
	return domain.Page{}, nil
}

func (u *pageUsecaseImpl) GetHead(ctx context.Context, listKey domain.ListKey) (domain.PageKey, error) {
	// implementation
	return domain.PageKey("abc"), nil
}

func (u *pageUsecaseImpl) SetPage(ctx context.Context, userID int64, listKey domain.ListKey, page domain.Page) error {
	// implementation
	return nil
}
