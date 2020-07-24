package di

import (
	"fmt"
	"github.com/go-zs/go-demo/di/digdemo/di/db"
	"github.com/go-zs/go-demo/di/digdemo/di/graph"
	"github.com/go-zs/go-demo/di/digdemo/di/routes"
	"go.uber.org/dig"
	"gopkg.in/mgo.v2"
)

type (
	Config struct {
		DatabaseUri string
		Database    string
	}

	Server struct {
		config         *Config
		authController *routes.Auth
		mainController *routes.MainController
		GraphqlServer  *graph.GraphqlServer
	}
)

func (s Server) Run() {
	fmt.Println("server start")
}

var Container *dig.Container

func NewConfig() *Config {
	return &Config{DatabaseUri: "mongodb://rwuser:YWprX29wcwo=@172.16.112.191:8635"}
}

func ConnectDatabase(config *Config) *mgo.Session {
	session, err := mgo.Dial(config.DatabaseUri)
	if err != nil {
		panic("DB connection error")
	}
	return session
}

func NewServer(config *Config, authC *routes.Auth, mainC *routes.MainController, graphqlServer *graph.GraphqlServer) *Server {
	return &Server{
		config:         config,
		authController: authC,
		mainController: mainC,
		GraphqlServer:  graphqlServer,
	}
}

func AuthController(models *db.ModelManager) *routes.Auth {
	return &routes.Auth{
		Models: models,
	}
}

func MainController(models *db.ModelManager) *routes.MainController {
	return &routes.MainController{
		Models: models,
	}
}

func ModelManager(config *Config, session *mgo.Session) *db.ModelManager {
	return &db.ModelManager{Session: session, DataBase: config.Database}
}

func BuildContainer() *dig.Container {
	Container := dig.New()
	Container.Provide(ModelManager)
	Container.Provide(NewConfig)
	Container.Provide(ConnectDatabase)
	Container.Provide(AuthController)
	Container.Provide(MainController)
	Container.Provide(graph.NewGraphqlServer)
	Container.Provide(NewServer)
	return Container
}
