package packages

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"proxy/utils"
	"strings"
)

var FontsReplacer = map[string]string{
	"Jetbrains+Mono": "JetBrainsMono-Regular.woff2",
}

func RegisterFonts(app *fiber.App) {
	app.Get("/fonts/:font", func(c *fiber.Ctx) error {
		font := strings.Replace(c.Params("font"), "-", "+", -1)
		for k, _ := range FontsReplacer {
			if font == k {
				return c.Redirect(fmt.Sprintf("/static/fonts/%s.css", k))
			}
		}

		uri := fmt.Sprintf("https://fonts.googleapis.com/css?family=%s", font)
		data, err := utils.Get(uri, nil)
		if err != nil {
			return c.SendString(fmt.Sprintf("error during get data: %s", err.Error()))
		}
		data = strings.Replace(data, "https://fonts.gstatic.com", "/gstatic", -1)

		return c.SendString(data)
	})

	app.All("/gstatic/*", func(c *fiber.Ctx) error {
		source := strings.TrimPrefix(c.Path(), "/gstatic")
		for _, v := range FontsReplacer {
			if strings.Contains(source, v) {
				return c.Redirect(fmt.Sprintf("/static/fonts/source/%s", v))
			}
		}

		uri := fmt.Sprintf("https://fonts.gstatic.com%s", source)
		data, err := utils.Get(uri, nil)
		if err != nil {
			return c.SendString(fmt.Sprintf("error during get data: %s", err.Error()))
		}

		return c.SendString(data)
	})
}
