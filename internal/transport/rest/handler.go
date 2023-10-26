package rest

import (
	"context"
	"encoding/json"
	"errors"
	"melvad/internal/model"
	"net/http"
	"time"
)

var (
	ErrBadRequest = errors.New("bad_request")
	ErrInternal   = errors.New("server_error")
)

//go:generate mockgen -source=handler.go -destination=mocks/mock.go
type Service interface {
	Save(ctx context.Context, u model.User) (uint64, error)
	Increment(ctx context.Context, data model.IncrementData) (uint8, error)
	Encrypt(ctx context.Context, h model.HashData) (string, error)
}
type api struct {
	svc Service
}

var timeout = time.Second * 5

func New(svc Service) *api {
	return &api{svc: svc}
}
func (a *api) Increment(w http.ResponseWriter, r *http.Request) error {
	var data *IncrementRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return ErrBadRequest
	}
	param, err := data.toModel()
	if err != nil {
		return ErrBadRequest
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	value, err := a.svc.Increment(ctx, param)
	if err != nil {
		return ErrInternal
	}
	response := &IncrementResponse{Value: value}

	resBytes, err := json.Marshal(&response)
	if err != nil {
		return ErrInternal
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resBytes)
	return nil
}
func (a *api) Save(w http.ResponseWriter, r *http.Request) error {
	var request *SaveRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return ErrBadRequest
	}
	params, err := request.toModel()
	if err != nil {
		return ErrBadRequest
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	id, err := a.svc.Save(ctx, params)
	if err != nil {
		return ErrInternal
	}
	response := SaveResponse{Id: id}
	resBytes, err := json.Marshal(&response)
	if err != nil {
		return ErrInternal
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(resBytes)
	return nil
}
func (a *api) Sign(w http.ResponseWriter, r *http.Request) error {
	var request *SignRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return ErrBadRequest
	}
	data, err := request.toModel()
	if err != nil {
		return ErrBadRequest
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	hash, err := a.svc.Encrypt(ctx, data)
	if err != nil {
		return ErrInternal
	}
	response := SignResponse{Hash: hash}
	resBytes, err := json.Marshal(&response)
	if err != nil {
		return ErrInternal
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resBytes)
	return nil
}
