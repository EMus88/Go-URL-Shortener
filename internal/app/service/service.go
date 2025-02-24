package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/EMus88/go-musthave-shortener-tpl/configs"
	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository"
	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository/model"
	"github.com/EMus88/go-musthave-shortener-tpl/pkg/idgenerator"
)

type Repository interface {
	SaveURL(shortModel *model.URL, sessionID string) (string, error)
	GetURL(id string) (string, error)
	PingDB() error
	GetCookie(s string) error
	SaveCookie(s string) error
	GetList(key string) ([]model.URL, error)
	SaveBatch(list *[]model.URL, sessionID string) error
	DeleteURLs(s []model.URL)
	AddToBuffer(model model.URL)
}
type Service struct {
	Repository
	Config configs.Config
	Auth   Auth
}

func NewService(repos *repository.Storage, config *configs.Config) *Service {
	return &Service{
		Repository: repos,
		Config:     *config,
	}
}

//save long URL in stotage and return short URL
func (s *Service) SaveURL(longURL string, sessionID string) (string, error) {
	var shortModel model.URL
	shortModel.URLID = idgenerator.CreateID(8)
	shortModel.ShortURL = fmt.Sprint(s.Config.BaseURL, "/", shortModel.URLID)
	shortModel.LongURL = longURL

	key, _ := s.Auth.ReadSessionID(sessionID)

	shortURL, err := s.Repository.SaveURL(&shortModel, key)
	if err != nil {
		return shortURL, err
	}

	//save to file
	var model model.FileModel
	file, err := os.OpenFile(s.Config.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return "", err
	}
	defer file.Close()

	model.ID = key
	model.LongURL = longURL

	data, err := json.MarshalIndent(model, "", " ")
	if err != nil {
		return "", err
	}
	file.Write(data)
	return shortURL, nil
}

//get long URL from stotage by short URL
func (s *Service) GetURL(key string) (string, error) {
	originURL, err := s.Repository.GetURL(key)
	if err != nil {
		return "", err
	}
	return originURL, nil
}

func (s *Service) CreateNewSession() (string, error) {
	id, encID, err := s.Auth.CreateSessionID()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Repository.SaveCookie(id); err != nil {
		return "", err
	}

	return encID, nil
}
func (s *Service) SaveBatch(list []model.BatchRequest, sessionID string) (*[]model.BatchResponse, error) {
	var batch []model.URL
	var response []model.BatchResponse

	for _, val := range list {

		var shortModel model.URL
		var responseModel model.BatchResponse

		shortModel.URLID = idgenerator.CreateID(8)
		shortModel.ShortURL = fmt.Sprint(s.Config.BaseURL, "/", shortModel.URLID)
		shortModel.LongURL = val.OriginalURL

		responseModel.CorrelationID = val.CorrelationID
		responseModel.ShortURL = shortModel.ShortURL

		response = append(response, responseModel)
		batch = append(batch, shortModel)
	}
	key, _ := s.Auth.ReadSessionID(sessionID)
	if err := s.Repository.SaveBatch(&batch, key); err != nil {
		return nil, err
	}

	return &response, nil
}
