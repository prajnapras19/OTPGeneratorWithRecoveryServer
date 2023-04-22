package sharedsecret

import (
	"gorm.io/gorm"
)

type Repository interface {
	Insert(sharedSecret []SharedSecret) error
	GetAndDelete(clientID string) ([]SharedSecret, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(sharedSecret []SharedSecret) error {
	// TODO
	return nil
}

func (r *repository) GetAndDelete(clientID string) ([]SharedSecret, error) {
	// TODO
	return nil, nil
}
