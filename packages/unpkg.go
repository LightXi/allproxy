package packages

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"proxy/utils"
	"strings"
)

func RegisterUnpkg(app *fiber.App) {
	app.All("/unpkg/*", func(c *fiber.Ctx) error {
		source := strings.TrimPrefix(c.Path(), "/unpkg")
		uri := fmt.Sprintf("https://unpkg.com%s", source)
		data, err := utils.Get(uri, nil)
		if err != nil {
			return Catch(c, err)
		}

		return End(c, data, func(data string) string {
			return strings.Replace(data, "https://unpkg.com", "/unpkg", -1)
		})
	})
}
