package service

import (
	"fmt"

	"github.com/napnap11/go-clean-template/internal/app/services/example/repo"
)

type ExampleService struct {
	exampleRepo repo.ExampleRepository
}

func NewExampleService(exampleRepo repo.ExampleRepository) *ExampleService {
	return &ExampleService{
		exampleRepo: exampleRepo,
	}
}

func (s *ExampleService) GetExampleData() (string, error) {
	data, err := s.exampleRepo.GetExampleData()
	if err != nil {
		return "", fmt.Errorf("failed to get example data: %w", err)
	}
	return data.Name, nil
}
