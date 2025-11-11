package handlers

import "net/http"

func (h *Handler) RegisterRouters(mux *http.ServeMux) {
	mux.HandleFunc("/questions", h.handleQuestions)
	mux.HandleFunc("/questions/", h.handleQuestionByID)
	mux.HandleFunc("/answers/", h.handleAnswerByID)
}
