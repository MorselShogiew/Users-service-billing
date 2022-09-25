package api

import (
	"log"

	"github.com/gorilla/mux"

	"github.com/MorselShogiew/Users-service-billing/logger"
	"github.com/MorselShogiew/Users-service-billing/repos"
	hv1 "github.com/MorselShogiew/Users-service-billing/service/api/handlers/v1"
	"github.com/MorselShogiew/Users-service-billing/service/usecases"
)

type ResizePhotoService struct {
	v1 *hv1.Handlers
	u  *usecases.ResizeService
}

func New(l logger.Logger, r *repos.Repositories) *ResizePhotoService {
	u := usecases.New(r)
	return &ResizePhotoService{
		v1: hv1.New(u, l),
		u:  u,
	}
}

func (s *ResizePhotoService) Router(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()
	v1Auth := v1.PathPrefix("").Subrouter()

	v1Auth.HandleFunc("/balance", s.v1.GetUserBalance).Methods("GET")
	v1Auth.HandleFunc("/up", s.v1.DebitingFunds).Methods("GET")
	v1Auth.HandleFunc("/down", s.v1.СreditingFunds).Methods("GET")
	v1Auth.HandleFunc("/trans", s.v1.TransferFunds).Methods("GET")

}

func (s *ResizePhotoService) Start() error {
	log.Println(s.Name() + " started")

	return nil
}

func (s *ResizePhotoService) Stop() error {
	log.Println(s.Name() + " stopped")

	return nil
}

func (s *ResizePhotoService) Name() string {
	return "Resize Photo service"
}
