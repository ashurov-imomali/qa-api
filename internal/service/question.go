package service

import (
	"errors"
	"github.com/ashurov-imomali/qa-api/internal/models"
	"net/http"
	"strings"
)

func (s *service) GetQuestions(page, limit int) ([]models.Question, error) {
	//offset, nLimit := 0, 1000
	//if page > 0 && limit > 0 {
	//	offset = (page - 1) * limit
	//	nLimit = limit
	//}

	list, err := s.repo.GetQuestionList()
	if err != nil {
		s.log.Errorf("Failed to fetch questions from database: %v", err)
		return nil, errors.New("internal server error")
	}

	s.log.Infof("Questions fetched successfully (count=%d)", len(list))
	return list, nil
}

func (s *service) CreateQuestion(text string) (*models.Question, int, error) {
	if len(strings.TrimSpace(text)) == 0 {
		s.log.Warnf("Invalid question text provided: %q", text)
		return nil, http.StatusUnprocessableEntity, errors.New("invalid question text")
	}

	question := &models.Question{
		Text: text,
	}

	if err := s.repo.CreateQuestion(question); err != nil {
		s.log.Errorf("Failed to create question: %v", err)
		return nil, http.StatusInternalServerError, errors.New("internal server error")
	}

	s.log.Infof("Question created successfully (id=%d, text=%q)", question.ID, question.Text)
	return question, http.StatusCreated, nil
}

func (s *service) GetQuestionWithAnswers(id int) (*models.QuestionWithAnswers, int, error) {
	result, notFound, err := s.repo.GetQuestionWithAnswers(id)

	if notFound {
		s.log.Warnf("Question not found (id=%d)", id)
		return nil, http.StatusNotFound, errors.New("question not found")
	}

	if err != nil {
		s.log.Errorf("Failed to fetch question with answers (id=%d): %v", id, err)
		return nil, http.StatusInternalServerError, errors.New("internal server error")
	}

	s.log.Infof("Question with answers retrieved successfully (id=%d)", id)
	return result, http.StatusOK, nil
}

func (s *service) DeleteQuestion(id int) error {
	if err := s.repo.DeleteQuestion(id); err != nil {
		s.log.Errorf("Failed to delete question (id=%d): %v", id, err)
		return errors.New("internal server error")
	}

	s.log.Infof("Question deleted successfully (id=%d)", id)
	return nil
}
