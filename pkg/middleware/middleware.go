package midl

import "net/http"

// MiddlewareFunc shortcut for handler func
type MiddlewareFunc func(w http.ResponseWriter, r *http.Request)

// RootgMiddleware ignore sub pathes
func RootMiddleware(f MiddlewareFunc) MiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		f(w, r)
	}
}

// JSONHeaderMiddleware sets response as JSON
func JSONHeaderMiddleware(f MiddlewareFunc) MiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header()["Content-Type"] = []string{"application/json"}
		f(w, r)
	}
}

func WrapMiddleware(f MiddlewareFunc) MiddlewareFunc {
	return JSONHeaderMiddleware(RootMiddleware(f))
}
