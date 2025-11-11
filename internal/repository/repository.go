package repository

import (
	"github.com/ashurov-imomali/qa-api/internal/models"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

type QuestionRepository interface {
	GetQuestionList() ([]models.Question, error)
	CreateQuestion(question *models.Question) error
	GetQuestionWithAnswers(id int) (*models.QuestionWithAnswers, bool, error)
	DeleteQuestion(id int) error
	GetQuestionById(id int) (*models.Question, bool, error)
}

type AnswerRepository interface {
	CreateAnswer(answer *models.Answer) error
	GetAnswerByID(id int) (*models.Answer, bool, error)
	DeleteAnswer(id int) error
}

type Repository interface {
	QuestionRepository
	AnswerRepository
}
