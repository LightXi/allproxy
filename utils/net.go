package utils

import "github.com/valyala/fasthttp"

type Headers map[string]string

func Get(uri string, headers *Headers) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}

	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
