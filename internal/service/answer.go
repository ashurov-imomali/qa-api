package service

import (
	"errors"
	"github.com/ashurov-imomali/qa-api/internal/models"
	"net/http"
	"strings"
)

func (s *service) AnswerToQuestion(questionId int, userId, answerText string) (*models.Answer, int, error) {
	_, notFound, err := s.repo.GetQuestionById(questionId)

	if notFound {
		s.log.Warnf("Question not found (id=%d)", questionId)
		return nil, http.StatusNotFound, errors.New("question not found")
	}

	if err != nil {
		s.log.Errorf("Failed to fetch question from database (id=%d): %v", questionId, err)
		return nil, http.StatusInternalServerError, errors.New("internal server error")
	}

	if len(strings.TrimSpace(answerText)) == 0 {
		s.log.Warnf("Invalid answer text provided (user=%s, questionId=%d)", userId, questionId)
		return nil, http.StatusUnprocessableEntity, errors.New("invalid answer text")
	}

	newAnswer := &models.Answer{
		QuestionID: questionId,
		UserID:     userId,
		Text:       answerText,
	}

	if err := s.repo.CreateAnswer(newAnswer); err != nil {
		s.log.Errorf("Failed to create answer for question (id=%d): %v", questionId, err)
		return nil, http.StatusInternalServerError, errors.New("internal server error")
	}

	s.log.Infof("Answer created successfully (id=%d, questionId=%d, userId=%s)", newAnswer.ID, questionId, userId)
	return newAnswer, http.StatusOK, nil
}

func (s *service) GetAnswerByID(id int) (*models.Answer, int, error) {
	answer, notFound, err := s.repo.GetAnswerByID(id)

	if notFound {
		s.log.Warnf("Answer not found (id=%d)", id)
		return nil, http.StatusNotFound, errors.New("answer not found")
	}

	if err != nil {
		s.log.Errorf("Failed to fetch answer (id=%d): %v", id, err)
		return nil, http.StatusInternalServerError, errors.New("internal server error")
	}

	s.log.Infof("Answer retrieved successfully (id=%d)", id)
	return answer, http.StatusOK, nil
}

func (s *service) DeleteAnswerByID(id int) error {
	if err := s.repo.DeleteAnswer(id); err != nil {
		s.log.Errorf("Failed to delete answer (id=%d): %v", id, err)
		return errors.New("internal server error")
	}

	s.log.Infof("Answer deleted successfully (id=%d)", id)
	return nil
}
