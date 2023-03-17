package jsonrpc

import (
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
)

type Config struct {
	Addr     string
	Endpoint string
}

type Server struct {
	rpc  *rpc.Server
	conf Config
}

func New(conf Config) *Server {
	return &Server{
		rpc:  rpc.NewServer(),
		conf: conf,
	}
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if contentType := r.Header.Get("Accept"); !strings.Contains(contentType, "application/json") {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	var (
		conn  = httpConn{r, w}
		codec = jsonrpc.NewServerCodec(&conn)
	)

	conn.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	conn.w.WriteHeader(http.StatusOK)
	srv.rpc.ServeCodec(codec)
}

func (srv *Server) Register(receivers ...any) *Server {
	for _, receiver := range receivers {
		_ = srv.rpc.Register(receiver)
	}
	return srv
}

func (srv *Server) RegisterName(name string, receiver any) *Server {
	_ = srv.rpc.RegisterName(name, receiver)
	return srv
}

func (srv *Server) ListenAndServe() error {
	http.Handle(srv.conf.Endpoint, srv)
	srv.rpc.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath) // used only for debug
	log.Print("Starting JSON-RPC 1.0 server on addr: ", srv.conf.Addr)
	return http.ListenAndServe(srv.conf.Addr, nil)
}

type httpConn struct {
	r *http.Request
	w http.ResponseWriter
}

func (c *httpConn) Read(p []byte) (n int, err error) {
	return c.r.Body.Read(p)
}

func (c *httpConn) Write(p []byte) (n int, err error) {
	return c.w.Write(p)
}

func (c *httpConn) Close() error {
	return c.r.Body.Close()
}
