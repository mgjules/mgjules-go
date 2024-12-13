package http

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/fvbock/endless"
	"github.com/mgjules/mgjules-go/auth"
	"github.com/mgjules/mgjules-go/fetcher"
	"github.com/mgjules/mgjules-go/logger"
	"github.com/mgjules/mgjules-go/projecter"
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
	static embed.FS,
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
	}

	if prod {
		s.static = static
	} else {
		s.static = os.DirFS("./static")
	}

	s.initRoutes()

	return s
}

func (s *Server) Start() error {
	logger.L.Infof("server listening on %s:%d...", s.host, s.port)

	es := endless.NewServer(fmt.Sprintf("%s:%v", s.host, s.port), s.mux)
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
	s.mux.HandleFunc("GET /{$}", s.IndexHandler())
	s.mux.HandleFunc("/blog", s.BlogIndexHandler())
	s.mux.HandleFunc("/blog/{slug}", s.BlogPostHandler())

	s.mux.Handle("/", http.StripPrefix("/", http.FileServerFS(s.static)))
}
