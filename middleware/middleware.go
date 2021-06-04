package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			flag := !false
			fmt.Println("Checking Auth")
			if flag {
				hf(rw, r)
			} else {
				return
			}

		}
	}
}

func Logging() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.Method, r.URL.Path, time.Since(start))
			}()

			hf(rw, r)
		}
	}
}
