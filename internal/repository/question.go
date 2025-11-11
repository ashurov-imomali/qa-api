package repository

import (
	"errors"
	"github.com/ashurov-imomali/qa-api/internal/models"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func (r *repo) GetQuestionList() ([]models.Question, error) {
	var result []models.Question
	//return result, r.db.Model(&models.Question{}).Limit(limit).Offset(offset).Find(&result).Error
	return result, r.db.Find(&result).Error

}

func (r *repo) CreateQuestion(question *models.Question) error {
	return r.db.Create(question).Error
}

func (r *repo) GetQuestionWithAnswers(id int) (*models.QuestionWithAnswers, bool, error) {
	var result models.QuestionWithAnswers

	gr := errgroup.Group{}
	gr.Go(func() error {
		var question models.Question
		if err := r.db.First(&question, id).Error; err != nil {
			return err
		}
		result.Question = question
		return nil
	})

	gr.Go(func() error {
		var answers []models.Answer
		if err := r.db.Where("question_id=?", id).Find(&answers).Error; err != nil {
			return err
		}
		result.Answers = answers
		return nil
	})

	if err := gr.Wait(); err != nil {
		return nil, errors.Is(err, gorm.ErrRecordNotFound), err
	}

	return &result, false, nil
}

func (r *repo) DeleteQuestion(id int) error {
	return r.db.Delete(&models.Question{}, id).Error
}

func (r *repo) GetQuestionById(id int) (*models.Question, bool, error) {
	var result models.Question
	if err := r.db.First(&result, id).Error; err != nil {
		return nil, errors.Is(gorm.ErrRecordNotFound, err), err
	}
	return &result, false, nil
}
