package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/gommon/log"
	"github.com/lazyboson/lostapi/pkg/restserver"
	"github.com/lazyboson/lostapi/pkg/storage"
	"github.com/subosito/gotenv"
)

const (
	envDatabaseHost     = "DATABASE_HOST"
	envDatabasePort     = "DATABASE_PORT"
	envDatabaseUserName = "DATABASE_USERNAME"
	envDatabasePassword = "DATABASE_PASSWORD"
	envDatabaseName     = "DATABASE_NAME"
	envSSLMode          = "SSL_MODE"
	envAPIVersion       = "API_VERSION"
)

type restService struct {
	restServer     *restserver.Server
	restConfig     *restserver.ServerConfig
	dataBaseConfig *storage.DatabaseConfig
}

func new() (s *restService) {
	s = &restService{}
	return
}

func (s *restService) configure() {
	gotenv.Load("rest.env")

	s.restConfig = loadRestConfig()
	s.dataBaseConfig = loadDataBaseConfig()
}

func (s *restService) start() {
	DB, err := storage.NewConnection(s.dataBaseConfig)
	if err != nil {
		log.Fatal("failed to make connection with DB")
		os.Exit(1)
	}

	s.restServer = restserver.GetServerInstance(3232, s.restConfig, DB)
	err = s.restServer.MigrateItems()
	if err != nil {
		log.Fatal("failed to run")
		os.Exit(1)
	}

	s.restServer.StartServer()
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	restService := new()
	restService.configure()
	restService.start()
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	<-done
}

func loadDataBaseConfig() *storage.DatabaseConfig {
	dbConf := &storage.DatabaseConfig{
		Host:         "localhost",
		DatabasePort: "5432",
		Password:     "password",
		UserName:     "postgres",
		DatabaseName: "inventory",
		SSLMode:      "disable",
	}

	confHost := os.Getenv(envDatabaseHost)
	if confHost != "" {
		dbConf.Host = confHost
	}

	confDBPort := os.Getenv(envDatabasePort)
	if confDBPort != "" {
		dbConf.DatabasePort = confDBPort
	}

	confDBUserName := os.Getenv(envDatabaseUserName)
	if confDBUserName != "" {
		dbConf.UserName = confDBUserName
	}

	confDBPassword := os.Getenv(envDatabasePassword)
	if confDBPassword != "" {
		dbConf.Password = confDBPassword
	}

	confDBName := os.Getenv(envDatabaseName)
	if confDBName != "" {
		dbConf.DatabaseName = confDBName
	}

	return dbConf
}

func loadRestConfig() *restserver.ServerConfig {
	restConf := &restserver.ServerConfig{APIVersion: "v1"}

	confApiVersion := os.Getenv(envAPIVersion)
	if confApiVersion != "" {
		restConf.APIVersion = confApiVersion
	}

	return restConf
}
