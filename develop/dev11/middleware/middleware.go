package middleware

import (
	"log"
	"net/http"
)

//каждый запрос проходит через данный обработчик и выводится его метод, URI и IP
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, req.RemoteAddr)
	})
}
