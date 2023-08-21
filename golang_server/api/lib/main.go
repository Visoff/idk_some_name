package Http

import (
	"net/http"
	"net/url"
)

type cluster struct {
	Mux  *http.ServeMux
	Path string
}

func Cluster(mux *http.ServeMux, path string) *cluster {
	return &cluster{mux, path}
}

func (cluster *cluster) Rest(path string) *endpoint {
	end_path, _ := url.JoinPath(cluster.Path, path)
	return &endpoint{cluster.Mux, end_path, []func(w http.ResponseWriter, r *http.Request) error{}, map[string]func(w http.ResponseWriter, r *http.Request){}}
}

func Rest(mux *http.ServeMux, path string) *endpoint {
	return &endpoint{mux, path, []func(w http.ResponseWriter, r *http.Request) error{}, map[string]func(w http.ResponseWriter, r *http.Request){}}
}

type endpoint struct {
	Mux         *http.ServeMux
	Path        string
	Middlewares []func(w http.ResponseWriter, r *http.Request) error
	Handlers    map[string]func(w http.ResponseWriter, r *http.Request)
}

func (endpoint *endpoint) Use(middleware func(w http.ResponseWriter, r *http.Request) error) *endpoint {
	endpoint.Middlewares = append(endpoint.Middlewares, middleware)
	return endpoint
}

func (endpoint *endpoint) Get(handler func(w http.ResponseWriter, r *http.Request)) *endpoint {
	endpoint.Handlers["GET"] = handler
	return endpoint
}
func (endpoint *endpoint) Post(handler func(w http.ResponseWriter, r *http.Request)) *endpoint {
	endpoint.Handlers["POST"] = handler
	return endpoint
}
func (endpoint *endpoint) Patch(handler func(w http.ResponseWriter, r *http.Request)) *endpoint {
	endpoint.Handlers["PATCH"] = handler
	return endpoint
}
func (endpoint *endpoint) Delete(handler func(w http.ResponseWriter, r *http.Request)) *endpoint {
	endpoint.Handlers["DELETE"] = handler
	return endpoint
}

func (endpoint *endpoint) Apply() {
	(*endpoint.Mux).HandleFunc(endpoint.Path, func(w http.ResponseWriter, r *http.Request) {
		if handler, ok := endpoint.Handlers[r.Method]; ok {
			defer handler(w, r)
		} else {
			w.WriteHeader(404)
		}
		for _, middleware := range endpoint.Middlewares {
			err := middleware(w, r)
			if err != nil {
				return
			}
		}
	})
}
