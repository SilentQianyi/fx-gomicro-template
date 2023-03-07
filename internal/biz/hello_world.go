package biz

import (
	"chat/internal/db"
	"chat/internal/domain/bo"
	"chat/internal/domain/do"
	"context"
)

type HelloWorldService interface {
	GetHelloWorldByLanguage(ctx context.Context, language string) ([]*do.HelloWorld, error)
}

type helloWorldServiceImpl struct {
	repo db.HelloWorldRepository
}

func NewHelloWorldService(repo db.HelloWorldRepository) HelloWorldService {
	return &helloWorldServiceImpl{
		repo: repo,
	}
}

func (s *helloWorldServiceImpl) GetHelloWorldByLanguage(ctx context.Context, language string) ([]*do.HelloWorld, error) {
	query := bo.NewHelloWorldQueryByLanguage(language)
	return s.repo.GetHelloWorldByLanguage(ctx, query)
}
