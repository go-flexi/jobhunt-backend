package web

type Middleware func(HandlerFunc) HandlerFunc

func WrapMiddlewares(h HandlerFunc, middlewares ...Middleware) HandlerFunc {
	for _, m := range middlewares {
		h = m(h)
	}

	return h
}
