package repos

import (
	"github.com/MorselShogiew/Users-service-billing/logger"
	"github.com/MorselShogiew/Users-service-billing/provider"
)

type Repositories struct {
	ResizeDBRepo ResizeDBRepo
}

func New(p provider.Provider, l logger.Logger) *Repositories {
	ResizeDBRepo := NewResizeDBRepo(p, l)
	return &Repositories{
		ResizeDBRepo,
	}
}
