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
	public     fs.FS
}

func New(prod bool, host string, port int, auth *auth.Auth, projection *projection.Projection, public embed.FS) *Server {
	if prod {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	desugared := logger.Logger.Desugar()

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
		s.public = public
	} else {
		s.public = os.DirFS(".")
	}

	s.AttachRoutes()

	return s
}

func (s *Server) Start() error {
	logger.Logger.Infof("server listening on %s:%d...", s.host, s.port)
	return endless.ListenAndServe(fmt.Sprintf("%s:%v", s.host, s.port), s.engine)
}

func (s *Server) AttachRoutes() {
	s.engine.GET("/", s.IndexHandler())

	s.engine.StaticFileFS("/favicon.ico", "public/favicon.ico", http.FS(s.public))

	assets, err := fs.Sub(s.public, "public/assets")
	if err != nil {
		logger.Logger.Errorf("error when creating assets FS handler: %v", err)
	} else {
		s.engine.StaticFS("/assets", http.FS(assets))
	}

	img, err := fs.Sub(s.public, "public/img")
	if err != nil {
		logger.Logger.Errorf("error when creating image FS handler: %v", err)
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
