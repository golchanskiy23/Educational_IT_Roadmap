package main

import(
	"net/http"
	"sync"
)

type Router struct{
	routes map[string]map[string]http.HandlerFunc
    mu     sync.RWMutex
}

func (r *Router) Handle(method, path string, h http.HandlerFunc){
	r.mu.Lock()
    defer r.mu.Unlock()

    if r.routes[method] == nil {
        r.routes[method] = make(map[string]http.HandlerFunc)
    }
    r.routes[method][path] = h
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request){
	router.mu.RLock()
    defer router.mu.RUnlock()

    methods, ok := router.routes[r.Method]
    if !ok {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    h, ok := methods[r.URL.Path]
    if !ok {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }

    h(w, r)
}