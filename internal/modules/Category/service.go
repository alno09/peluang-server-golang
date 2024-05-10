package category

import "peluang-server/domain"

type service struct {
	cateRepo domain.CategoryRepository
}

func newService(cateRepo domain.CategoryRepository) domain.CategoryService {
	return &service{
		cateRepo: cateRepo,
	}
}

// GetAllCAtegory implements domain.CategoryService.
func (s *service) GetAllCAtegory() ([]domain.Category, error) {
	panic("unimplemented")
}
