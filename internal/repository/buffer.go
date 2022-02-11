package repository

import (
	"sync"

	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository/model"
)

type DeleteBuffer struct {
	Buffer []model.URL
	Mutex  sync.Mutex
	Full   chan struct{}
}

func NewDeleteBuffer() *DeleteBuffer {
	return &DeleteBuffer{Buffer: make([]model.URL, 0, 5)}
}
