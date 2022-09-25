package v1

import (
	"net/http"

	"github.com/MorselShogiew/Users-service-billing/logger"
	"github.com/MorselShogiew/Users-service-billing/service/usecases"
)

type Handlers struct {
	u *usecases.ResizeService
	l logger.Logger
}

func New(u *usecases.ResizeService, l logger.Logger) *Handlers {
	return &Handlers{u, l}
}

func (h *Handlers) CheckErrWriteResp(e error, w http.ResponseWriter, requestID string) {
	if e == nil {
		w.WriteHeader(200)
		return
	}

	if err, ok := e.(interface{ StatusCode() int }); ok {
		h.l.Error(requestID, e)
		w.WriteHeader(err.StatusCode())
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(e.Error()))
		return
	}

	h.l.Error(requestID, e)
	w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(e.Error()))
}
