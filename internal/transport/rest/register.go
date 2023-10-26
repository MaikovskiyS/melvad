package rest

import "net/http"

func (a *api) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("/redis/incr", ErrorHandle(a.Increment))
	r.HandleFunc("/postgres/users", ErrorHandle(a.Save))
	r.HandleFunc("/sign/hmacsha512", ErrorHandle(a.Sign))
}
