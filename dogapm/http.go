package dogapm

import (
	"context"
	"net/http"
)

type HttpServer struct {
	mux *http.ServeMux
	*http.Server
}

func NewHttpServer(addr string) *HttpServer {
	mux := http.NewServeMux()
	server := &http.Server{Addr: addr, Handler: mux}
	s := &HttpServer{mux: mux, Server: server}
	globalClosers = append(globalClosers, s)
	globalStarters = append(globalStarters, s)
	return s
}

func (h *HttpServer) Handle(pattern string, handler http.Handler) {
	h.mux.Handle(pattern, handler)
}

func (h *HttpServer) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	h.mux.HandleFunc(pattern, handler)
}

func (h *HttpServer) Start() {
	go func() {
		err := h.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
}

func (h *HttpServer) Close() {
	h.Shutdown(context.TODO())
}
