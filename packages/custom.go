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
			resp, err := utils.Get(uri, nil)
			if err != nil {
				return Catch(c, err)
			}
			return End(c, resp)
		})
	}
}
