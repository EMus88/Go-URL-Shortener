package service

import (
	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository"
	"github.com/EMus88/go-musthave-shortener-tpl/pkg/idgenerator"
)

type Repository interface {
	SaveURL(key string, value string)
	GetURL(id string) string
}
type Service struct {
	Repository
}

func NewService(repos *repository.URLStorage) *Service {
	return &Service{Repository: repos}
}

func (s *Service) SaveURL(value string) string {
	key := idgenerator.CreateID()
	s.Repository.SaveURL(key, value)
	return key
}

func (s *Service) GetURL(key string) string {
	value := s.Repository.GetURL(key)
	return value
}
