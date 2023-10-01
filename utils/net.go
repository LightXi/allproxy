package utils

import "github.com/valyala/fasthttp"

type Headers map[string]string

func Get(uri string, headers *Headers) (string, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}

	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, resp); err != nil {
		return "", err
	}

	return string(resp.Body()), nil
}
