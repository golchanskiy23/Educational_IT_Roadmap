package main

import (
	"log"
	"net/http"
	"time"
	"runtime/debug"
)

// type Middleware func(http.Handler) http.Handler

func Auth(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Header.Get("Authorization") == ""{
			http.Error(w, "error in authorization", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w,r)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("[%d] %s %s — %v", http.StatusAccepted, r.Method, r.URL.Path, time.Since(start))
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v\n%s", err, debug.Stack())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func chain(handler http.Handler, middlewares ...func(http.Handler) http.Handler)  http.Handler{
	for _, m := range middlewares{
		handler = m(handler)
	}
	return handler
}

func listUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{["alice" : 42, "bob" : 32]}`))
}

func Param(r *http.Request, key string) string {
    params, _ := r.Context().Value(paramsKey).(map[string]string)
    return params[key]
}

func getUser(w http.ResponseWriter, r *http.Request) {
    id := Param(r, "id")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"id":"` + id + `"}`))
}