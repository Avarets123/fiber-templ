package main

import (
	"fiber-templ/config"
	"fiber-templ/internal/home"
	"fiber-templ/pkg/logger/zlogger"
	"log"
	"strings"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/rs/zerolog"
)

func main() {
	err := config.Init()
	logCfg := config.NewLoggerCfg()
	zlogger := zlogger.New(logCfg)

	if err != nil {
		zlogger.Warn().Msg(".env file not loaded!")
	} else {
		zlogger.Info().Msg(".env file loaded!")
	}

	config.NewDbCfg()

	app := configureFiber(zlogger)
	app.Static("/public", "./public")

	home.ApplyHanlder(app)

	zlogger.Info().Msg("App listening port: 3000")
	err = app.Listen(":3000")
	if err != nil {
		log.Panicln(err)
	}

}

func configureFiber(zlogger *zerolog.Logger) *fiber.App {

	engine := html.New("./html", ".html")
	engine.AddFuncMap(map[string]any{
		"ToUpper": strings.ToUpper,
	})

	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 10,
		Views:        engine,
	})

	errHandler := recover.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: zlogger,
	}))

	app.Use(errHandler)

	return app
}
