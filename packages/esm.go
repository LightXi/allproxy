package packages

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"proxy/utils"
	"strings"
)

func RegisterEsm(app *fiber.App) {
	app.All("/esm/*", func(c *fiber.Ctx) error {
		source := strings.TrimPrefix(c.Path(), "/esm")
		uri := fmt.Sprintf("https://cdn.jsdelivr.net/npm%s", source)
		data, err := utils.Get(uri, nil)
		if err != nil {
			return Catch(c, err)
		}

		return End(c, data, func(data string) string {
			return strings.Replace(data, "https://cdn.jsdelivr.net/npm", "/esm", -1)
		})
	})
}
