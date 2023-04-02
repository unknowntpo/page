package redis

import "github.com/unknowntpo/page/internal/domain"

type pageRepoImpl struct {
	// any fields needed for implementation
}

func NewPageRepo() domain.PageRepo {
	return &pageRepoImpl{}
}

func (r *pageRepoImpl) GetPage(pageKey domain.PageKey) (domain.Page, error) {
	// implementation
	return domain.Page{}, nil
}

func (r *pageRepoImpl) GetHead(listKey domain.ListKey) (domain.PageKey, error) {
	// implementation
	return domain.PageKey("asdf"), nil
}

func (r *pageRepoImpl) SetPage(p domain.Page) error {
	// implementation
	return nil
}
