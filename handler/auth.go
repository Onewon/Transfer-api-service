package handler

import (
	"net/http"
)

// Auth Interceptor handler
func AuthInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var user_id string
			var auth_code string
			if r.Method == http.MethodGet {
				vars := r.URL.Query()
				user_id = vars.Get("uid")
				auth_code = vars.Get("auth")
			} else {
				r.ParseForm()
				user_id = r.Form.Get("uid")
				auth_code = r.Form.Get("auth")
			}
			if len(user_id) < 3 || !IsAuthValid(user_id, auth_code) {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			h(w, r)
		})
}
