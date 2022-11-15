package main

import (
	"crypto/sha1"

	"github.com/ellipizle/crime-report/config"
	"github.com/ellipizle/crime-report/pkg/middleware/jwt"
	"github.com/ellipizle/crime-report/pkg/postgresql"
	"github.com/ellipizle/crime-report/pkg/rbac"
	"github.com/ellipizle/crime-report/pkg/secure"
	"github.com/ellipizle/crime-report/pkg/zlog"

	crime_case "github.com/ellipizle/crime-report/internal/crime-case"
	ccl "github.com/ellipizle/crime-report/internal/crime-case/logging"
	cct "github.com/ellipizle/crime-report/internal/crime-case/transport"

	"github.com/ellipizle/crime-report/internal/auth"
	authl "github.com/ellipizle/crime-report/internal/auth/logging"
	autht "github.com/ellipizle/crime-report/internal/auth/transport"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// "github.com/ory/viper"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/ellipizle/crime-report/docs"
)

// var configuration c.Configurations

// @title Crime Report API
// @version 1.0
// @description Crime report server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.crimereporter.io/support
// @contact.email support@crimereporter.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.crimereporter.io
// @BasePath /v2
func main() {
	config.Setup()
	db := postgresql.New(config.App.PostgresHost, config.App.PostgresDb, config.App.PostgresUser, config.App.PostgresPassword, config.App.PostgresPort)
	rbac := rbac.New()
	sec := secure.New(sha1.New())
	log := zlog.New()
	jwt := jwt.New(config.App.JWTSecret, config.App.JWTSigningAlgorithm, config.App.JWTDuration)

	// initialize new echo
	e := echo.New()
	e.Use(middleware.Logger())  // Logger
	e.Use(middleware.Recover()) // Recover
	// at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac, sms, smpt), log), e, jwt.MWFunc())
	autht.NewHTTP(authl.New(auth.Initialize(db, jwt, sec, rbac), log), e, jwt.MWFunc())
	// create a group v1
	v1 := e.Group("/v1")
	v1.Use(jwt.MWFunc())
	v1.Use(jwt.MWFunc())
	cct.NewHTTP(ccl.New(crime_case.Initialize(db, rbac), log), v1)

	//start server
	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
