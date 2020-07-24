package routes

import "github.com/go-zs/go-demo/di/digdemo/di/db"

type (
	Auth struct {
		Models *db.ModelManager
	}

	MainController struct {
		Models *db.ModelManager
	}
)
