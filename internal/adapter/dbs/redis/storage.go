package rstorage

import (
	"context"
	"melvad/internal/model"

	"github.com/redis/go-redis/v9"
)

type store struct {
	cl *redis.Client
}

func New(cl *redis.Client) *store {
	return &store{
		cl: cl,
	}
}
func (s *store) Increment(ctx context.Context, d model.IncrementData) (uint8, error) {
	result := s.cl.IncrBy(ctx, d.Key, int64(d.Value))
	value, err := result.Uint64()
	if err != nil {
		return 0, err
	}
	return uint8(value), nil
}
