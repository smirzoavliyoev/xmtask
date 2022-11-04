package handlers

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/smirzoavliyoev/xmtask/internal/regionservice"
	"github.com/smirzoavliyoev/xmtask/pkg/jwt"
	"github.com/smirzoavliyoev/xmtask/pkg/responser"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, req)
	})
}

// Middleware function, which will be called for each request
func CountryResrictionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		regionService := regionservice.NewRegionService("192.168.0.1")

		country, err := regionService.GetRegionBasedOnIp(r.RemoteAddr)
		if err != nil {
			responser.Response(responser.InternalError, w)
			return
		}
		if r.Method == "POST" || r.Method == "DELETE" {
			if country != "Cyprus" {
				responser.Response(responser.Forbbiden, w)
				return
			}
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)

	})
}

func AuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" && r.Method != "DELETE" {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("auth-token")
		err := jwt.Validate(token, "secret")
		if err != nil {
			responser.Response(responser.Forbbiden, w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
