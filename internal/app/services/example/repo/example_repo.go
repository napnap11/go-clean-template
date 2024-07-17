package repo

import "github.com/napnap11/go-clean-template/internal/pkg/model"

type ExampleRepository interface {
	GetExampleData() (*model.Example, error)
}
