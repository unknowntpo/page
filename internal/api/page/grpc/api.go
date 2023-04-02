package page

import (
	"context"

	"github.com/pkg/errors"
	"github.com/unknowntpo/page/internal/domain"
)

type pageAPIImpl struct {
	// any fields needed for implementation
	useCase domain.PageUsecase
}

func NewPageAPI(uCase domain.PageUsecase) domain.PageAPI {
	return &pageAPIImpl{useCase: uCase}
}

func (api *pageAPIImpl) GetPage(ctx context.Context, pageKey domain.PageKey) (domain.Page, error) {
	// implementation
	return domain.Page{}, nil
}

func (api *pageAPIImpl) GetHead(ctx context.Context, listKey domain.ListKey) (domain.PageKey, error) {
	// implementation
	pageKey, err := api.useCase.GetHead(ctx, listKey)
	if err != nil {
		return "", errors.Wrap(err, "failed on api.useCase.GetHead")
	}
	return pageKey, nil
}

func (api *pageAPIImpl) SetPage(ctx context.Context, p domain.Page) error {
	// implementation
	return nil
}
