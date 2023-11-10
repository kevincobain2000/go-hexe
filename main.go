package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/kevincobain2000/go-hexe/pkg"
	"github.com/labstack/echo/v4"
)

const (
	PUBLIC_DIR   = "dist"
	DEFAULT_PORT = "3000"
)

//go:embed all:dist/*
var publicDir embed.FS

func main() {

	e := pkg.NewEcho()

	e.GET("*", func(c echo.Context) error {
		filename := fmt.Sprintf("%s%s", PUBLIC_DIR, c.Request().URL.Path)

		filename = pkg.SlashIndexFile(filename)

		content, err := publicDir.ReadFile(filename)
		if err != nil {
			return c.String(http.StatusNotFound, pkg.ERROR_MSG_404_PAGE_NOT_FOUND)
		}

		return c.Blob(http.StatusOK, pkg.GetContentType(filename), content)

	})

	port := getPort()
	pkg.GracefulServerWithPid(e, port)
}

func getPort() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return DEFAULT_PORT
}
