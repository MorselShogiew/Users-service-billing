package usecases

import (
	"github.com/MorselShogiew/Users-service-billing/repos"
)

type ResizeService struct {
	resizeAPIRepo repos.ResizeDBRepo
}

func New(r *repos.Repositories) *ResizeService {
	return &ResizeService{
		r.ResizeDBRepo,
	}
}
