package handlers

import (
	"github.com/ashurov-imomali/qa-api/internal/service"
	"github.com/ashurov-imomali/qa-api/pkg/logger"
)

type Handler struct {
	qs  service.QuestionService
	as  service.AnswerService
	log logger.Logger
}

func New(qs service.QuestionService, as service.AnswerService, l logger.Logger) *Handler {
	return &Handler{
		as:  as,
		qs:  qs,
		log: l,
	}
}
