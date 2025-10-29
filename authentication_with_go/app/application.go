package app

import (
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

func NewConfig() Config {
	port := config.GetString("port", "3001")
	return Config{
		Addr: fmt.Sprintf(":%s", port),
	}
}

type Application struct {
	Config Config
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
	}
}

func (app *Application) Run() error {

	db, err := dbConfig.SetupDb()
	if err != nil {
		return err
	}

	ur := repo.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	rr := repo.NewRoleRepository(db)
	rpr := repo.NewRolePermissionRepository(db)
	urr := repo.NewUserRoleRepository(db)
	rs := services.NewRoleService(rr, rpr, urr)
	rc := controllers.NewRoleController(rs)
	roleRouter := router.NewRoleRouter(rc)

	server := &http.Server{ // the & will return the reference of the http server object
		Addr:    app.Config.Addr,
		Handler: router.SetupRouter(uRouter, roleRouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting the server on ", app.Config.Addr);

	return server.ListenAndServe()
}