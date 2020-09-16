package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go_echo_ent/datasource"
	"go_echo_ent/handler"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
)

func main() {

	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	c := jaegertracing.New(e, nil)
	defer c.Close()

	e.Use(middleware.Recover())

	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	e.Use(middleware.Gzip())
	//e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5,}))

	client, err := datasource.Clients()

	if err != nil {
		panic(err)
	}
	fmt.Println("dddd")

	fmt.Println(client)
	fmt.Println("eeee")
	ctx := context.Background()

	autoMigration := datasource.AutoMigration
	autoMigration(client, ctx)

	debugMode := datasource.DebugMode

	debugMode(err, client, ctx)

	//CreateUser(ctx, client)
	//println(CreateUser(ctx, client_model))

	fmt.Println("bbbbb")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world!!!")
	})
	e.GET("/users", handler.GetAllUser())
	e.POST("/user", handler.CreateUser())

	e.GET("/user/username", handler.GetUserByUserName())
	e.GET("/user/name", handler.GetUserByName())
	e.GET("/user", handler.GetUserByEmail())
	e.PUT("/user", handler.UpdateUser())

	e.Logger.Fatal(e.Start(":2020"))

}
