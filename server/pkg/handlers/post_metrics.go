package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/glena/electric-monitor/pkg/db"
	"github.com/labstack/echo/v4"
)

func PostMetrics(database db.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		// Create a point and add to batch
		tags := map[string]string{}

		p, err := strconv.ParseFloat(c.FormValue("p"), 64)
		if err != nil {
			fmt.Println("parse p")
			log.Fatal(err)
		}
		v, err := strconv.ParseFloat(c.FormValue("v"), 64)
		if err != nil {
			fmt.Println("parse v")
			log.Fatal(err)
		}
		i, err := strconv.ParseFloat(c.FormValue("i"), 64)
		if err != nil {
			fmt.Println("parse i")
			log.Fatal(err)
		}

		fields := map[string]interface{}{
			"p": p,
			"v": v,
			"i": i,
		}

		err = database.InsertPoint("meassures", tags, fields)

		if err != nil {
			fmt.Println("Failed to push data")
			log.Fatal(err)
		}

		return c.NoContent(204)
	}
}
