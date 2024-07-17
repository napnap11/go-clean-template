package repository

import (
	"time"

	"github.com/napnap11/go-clean-template/internal/pkg/model"
)

type ExampleRepo struct {
	// Add any dependencies here if needed, such as a database connection
}

func NewExampleRepo() *ExampleRepo {
	return &ExampleRepo{}
}

func (r *ExampleRepo) GetExampleData() (*model.Example, error) {
	// Create a mock Example instance
	mockExample := &model.Example{
		ID:        1,
		Name:      "Mock Example",
		CreatedAt: time.Now().Add(-24 * time.Hour), // 1 day ago
		UpdatedAt: time.Now(),
	}

	return mockExample, nil
}
