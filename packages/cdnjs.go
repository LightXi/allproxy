package packages

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"proxy/utils"
	"strings"
)

func RegisterCdnJs(app *fiber.App) {
	app.All("/cdnjs/*", func(c *fiber.Ctx) error {
		source := strings.TrimPrefix(c.Path(), "/cdnjs")
		uri := fmt.Sprintf("https://cdnjs.cloudflare.com%s", source)
		data, err := utils.Get(uri, nil)
		if err != nil {
			return Catch(c, err)
		}

		return End(c, data, func(data string) string {
			return strings.Replace(data, "/ajax/libs", "/cdnjs/ajax/libs", -1)
		})
	})
}
