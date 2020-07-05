package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func upload(c echo.Context) error {

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files)))
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})
	e.GET("/file", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.Logger.Fatal(e.Start(":4000"))
}
