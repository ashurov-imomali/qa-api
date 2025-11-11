package repository

import (
	"errors"
	"github.com/ashurov-imomali/qa-api/internal/models"
	"gorm.io/gorm"
)

func (r *repo) CreateAnswer(answer *models.Answer) error {
	return r.db.Create(answer).Error
}

func (r *repo) GetAnswerByID(id int) (*models.Answer, bool, error) {
	var result models.Answer
	if err := r.db.First(&result, id).Error; err != nil {
		return nil, errors.Is(gorm.ErrRecordNotFound, err), err
	}
	return &result, false, nil
}

func (r *repo) DeleteAnswer(id int) error {
	return r.db.Delete(models.Answer{}, id).Error
}
