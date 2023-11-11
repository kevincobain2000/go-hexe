package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"

	"github.com/kevincobain2000/go-hexe/pkg"
	"github.com/labstack/echo/v4"
)

const (
	PUBLIC_DIR = "dist"
)

var (
	port     string
	basePath string
)

//go:embed all:dist/*
var publicDir embed.FS

func main() {
	cliArgs()
	e := pkg.NewEcho()

	e.GET(basePath+"*", func(c echo.Context) error {
		embedPath := c.Request().URL.Path
		embedPath = embedPath[len(basePath):]
		filename := fmt.Sprintf("%s/%s", PUBLIC_DIR, embedPath)

		filename = pkg.SlashIndexFile(filename)

		content, err := publicDir.ReadFile(filename)
		if err != nil {
			return c.String(http.StatusNotFound, pkg.ERROR_MSG_404_PAGE_NOT_FOUND)
		}

		return c.Blob(http.StatusOK, pkg.GetContentType(filename), content)

	})

	pkg.GracefulServerWithPid(e, port)
}

func cliArgs() {
	flag.StringVar(&port, "port", "3000", "port to serve")
	flag.StringVar(&basePath, "base-path", "/", "base path")
	flag.Parse()
}
