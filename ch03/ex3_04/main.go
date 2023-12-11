package main

import (
	"ch03/example3"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/:number", func(c echo.Context) error {
		nstr := c.Param("number")
		n, err := strconv.Atoi(nstr)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, strconv.FormatBool(example3.IsPrime(n)))
	})

	e.Logger.Fatal(e.Start(":6014"))
}
