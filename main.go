package main

import (
    "github.com/danieljuvito/iohub-server/internal/controller"
    "github.com/danieljuvito/iohub-server/internal/controller/middleware"
    "github.com/danieljuvito/iohub-server/internal/repository"
    "github.com/danieljuvito/iohub-server/internal/service"
    "github.com/danieljuvito/iohub-server/internal/util/mongo"
    "github.com/joho/godotenv"
    "github.com/labstack/echo/v4"
    "log"
    "os"
    "strconv"
    "strings"

    _ "github.com/danieljuvito/iohub-server/docs"
)

// @title IOHub Server
// @version 1.0
// @description This is API server for IOHub App.
// @termsOfService http://swagger.io/terms/

// @contact.name Daniel Juvito
// @contact.url danieljuvito.github.io
// @contact.email danieljuvito@outlook.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbURL := "mongodb://localhost:27017"
    if dbURLEnv := strings.TrimSpace(os.Getenv("DB_URL")); dbURLEnv != "" {
        dbURL = dbURLEnv
    }

    // Get Client, Context, CancelFunc and 
    // err from connect method.
    client, ctx, cancel, err := mongo.Connect(dbURL)
    if err != nil {
        panic(err)
    }

    // Release resource when the main
    // function is returned.
    defer mongo.Close(client, ctx, cancel)

    // Ping mongoDB with Ping method
    err = mongo.Ping(client, ctx)
    if err != nil {
        panic(err)
    }

    db := client.Database("iohub")

    middleware.InitAuth(db)

    e := echo.New()
    e.GET("/ws", func(c echo.Context) error {
        
    })

    r := repository.NewRepository(db)

    controller.NewController(e,
        service.NewService(
            repository.NewRepository(db),
        ),
    )

    apiPort := "8080"
    if apiPortEnv, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
        apiPort = strconv.Itoa(apiPortEnv)
    }

    err = e.Start("0.0.0.0:" + apiPort)
    if err != nil {
        panic(err)
    }
}
