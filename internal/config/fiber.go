package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetString("app.name"),
		ErrorHandler: NewErrorHandler(),
		Prefork:      config.GetBool("web.prefork"),
		ColorScheme:  fiber.DefaultColors,
	})
	//corsSettings := cors.New(cors.Config{
	//	AllowCredentials: true,
	//	AllowOrigins:     "",
	//	AllowMethods:     "GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH",
	//	AllowHeaders:     "Origin, Content-Type, Accept,  Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With",
	//	//   ExposeHeaders:    "Origin",
	//})
	//app.Use(corsSettings)

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
