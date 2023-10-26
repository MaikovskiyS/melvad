package service

import (
	"context"
	"melvad/internal/model"
)

type Incrementer interface {
	Increment(ctx context.Context, d model.IncrementData) (uint8, error)
}
type Storage interface {
	Save(ctx context.Context, u model.User) (uint64, error)
}
type Hasher interface {
	Encrypt(ctx context.Context, h model.HashData) (string, error)
}
type service struct {
	hasher Hasher
	i      Incrementer
	user   Storage
}

func New(i Incrementer, s Storage, h Hasher) *service {
	return &service{
		hasher: h,
		i:      i,
		user:   s,
	}
}
func (s *service) Increment(ctx context.Context, d model.IncrementData) (uint8, error) {
	return s.i.Increment(ctx, d)
}
func (s *service) Save(ctx context.Context, u model.User) (uint64, error) {
	return s.user.Save(ctx, u)
}
func (s *service) Encrypt(ctx context.Context, h model.HashData) (string, error) {
	return s.hasher.Encrypt(ctx, h)
}
