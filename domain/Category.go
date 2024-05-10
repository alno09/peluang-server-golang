package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ImgURL    string    `json:"img_url" gorm:"default:null"`
	CreatedAt time.Time `json:"created_At" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_At" gorm:"autoUpdateTime"`
}

type CategoryRepository interface {
	FindAll() ([]Category, error)
}

type CategoryService interface {
	GetAllCAtegory() ([]Category, error)
}
