package service

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository"
	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository/models/file"
	"github.com/EMus88/go-musthave-shortener-tpl/pkg/idgenerator"
)

type Repository interface {
	SaveURL(key string, value string)
	GetURL(id string) string
}
type Service struct {
	Repository
	Model file.Model
}

func NewService(repos *repository.URLStorage, model *file.Model) *Service {
	return &Service{Repository: repos, Model: *model}
}

func (s *Service) SaveURL(value string) string {
	//save to map
	key := idgenerator.CreateID()
	s.Repository.SaveURL(key, value)
	//save to file
	path := os.Getenv("FILE_STORAGE_PATH")
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	s.Model.ID = key
	s.Model.LongURL = value

	data, err := json.MarshalIndent(s.Model, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	file.Write(data)

	return key
}

func (s *Service) GetURL(key string) string {
	value := s.Repository.GetURL(key)
	return value
}

func (s *Service) LoadFromFile() {
	var model file.Model
	path := os.Getenv("FILE_STORAGE_PATH")
	file, err := os.ReadFile(path)
	if err != nil {
		return
	}
	str := strings.Split(string(file), "}")
	for i := 0; i < (len(str) - 1); i++ {

		if err := json.Unmarshal([]byte(str[i]+"}"), &model); err != nil {
			log.Fatal(err)
		}
		s.Repository.SaveURL(model.ID, model.LongURL)
	}
}
