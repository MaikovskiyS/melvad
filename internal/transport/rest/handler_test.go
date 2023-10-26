package rest

import (
	"bytes"
	"errors"
	"melvad/internal/model"
	mock_rest "melvad/internal/transport/rest/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Increment(t *testing.T) {
	type mockBehavior func(s *mock_rest.MockService, d model.IncrementData)
	tests := []struct {
		name            string
		inputBody       string
		input           *IncrementRequest
		data            model.IncrementData
		mockBehavior    mockBehavior
		expStatusCode   int
		expErr          error
		expResponseBody []byte
	}{
		{
			name:      "ok",
			inputBody: `{"key":"age","value":19}`,
			data:      model.IncrementData{Key: "age", Value: 19},
			mockBehavior: func(s *mock_rest.MockService, d model.IncrementData) {
				s.EXPECT().Increment(gomock.Any(), d).Return(uint8(20), nil)
			},
			expErr:          ErrBadRequest,
			expStatusCode:   http.StatusOK,
			expResponseBody: []byte(`{"value":20}`),
		},
		{
			name:      "internal_err",
			inputBody: `{"key":"age","value":19}`,
			data:      model.IncrementData{Key: "age", Value: 19},
			mockBehavior: func(s *mock_rest.MockService, d model.IncrementData) {
				s.EXPECT().Increment(gomock.Any(), d).Return(uint8(0), errors.New(""))
			},
			expErr:          ErrInternal,
			expStatusCode:   http.StatusInternalServerError,
			expResponseBody: []byte(`server_error`),
		},
	}
	for _, usecase := range tests {
		t.Run(usecase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			svc := mock_rest.NewMockService(c)
			usecase.mockBehavior(svc, usecase.data)
			handler := New(svc)
			mux := http.NewServeMux()
			mux.HandleFunc("/redis/incr", ErrorHandle(handler.Increment))
			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST", "/redis/incr", bytes.NewBufferString(usecase.inputBody))
			mux.ServeHTTP(w, req)
			assert.Equal(t, usecase.expResponseBody, w.Body.Bytes())
			assert.Equal(t, usecase.expStatusCode, w.Code)
		})
	}
}
