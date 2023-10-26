package rest

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func ErrorHandle(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			if errors.Is(ErrBadRequest, err) {
				w.WriteHeader(http.StatusBadRequest)
			}
			if errors.Is(err, ErrInternal) {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte(err.Error()))

		}
	}
}
