package page

import "github.com/unknowntpo/page/domain"

type pageAPIImpl struct {
	// any fields needed for implementation
}

func NewPageAPI() domain.PageAPI {
	return &pageAPIImpl{}
}

func (api *pageAPIImpl) GetPage(pageKey domain.PageKey) (domain.Page, error) {
	// implementation
	return domain.Page{}, nil
}

func (api *pageAPIImpl) GetHead(listKey domain.ListKey) (domain.PageKey, error) {
	// implementation
	return domain.PageKey("dsfsd"), nil
}

func (api *pageAPIImpl) SetPage(p domain.Page) error {
	// implementation
	return nil
}
