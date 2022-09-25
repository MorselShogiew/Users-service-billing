package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MorselShogiew/Users-service-billing/errs"
	"github.com/MorselShogiew/Users-service-billing/middleware"
	// _ "image/jpeg"
)

func (h *Handlers) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r)

	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

	res, err := h.u.GetUserBalance(reqID, id)
	h.CheckErrWriteResp(err, w, reqID)
	// Encode uses a Writer, use a Buffer if you need the raw []byte

	if err := json.NewEncoder(w).Encode(res); err != nil {
		err := errs.New(nil, errs.ErrJSONEncode, false, 400)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

}
func (h *Handlers) DebitingFunds(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r)

	idStr := r.URL.Query().Get("id")
	valueStr := r.URL.Query().Get("value")
	fmt.Println("asfqef")
	if idStr == "" || valueStr == "" {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

	err = h.u.DebitingFunds(reqID, id, value)
	h.CheckErrWriteResp(err, w, reqID)
	// Encode uses a Writer, use a Buffer if you need the raw []byte

}
func (h *Handlers) СreditingFunds(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r)

	idStr := r.URL.Query().Get("id")
	valueStr := r.URL.Query().Get("value")
	if idStr == "" || valueStr == "" {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

	err = h.u.СreditingFunds(reqID, id, value)
	h.CheckErrWriteResp(err, w, reqID)
	// Encode uses a Writer, use a Buffer if you need the raw []byte

}
func (h *Handlers) TransferFunds(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r)

	idFromStr := r.URL.Query().Get("idFrom")
	idToStr := r.URL.Query().Get("idTo")
	valueStr := r.URL.Query().Get("value")

	if idFromStr == "" || idToStr == "" || valueStr == "" {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}
	idFrom, err := strconv.Atoi(idFromStr)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}
	idTo, err := strconv.Atoi(idToStr)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

	err = h.u.TransferFunds(reqID, idFrom, idTo, value)
	h.CheckErrWriteResp(err, w, reqID)
	// Encode uses a Writer, use a Buffer if you need the raw []byte

}
