package main

import (
	"net/http"
	"sync"
    "context"
    "strings"
)

type route struct{
    method string
    parts []string
    handler http.HandlerFunc
}

type Router struct{
	routes []route
    mu     sync.RWMutex
    middlewares []func(http.Handler) http.Handler
}

type contextKey string
const paramsKey contextKey = "params"


func splitPath(path string) []string {
    return strings.Split(strings.Trim(path, "/"), "/")
}

func matchRoute(routeParts, reqParts []string) map[string]string {
    if len(routeParts) != len(reqParts) {
        return nil
    }
    params := make(map[string]string)
    for i, part := range routeParts {
        switch {
        case strings.HasPrefix(part, ":"):
            params[part[1:]] = reqParts[i]
        case part != reqParts[i]:
            return nil
        }
    }
    return params
}

func (r *Router) Handle(method, path string, h http.HandlerFunc) {
    r.mu.Lock()
    defer r.mu.Unlock()

    r.routes = append(r.routes, route{
        method:  method,
        parts:   splitPath(path),
        handler: h,
    })
}

func (r *Router) Use(m func(http.Handler)http.Handler){
    r.middlewares = append(r.middlewares, m)
}

func (r *Router) wrap(h http.Handler) http.Handler{
    curr := h
    for _, m := range r.middlewares{
        curr = m(curr)
    }

    return curr
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request){
	router.mu.RLock()
    defer router.mu.RUnlock()

    reqParts := splitPath(r.URL.Path)
    for _, route := range router.routes{
        if route.method != r.Method{
            continue
        }

        params := matchRoute(route.parts, reqParts)
        if params == nil{
            continue
        }
        ctx := context.WithValue(r.Context(), paramsKey, params)
        router.wrap(route.handler).ServeHTTP(w, r.WithContext(ctx))
        return
    }
    http.Error(w, "Nor Found", http.StatusNotFound)
}