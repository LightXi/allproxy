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
			return Catch(c, err)
		}

		return End(c, data, func(data string) string {
			return strings.Replace(data, "fonts/", "/jsdelivr/fonts/", -1)
		})
	})
}
