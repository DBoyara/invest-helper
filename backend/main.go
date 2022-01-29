package main

import (
	"fmt"
	"os"
	"time"

	"github.com/DBoyara/invest-helper/common"
	"github.com/DBoyara/invest-helper/pkg/repository"
	"github.com/DBoyara/invest-helper/router"
	"github.com/DBoyara/invest-helper/third_party"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var configDefaultCORS = cors.Config{
	AllowOrigins: "*",
	AllowMethods: "GET,POST",
	AllowHeaders: "*",
}

var configDefaultLogger = fiberLogger.Config{
	Format:       "${red}[${time}] ${green}${status} - ${blue}${latency} ${method} ${path}${reset}\n",
	TimeFormat:   "15:04:05",
	TimeZone:     "Local",
	TimeInterval: 500 * time.Millisecond,
	Output:       os.Stderr,
}

func main() {
	logger := third_party.SetLogConf()
	app := fiber.New()
	settings := common.GetSettings()

	app.Use(requestid.New())
	app.Use(cors.New(configDefaultCORS))
	app.Use(fiberLogger.New(configDefaultLogger))

	router.SetupRoutes(app)
	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	if err := repository.SetupDB(settings); err != nil {
		logger.Panic().Msgf("failed to connect to db: %s", err)
	}
	defer repository.CloseDB()

	if err := app.Listen(fmt.Sprintf(":%s", settings.HttpPort)); err != nil {
		logger.Fatal().Err(err)
	}

	if err := app.Shutdown(); err != nil {
		logger.Panic().Err(err)
	}
}
