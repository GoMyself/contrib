package helper

import (
	"github.com/valyala/fasthttp"
	"time"
)

const (
	apiTimeOut = time.Second * 12
)

func httpPost(requestBody []byte, requestURI string, headers map[string]string) (int, []byte, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	req.SetRequestURI(requestURI)
	req.Header.SetMethod("POST")
	req.SetBody(requestBody)

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	err := fc.DoTimeout(req, resp, apiTimeOut)

	return resp.StatusCode(), resp.Body(), err
}

func httpGet(requestURI string) (int, []byte, error) {

	return fc.GetTimeout(nil, requestURI, apiTimeOut)
}

func httpGetHeader(requestURI string, headers map[string]string) (int, []byte, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	req.SetRequestURI(requestURI)
	req.Header.SetMethod("GET")

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	err := fc.DoTimeout(req, resp, apiTimeOut)

	return resp.StatusCode(), resp.Body(), err
}
