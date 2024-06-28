package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/vaberof/smartway-task/cmd/smartway-task/docs"
	"github.com/vaberof/smartway-task/internal/app/entrypoint/http"
	"github.com/vaberof/smartway-task/internal/domain/employee"
	"github.com/vaberof/smartway-task/internal/infra/storage/postgres/pgcompany"
	"github.com/vaberof/smartway-task/internal/infra/storage/postgres/pgemployee"
	"github.com/vaberof/smartway-task/pkg/database/postgres"
	"github.com/vaberof/smartway-task/pkg/http/httpserver"
	"github.com/vaberof/smartway-task/pkg/logging/logs"
	"os"
)

var appConfigPaths = flag.String("config.files", "not-found.yaml", "List of application config files separated by comma")
var environmentVariablesPath = flag.String("env.vars.file", "not-found.env", "Path to environment variables file")

//	@title			Employee API
//	@version		1.0
//	@description	API Server for Employee Application

// @host		localhost:8000
// @BasePath	/api/v1
func main() {
	flag.Parse()
	if err := loadEnvironmentVariables(); err != nil {
		panic(err)
	}

	appConfig := mustGetAppConfig(*appConfigPaths)

	fmt.Printf("%+v\n", appConfig)

	logger := logs.New(os.Stdout, nil)

	postgresManagedDb, err := postgres.New(&appConfig.Postgres)
	if err != nil {
		panic(err)
	}

	pgEmployeeStorage := pgemployee.NewPgEmployeeStorage(postgresManagedDb.PostgresDb)
	pgCompanyStorage := pgcompany.NewPgCompanyStorage(postgresManagedDb.PostgresDb)

	domainEmployeeService := employee.NewEmployeeService(pgEmployeeStorage, pgCompanyStorage, logger)

	httpHandler := http.NewHandler(domainEmployeeService)

	appServer := httpserver.New(&appConfig.Server, logger)

	httpHandler.InitRoutes(appServer.ChiRouter)

	serverExitChannel := appServer.StartAsync()

	<-serverExitChannel
}

func loadEnvironmentVariables() error {
	return godotenv.Load(*environmentVariablesPath)
}
