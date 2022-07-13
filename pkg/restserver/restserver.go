package restserver

import (
	"context"
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var (
	once   *sync.Once
	server *Server
)

type Server struct {
	serverPort int
	conf       ServerConfig
	restServer *echo.Echo
	DB         *gorm.DB
}

type ServerConfig struct {
	APIVersion string
}

func GetServerInstance(serverPort int, conf *ServerConfig, DB *gorm.DB) *Server {
	server = createServer(serverPort, conf, DB)
	return server
}

func createServer(serverPort int, conf *ServerConfig, DB *gorm.DB) *Server {
	apiServer := echo.New()
	apiServer.Use(middleware.Recover())
	apiServer.Pre(middleware.RemoveTrailingSlash())
	serverConf := ServerConfig{
		APIVersion: conf.APIVersion,
	}

	s := &Server{
		serverPort: serverPort,
		conf:       serverConf,
		restServer: apiServer,
		DB:         DB,
	}

	return s
}

func (s *Server) StartServer() error {
	s.loadAPIGroups()

	go func(port int) {
		addr := fmt.Sprintf(":%d", port)
		log.Fatal(s.restServer.Start(addr))
	}(s.serverPort)

	return nil
}

func (s *Server) Shutdown() {
	s.restServer.Shutdown(context.Background())
}
