package packages

import (
	"crypto/tls"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/valyala/fasthttp"
)

func init() {
	proxy.WithTlsConfig(&tls.Config{
		InsecureSkipVerify: true,
	})
	// if you need to use global self-custom client, you should use proxy.WithClient.
	proxy.WithClient(&fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		DisablePathNormalizing:   true,
	})
}

func End(c *fiber.Ctx, resp *fasthttp.Response, handler ...func(data string) string) error {
	c.Response().Header.SetContentType(string(resp.Header.ContentType()))
	c.Response().Header.SetStatusCode(resp.StatusCode())

	data := string(resp.Body())
	for _, v := range handler {
		data = v(data)
	}
	return c.SendString(data)
}

func Catch(c *fiber.Ctx, err error) error {
	return c.SendString(fmt.Sprintf("error during get data: %s", err.Error()))
}
