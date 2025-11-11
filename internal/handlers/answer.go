package handlers

import (
	"encoding/json"
	"github.com/ashurov-imomali/qa-api/internal/models"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func (h *Handler) handleAnswerByID(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	strID := strings.TrimPrefix(cleanPath, "/answers/")

	id, err := strconv.Atoi(strID)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getAnswer(w, r, id)
	case http.MethodDelete:
		h.deleteAnswer(w, r, id)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *Handler) getAnswer(w http.ResponseWriter, r *http.Request, id int) {
	answer, status, err := h.as.GetAnswerByID(id)
	if err != nil {
		writeError(w, status, err.Error())
		return
	}
	writeJSON(w, status, answer)
}

func (h *Handler) deleteAnswer(w http.ResponseWriter, r *http.Request, id int) {
	err := h.as.DeleteAnswerByID(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) answerToQuestion(w http.ResponseWriter, r *http.Request, questionId int) {
	var answer models.Answer
	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	newAns, status, err := h.as.AnswerToQuestion(questionId, answer.UserID, answer.Text)
	if err != nil {
		writeError(w, status, err.Error())
		return
	}
	writeJSON(w, status, newAns)
}
