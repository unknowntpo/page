package page

import "github.com/unknowntpo/page/domain"

type pageUsecaseImpl struct {
	// any fields needed for implementation
}

func NewPageUsecase() domain.PageUsecase {
	return &pageUsecaseImpl{}
}

func (u *pageUsecaseImpl) GetPage(pageKey domain.PageKey) (domain.Page, error) {
	// implementation
	return domain.Page{}, nil
}

func (u *pageUsecaseImpl) GetHead(listKey domain.ListKey) (domain.PageKey, error) {
	// implementation
	return domain.PageKey("abc"), nil
}

func (u *pageUsecaseImpl) SetPage(p domain.Page) error {
	// implementation
	return nil
}
