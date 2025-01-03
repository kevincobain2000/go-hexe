package pkg

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/fvbock/endless"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	SetupLogger(e)
	return e
}
func SetupLogger(e *echo.Echo) {
	log := logrus.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(_ echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))
}

// GracefulServerWithPid reloads server with pid
// kill -HUP when binary is changed
// kill -9 when want to kill the process and make the application dead and want to restart
// kill -9 is NOT FOR FAINT HEARTED and must not be done on prod unless SOUT
func GracefulServerWithPid(e *echo.Echo, port string) {
	log := logrus.New()
	server := endless.NewServer("localhost:"+port, e)
	server.BeforeBegin = func(_ string) {
		pidFile := filepath.Join(port + ".pid") //nolint: gocritic
		_ = os.Remove(pidFile)
		err := os.WriteFile(pidFile, []byte(strconv.Itoa(os.Getpid())), 0600)
		if err != nil {
			log.Error("write pid file error: ", err)
		}
		log.Info("started server on localhost:", port)
	}
	if err := server.ListenAndServe(); err != nil {
		log.Error("graceful error: ", err)
	}
}
