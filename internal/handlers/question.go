package handlers

import (
	"encoding/json"
	"github.com/ashurov-imomali/qa-api/internal/models"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func (h *Handler) handleQuestions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listQuestions(w, r)
	case http.MethodPost:
		h.createQuestion(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *Handler) handleQuestionByID(w http.ResponseWriter, r *http.Request) {
	cleanPath := path.Clean(r.URL.Path)
	split := strings.Split(cleanPath, "/")
	var idStr string
	if len(split) > 2 {
		idStr = split[2]
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	switch {
	case r.Method == http.MethodPost && len(split) == 4 && split[3] == "answers":
		h.answerToQuestion(w, r, id)
	case r.Method == http.MethodGet:
		h.getQuestion(w, r, id)
	case r.Method == http.MethodDelete:
		h.deleteQuestion(w, r, id)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *Handler) listQuestions(w http.ResponseWriter, r *http.Request) {
	strPage := r.URL.Query().Get("page")
	strLimit := r.URL.Query().Get("limit")
	page, err := strconv.Atoi(strPage)
	if err != nil {
		h.log.Warnf("Error parsing page number to int: %v", err)
		// опционально поэтому продолжаем работу функции
	}
	limit, err := strconv.Atoi(strLimit)
	if err != nil {
		h.log.Warnf("Error parsing limit to int: %v", err)
		// опционально поэтому продолжаем работу функции
	}

	questions, err := h.qs.GetQuestions(page, limit)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, questions)
}

func (h *Handler) createQuestion(w http.ResponseWriter, r *http.Request) {
	var q models.Question
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	created, status, err := h.qs.CreateQuestion(q.Text)
	if err != nil {
		writeError(w, status, err.Error())
		return
	}

	writeJSON(w, status, created)
}

func (h *Handler) getQuestion(w http.ResponseWriter, r *http.Request, id int) {
	q, status, err := h.qs.GetQuestionWithAnswers(id)
	if err != nil {
		writeError(w, status, err.Error())
		return
	}
	writeJSON(w, status, q)
}

func (h *Handler) deleteQuestion(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.qs.DeleteQuestion(id); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
