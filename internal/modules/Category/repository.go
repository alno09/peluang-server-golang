package category

import (
	"peluang-server/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(con *gorm.DB) domain.CategoryRepository {
	return &repository{
		db: con,
	}
}

// FindAll implements domain.CategoryRepository.
func (r *repository) FindAll() ([]domain.Category, error) {
	panic("unimplemented")
}
