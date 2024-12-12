package http

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mgjules/mgjules-go/auth"
	"github.com/mgjules/mgjules-go/fetcher"
	"github.com/mgjules/mgjules-go/logger"
	"github.com/mgjules/mgjules-go/projecter"
	"golang.org/x/crypto/acme/autocert"
)

type Server struct {
	engine    *gin.Engine
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
	if prod {
		gin.SetMode(gin.ReleaseMode)
	}
	if auth == nil {
		panic("auth cannot be nil")
	}
	if fetcher == nil {
		panic("fetcher cannot be nil")
	}
	if projection == nil {
		panic("projection cannot be nil")
	}

	engine := gin.New()

	desugared := logger.L.Desugar()

	engine.Use(ginzap.Ginzap(desugared, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(desugared, true))

	s := &Server{
		engine:    engine,
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
		s.static = os.DirFS(".")
	}

	s.AttachRoutes()

	return s
}

func (s *Server) Start() error {
	logger.L.Infof("server listening on %s:%d...", s.host, s.port)

	es := endless.NewServer(fmt.Sprintf("%s:%v", s.host, s.port), s.engine)
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

func (s *Server) AttachRoutes() {
	if !s.prod {
		pprof.Register(s.engine)
	}

	s.engine.GET("/", s.IndexHandler())
	{
		blog := s.engine.Group("/blog")
		blog.GET("/", s.BlogIndexHandler())
		blog.GET("/:slug", s.BlogPostHandler())
	}

	s.engine.NoRoute(s.NotFoundHandler())

	s.engine.StaticFileFS("/favicon.ico", "static/favicon.ico", http.FS(s.static))

	css, err := fs.Sub(s.static, "static/css")
	if err != nil {
		logger.L.Errorf("error when creating css FS handler: %v", err)
	} else {
		cssR := s.engine.Group("/css")
		cssR.Use(gzip.Gzip(gzip.BestCompression))
		cssR.StaticFS("/", http.FS(css))
	}

	img, err := fs.Sub(s.static, "static/img")
	if err != nil {
		logger.L.Errorf("error when creating image FS handler: %v", err)
	} else {
		imgR := s.engine.Group("/img")
		imgR.Use(gzip.Gzip(gzip.BestCompression))
		imgR.StaticFS("/", http.FS(img))
	}

	fonts, err := fs.Sub(s.static, "static/fonts")
	if err != nil {
		logger.L.Errorf("error when creating image FS handler: %v", err)
	} else {
		fontR := s.engine.Group("/fonts")
		fontR.Use(gzip.Gzip(gzip.BestCompression))
		fontR.StaticFS("/", http.FS(fonts))
	}

	authenticated := s.engine.Group("/_")
	authenticated.Use(s.Authorize())
	{
		authenticated.POST("/refetch-data", s.RefetchDataHandler())
		authenticated.POST("/rebuild-projections", s.RebuildProjectionsHandler())
	}
}
