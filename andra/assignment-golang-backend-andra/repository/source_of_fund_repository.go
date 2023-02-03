package repository

import (
	"errors"

	"git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/assignment-golang-backend/utils/errors"
	"gorm.io/gorm"
)

type SourceOfFundRepository interface {
	GetById(id uint) (*entity.SourceOfFund, error)
}

type sourceOfFundRepositoryImpl struct {
	db *gorm.DB
}

type SourceOfFundRConfig struct {
	DB *gorm.DB
}

func NewSourceOfFundRepository(cfg *SourceOfFundRConfig) SourceOfFundRepository {
	return &sourceOfFundRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *sourceOfFundRepositoryImpl) GetById(id uint) (*entity.SourceOfFund, error) {
	var res entity.SourceOfFund
	err := r.db.First(&res, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrSourceOfFundNotFound
		}
		return nil, err
	}

	return &res, nil
}
