package models

import "time"

type Question struct {
	ID        int       `json:"id" gorm:"primaryKey;column:id"`
	Text      string    `json:"text" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

type Answer struct {
	ID         int       `json:"id" gorm:"primaryKey;column:id"`
	QuestionID int       `json:"question_id" gorm:"not null;column:question_id"`
	UserID     string    `json:"user_id" gorm:"type:text;not null;column:user_id"`
	Text       string    `json:"text" gorm:"type:text;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
}

type QuestionWithAnswers struct {
	Question Question `json:"question"`
	Answers  []Answer `json:"answers"`
}
