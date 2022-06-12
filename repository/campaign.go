package repository

import (
	"basicDB/entity"
	"gorm.io/gorm"
)

type CampaignInterface interface {
	Create(model entity.Campaign) (entity.Campaign, error)
	Update(model entity.Campaign) (entity.Campaign, error)
}
type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *campaignRepository {
	return &campaignRepository{db}
}

func (r *campaignRepository) Create(model entity.Campaign) (entity.Campaign, error) {
	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r *campaignRepository) Update(model entity.Campaign) (entity.Campaign, error) {

	return model, nil
}
