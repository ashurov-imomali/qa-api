package service

import (
	"github.com/ashurov-imomali/qa-api/internal/models"
	"github.com/ashurov-imomali/qa-api/internal/repository"
	"github.com/ashurov-imomali/qa-api/pkg/logger"
)

type service struct {
	repo repository.Repository
	log  logger.Logger
}

type QuestionService interface {
	GetQuestions(page, limit int) ([]models.Question, error)
	CreateQuestion(text string) (*models.Question, int, error)
	GetQuestionWithAnswers(id int) (*models.QuestionWithAnswers, int, error)
	DeleteQuestion(id int) error
}

func NewQuestionService(repo repository.Repository, log logger.Logger) QuestionService {
	return &service{repo: repo, log: log}
}

type AnswerService interface {
	AnswerToQuestion(questionId int, userId, answerText string) (*models.Answer, int, error)
	GetAnswerByID(id int) (*models.Answer, int, error)
	DeleteAnswerByID(id int) error
}

func NewAnswerService(repo repository.Repository, log logger.Logger) AnswerService {
	return &service{repo: repo, log: log}
}
