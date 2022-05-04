package applestore

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func HttpRequest(token string, method string, url string, body []byte) (*fasthttp.Response, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetContentType("application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.SetMethod(method)
	req.SetRequestURI(url)
	req.SetBody(body)
	if err := fasthttp.DoTimeout(req, resp, 20*time.Second); err != nil {
		return nil, err
	}
	return resp, nil
}
