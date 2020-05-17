package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/glena/electric-monitor/pkg/db"
	"github.com/glena/electric-monitor/pkg/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	HOST := os.Getenv("HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	PORT := os.Getenv("PORT")
	CLIENT_NAME := os.Getenv("CLIENT_NAME")
	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")

	e := echo.New()

	database, err := db.Init(HOST, DB_NAME, DB_USER, DB_PASSWORD)
	defer database.Disconnect()

	if err != nil {
		fmt.Println("Failed to initialize te database")
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	e.Pre(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte(CLIENT_NAME)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(CLIENT_SECRET)) == 1 {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	e.GET("/metrics", handlers.GetMetrics(database))
	e.POST("/metrics", handlers.PostMetrics(database))

	e.Logger.Fatal(e.Start(":" + PORT))
}
