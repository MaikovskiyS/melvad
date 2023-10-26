package hasher

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"melvad/internal/model"
)

type hasher struct {
}

func New() *hasher {
	return &hasher{}
}
func (h *hasher) Encrypt(_ context.Context, d model.HashData) (string, error) {

	hash := hmac.New(sha512.New, []byte(d.Key))
	hash.Write([]byte(d.Text))
	hashStr := hex.EncodeToString(hash.Sum(nil))
	return hashStr, nil

}
