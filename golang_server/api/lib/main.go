package Http

import (
	"encoding/json"
	"fmt"
	api_middleware "idk/main/api/middleware"
	"net/http"
	"net/url"
)

type BetterResponseWriter struct {
	w http.ResponseWriter
}

func (brw *BetterResponseWriter) AddHeader(key string, value string) *BetterResponseWriter {
	brw.w.Header().Add(key, value)
	return brw
}
func (brw *BetterResponseWriter) DeleteHeader(key string) *BetterResponseWriter {
	brw.w.Header().Del(key)
	return brw
}
func (brw *BetterResponseWriter) Status(code int) *BetterResponseWriter {
	brw.w.WriteHeader(code)
	return brw
}
func (brw *BetterResponseWriter) Send(message interface{}) {
	data, ok := message.([]byte)
	if !ok {
		data_str, ok := message.(string)
		if ok {
			data = []byte(data_str)
		} else {
			var err error
			data, err = json.Marshal(message)
			if err != nil {
				data = []byte(fmt.Sprint(message))
			}
		}
	}
	brw.w.Write(data)
}
func NewBetterResponseWriter(w http.ResponseWriter) *BetterResponseWriter {
	return &BetterResponseWriter{w}
}

type cluster struct {
	Mux  *http.ServeMux
	Path string
}

func Cluster(mux *http.ServeMux, path string) *cluster {
	return &cluster{mux, path}
}

func (cluster *cluster) Rest(path string) *endpoint {
	end_path, _ := url.JoinPath(cluster.Path, path)
	return Rest(cluster.Mux, end_path)
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
		if r.Method == "OPTIONS" {
			api_middleware.AllowCors(w, r)
			return
		}
		var handler func(w http.ResponseWriter, r *http.Request)
		if h, ok := endpoint.Handlers[r.Method]; ok {
			handler = h
		} else {
			w.WriteHeader(404)
			return
		}
		for _, middleware := range endpoint.Middlewares {
			err := middleware(w, r)
			if err != nil {
				return
			}
		}
		handler(w, r)
	})
}
