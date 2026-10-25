package core_http_middleware

import "slices"

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	if len(middlewares) == 0 {
		return h
	}

	for _, middleware := range slices.Backward(middlewares) {
		h = middleware(h)
	}

	return h
}
