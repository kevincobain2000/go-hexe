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
	PublicDir = "dist"
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
		filename := fmt.Sprintf("%s/%s", PublicDir, embedPath)

		filename = pkg.SlashIndexFile(filename)
		filename = pkg.ReplaceDoubleSlash(filename)

		content, err := publicDir.ReadFile(filename)
		if err != nil {
			return c.String(http.StatusNotFound, pkg.ErrorMsg404)
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
