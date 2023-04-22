package sharedsecret

import (
	"gorm.io/gorm"
)

type Repository interface {
	Insert(sharedSecrets []SharedSecret) error
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

func (r *repository) Insert(sharedSecrets []SharedSecret) error {
	return r.db.Create(&sharedSecrets).Error
}

func (r *repository) GetAndDelete(clientID string) ([]SharedSecret, error) {
	var res []SharedSecret

	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("client_id = ?", clientID).Find(&res).Error
		if err != nil {
			return err
		}
		return tx.Where("client_id = ?", clientID).Delete(&SharedSecret{}).Error
	})

	if err != nil {
		return nil, err
	}
	return res, nil
}
