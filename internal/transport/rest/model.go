package rest

import (
	"errors"
	"melvad/internal/model"
)

var (
	ErrBadParams = errors.New("bad_params")
)

type IncrementResponse struct {
	Value uint8 `json:"value"`
}
type IncrementRequest struct {
	Key   string `json:"key"`
	Value uint8  `json:"value"`
}

func (r *IncrementRequest) toModel() (model.IncrementData, error) {
	if r == nil {
		return model.IncrementData{}, ErrBadParams
	}
	if r.Key == "" {
		return model.IncrementData{}, ErrBadParams
	}
	if r.Value <= 0 {
		return model.IncrementData{}, ErrBadParams
	}
	data := model.IncrementData{
		Key:   r.Key,
		Value: r.Value,
	}
	return data, nil
}

type SaveResponse struct {
	Id uint64 `json:"id"`
}
type SaveRequest struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func (r *SaveRequest) toModel() (model.User, error) {
	if r == nil {
		return model.User{}, ErrBadParams
	}
	if r.Name == "" {
		return model.User{}, ErrBadParams
	}
	if r.Age <= 0 {
		return model.User{}, ErrBadParams
	}
	data := model.User{
		Name: r.Name,
		Age:  r.Age,
	}
	return data, nil
}

type SignRequest struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}
type SignResponse struct {
	Hash string `json:"hash"`
}

func (r *SignRequest) toModel() (model.HashData, error) {
	if r == nil {
		return model.HashData{}, ErrBadParams
	}
	if r.Key == "" || r.Text == "" {
		return model.HashData{}, ErrBadParams
	}
	data := model.HashData{
		Text: r.Text,
		Key:  r.Key,
	}
	return data, nil
}
