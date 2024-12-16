package http

import (
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/mgjules/mgjules-go/internal/auth"
	"github.com/mgjules/mgjules-go/internal/fetcher"
	"github.com/mgjules/mgjules-go/internal/projecter"
	"github.com/mgjules/mgjules-go/internal/static"
	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	mux       *http.ServeMux
	prod      bool
	host      string
	port      int
	tlsDomain string
	auth      *auth.Auth
	fetcher   *fetcher.Fetcher
	projecter *projecter.Projecter
	static    fs.FS
}

func NewServer(
	prod bool,
	host string,
	port int,
	tlsDomain string,
	auth *auth.Auth,
	fetcher *fetcher.Fetcher,
	projection *projecter.Projecter,
) *Server {
	if auth == nil {
		panic("auth cannot be nil")
	}
	if fetcher == nil {
		panic("fetcher cannot be nil")
	}
	if projection == nil {
		panic("projection cannot be nil")
	}

	s := &Server{
		mux:       http.NewServeMux(),
		prod:      prod,
		host:      host,
		port:      port,
		tlsDomain: tlsDomain,
		auth:      auth,
		fetcher:   fetcher,
		projecter: projection,
		static:    static.FS,
	}

	s.initRoutes()

	return s
}

func (s *Server) Start() error {
	hostport := fmt.Sprintf("%s:%v", s.host, s.port)
	slog.Info("server listening on", "host:port", hostport)

	es := endless.NewServer(hostport, s.mux)
	es.ReadTimeout = 10 * time.Second
	if s.prod {
		es.WriteTimeout = 10 * time.Second
	} else {
		es.WriteTimeout = 60 * time.Second
	}
	es.MaxHeaderBytes = 1 << 20

	if s.tlsDomain != "" {
		es.EndlessListener = autocert.NewListener(s.tlsDomain)
		return es.Serve()
	}

	return es.ListenAndServe()
}

func (s *Server) initRoutes() {
	s.mux.HandleFunc("GET /", s.rootHandler())
	// s.mux.HandleFunc("GET /blog", s.BlogIndexHandler())
	// s.mux.HandleFunc("GET /blog/{slug}", s.BlogPostHandler())

	// s.mux.Handle("/", http.StripPrefix("/", http.FileServerFS(s.static)))
}
