package server

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/fvbock/endless"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mgjules/mgjules-go/auth"
	"github.com/mgjules/mgjules-go/logger"
	"github.com/mgjules/mgjules-go/projection"
)

type Server struct {
	engine     *gin.Engine
	host       string
	port       int
	auth       *auth.Auth
	projection *projection.Projection
	static     fs.FS
}

func New(prod bool, host string, port int, auth *auth.Auth, projection *projection.Projection, static embed.FS) *Server {
	if prod {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	desugared := logger.L.Desugar()

	engine.Use(ginzap.Ginzap(desugared, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(desugared, true))

	s := &Server{
		engine:     engine,
		host:       host,
		port:       port,
		auth:       auth,
		projection: projection,
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
	es.Server.ReadTimeout = 10 * time.Second
	es.Server.WriteTimeout = 10 * time.Second
	es.Server.MaxHeaderBytes = 1 << 20

	return es.ListenAndServe()
}

func (s *Server) AttachRoutes() {
	s.engine.GET("/", s.IndexHandler())
	s.engine.GET("/cv/:section", s.CVHandler())

	s.engine.StaticFileFS("/favicon.ico", "static/favicon.ico", http.FS(s.static))

	css, err := fs.Sub(s.static, "static/css")
	if err != nil {
		logger.L.Errorf("error when creating css FS handler: %v", err)
	} else {
		s.engine.StaticFS("/css", http.FS(css))
	}

	img, err := fs.Sub(s.static, "static/img")
	if err != nil {
		logger.L.Errorf("error when creating image FS handler: %v", err)
	} else {
		s.engine.StaticFS("/img", http.FS(img))
	}

	authenticated := s.engine.Group("/_")
	authenticated.Use(s.Authorize())
	{
		authenticated.POST("/refetch-data", s.RefetchDataHandler())
		authenticated.POST("/rebuild-projections", s.RebuildProjectionsHandler())
	}
}
