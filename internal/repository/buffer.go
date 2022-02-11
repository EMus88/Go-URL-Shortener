package repository

import (
	"sync"

	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository/model"
)

type Buffer struct {
	Buffer []model.URL
	Mutex  sync.Mutex
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: make([]model.URL, 0, 5)}
}
