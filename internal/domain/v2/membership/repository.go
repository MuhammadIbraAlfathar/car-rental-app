package membershipV2

import (
	v2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v2"
	"gorm.io/gorm"
)

type Repository interface {
	FindById(membershipId int) (*v2.Membership, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindById(membershipId int) (*v2.Membership, error) {
	var membership *v2.Membership
	if err := r.db.First(&membership, membershipId).Error; err != nil {
		return nil, err
	}

	return membership, nil
}
