package repository

import (
	"sync"

	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository/models/file"
)

type URLStorage struct {
	storage map[string]string
	mx      sync.Mutex
	Model   file.Model
}

func NewStorage() *URLStorage {
	return &URLStorage{
		storage: make(map[string]string, 10),
		mx:      sync.Mutex{},
		Model:   file.Model{},
	}
}
func (us *URLStorage) SaveURL(key string, value string) {
	us.mx.Lock()
	us.storage[key] = value
	us.mx.Unlock()
}

func (us *URLStorage) GetURL(key string) string {
	return us.storage[key]
}
