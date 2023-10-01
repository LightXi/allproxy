package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"proxy/packages"
)

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Sprintf("error during read config: %s", err.Error()))
		return
	}

	app := fiber.New()
	app.Static("/static", "./static")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	{
		packages.RegisterJsdelivr(app)
		packages.RegisterFonts(app)
	}

	if err := app.Listen(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		fmt.Println(fmt.Sprintf("error during listen: %s", err.Error()))
	}
}
