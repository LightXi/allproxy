package packages

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"proxy/utils"
	"strings"
)

func RegisterJsdelivr(app *fiber.App) {
	app.All("/jsdelivr/*", func(c *fiber.Ctx) error {
		source := strings.TrimPrefix(c.Path(), "/jsdelivr")
		uri := fmt.Sprintf("https://cdn.jsdelivr.net%s", source)
		data, err := utils.Get(uri, nil)
		if err != nil {
			return c.SendString(fmt.Sprintf("error during get data: %s", err.Error()))
		}

		return c.SendString(data)
	})
}
