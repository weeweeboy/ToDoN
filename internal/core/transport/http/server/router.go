package core_http_server

import (
	core_http_middleware "ToDoN/internal/core/transport/http/middleware"
	"fmt"
	"net/http"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	middleware []core_http_middleware.Middleware
}

func NewAPIVersionRouter(apiVersion ApiVersion, middleware ...core_http_middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   &http.ServeMux{},
		apiVersion: apiVersion,
		middleware: middleware,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		h := core_http_middleware.ChainMiddlewares(route.Handler, r.middleware...)
		r.Handle(pattern, h)
	}
}

func (r *APIVersionRouter) WithMiddleware(middleware ...core_http_middleware.Middleware) http.Handler {
	return core_http_middleware.ChainMiddlewares(r, r.middleware...)
}
