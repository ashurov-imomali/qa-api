package repository_test

import (
	"github.com/ashurov-imomali/qa-api/internal/db"
	"testing"

	"github.com/ashurov-imomali/qa-api/internal/models"
	"github.com/ashurov-imomali/qa-api/internal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	dns := "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable"
	pgConn, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	err = db.RunMigrations(dns)
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return pgConn
}

func TestQuestionRepository(t *testing.T) {
	pgConn := setupTestDB(t)
	repo := repository.NewRepository(pgConn)
	questionTestId := 0
	t.Run("CreateQuestion", func(t *testing.T) {
		q := &models.Question{
			Text: "Test question",
		}

		err := repo.CreateQuestion(q)
		questionTestId = q.ID
		assert.NoError(t, err)
		assert.NotZero(t, q.ID, "Question ID should be set after creation")
	})

	t.Run("GetQuestionList", func(t *testing.T) {
		questions, err := repo.GetQuestionList()
		assert.NoError(t, err)
		for _, question := range questions {
			if question.ID == questionTestId {
				assert.Equal(t, "Test question", question.Text)
			}
		}
	})

	t.Run("DeleteQuestion", func(t *testing.T) {
		err := repo.DeleteQuestion(questionTestId)
		assert.NoError(t, err)
	})
}
