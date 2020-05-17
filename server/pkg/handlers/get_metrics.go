package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/glena/electric-monitor/pkg/db"
	"github.com/labstack/echo/v4"
)

func GetMetrics(database db.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		q := fmt.Sprintf("SELECT * FROM meassures WHERE time > now() - 1m")

		res, err := database.Query(q)
		if err != nil {
			log.Fatal(err)
		}

		return c.JSON(http.StatusOK, db.MapDataPoints(res))
	}
}
