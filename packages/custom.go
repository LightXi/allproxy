package packages

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"proxy/utils"
	"strings"
)

var CustomReplacer = map[string]string{
	"chatnio": "https://chatnio.net",
}

func RegisterCustom(app *fiber.App) {
	for _, v := range CustomReplacer {
		app.All(fmt.Sprintf("/%s/*", v), func(c *fiber.Ctx) error {
			source := strings.TrimPrefix(c.Path(), fmt.Sprintf("/%s", v))
			uri := fmt.Sprintf("%s%s", v, source)
			data, err := utils.Get(uri, nil)
			if err != nil {
				return c.SendString(fmt.Sprintf("error during get data: %s", err.Error()))
			}

			return c.SendString(data)
		})
	}
}
