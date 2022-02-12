package repository

import (
	"sync"
	"time"

	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository/model"
)

type DeleteBuffer struct {
	Buffer     []model.URL
	Mutex      sync.Mutex
	Full       chan struct{}
	LastUpdate time.Duration
}

func NewDeleteBuffer() *DeleteBuffer {
	return &DeleteBuffer{Buffer: make([]model.URL, 0, 10), Full: make(chan struct{})}
}
func (buf *DeleteBuffer) ClearBuffer() {
	buf.Buffer = buf.Buffer[:0]
}
