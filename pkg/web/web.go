package web

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

// JSONEncoder is an interface that wraps the JSONEncode method to get the JSON encoded bytes
type JSONEncoder interface {
	// JSONEncode returns the JSON encoded bytes
	JSONEncode() ([]byte, error)
}

// HandlerFunc is a function that takes a context and a http.Request and returns a JSONEncoder and an error
type HandlerFunc func(ctx context.Context, r *http.Request) (JSONEncoder, error)

type App struct {
	router      *chi.Mux
	middleWares []Middleware
}

func NewApp(middlewares []Middleware) *App {
	r := chi.NewRouter()

	return &App{
		router:      r,
		middleWares: middlewares,
	}
}

func (a *App) AddHandler(method, path string, handler HandlerFunc, middlewares ...Middleware) {
	handler = WrapMiddlewares(handler, append(a.middleWares, middlewares...)...)

}
