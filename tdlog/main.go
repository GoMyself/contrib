package tdlog

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

var fc = &fasthttp.Client{
	Name:                     "TDlog",
	NoDefaultUserAgentHeader: true,
	TLSConfig:                &tls.Config{InsecureSkipVerify: true},
	MaxConnsPerHost:          2000,
	MaxIdleConnDuration:      5 * time.Second,
	MaxConnDuration:          5 * time.Second,
	ReadTimeout:              5 * time.Second,
	WriteTimeout:             5 * time.Second,
	MaxConnWaitTimeout:       5 * time.Second,
}

var (
	timeout       time.Duration
	internalUrl   string
	internalToken string
)

func httpDoTimeout(requestBody []byte, method string, requestURI string, headers map[string]string) ([]byte, int, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	req.SetRequestURI(requestURI)
	req.Header.SetMethod(method)

	switch method {
	case "POST":
		req.SetBody(requestBody)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	// time.Second * 30
	err := fc.DoTimeout(req, resp, timeout)

	return resp.Body(), resp.StatusCode(), err
}

func write(fields map[string]string, flags string) error {

	var b strings.Builder
	
	t := time.Now()
	b.WriteString("INSERT INTO zlog_")
	b.WriteString(t.Format("01-02"))
	b.WriteString(" USING zlog TAGS (")
	b.WriteByte('"')
	b.WriteString(fields["id"])
	b.WriteByte('"')
	b.WriteByte(',')
	b.WriteByte('"')
	b.WriteString(fields["project"])
	b.WriteByte('"')
	b.WriteByte(')')
	
	b.WriteString(" VALUES(")
	b.WriteString(fmt.Sprintf("%d,", t.UnixMilli()))
	b.WriteByte('"')
	b.WriteString(fields["filename"])
	b.WriteByte('"')
	b.WriteByte(',')
	b.WriteByte('"')
	b.WriteString(fields["content"])
	b.WriteByte('"')
	b.WriteByte(',')
	b.WriteByte('"')
	b.WriteString(fields["fn"])
	b.WriteByte('"')
	b.WriteByte(',')
	b.WriteByte('"')
	b.WriteString(flags)
	b.WriteByte('"')
	b.WriteByte(')')
	
	fmt.Println("b = ", b.String())
	headers := map[string]string{
		"Authorization": "Basic " + internalToken,
	}
	body, statusCode, err := httpDoTimeout([]byte(b.String()), "POST", "/rest/sql", headers)
	if err != nil {
		return err
	}
	if statusCode != fasthttp.StatusOK {
		return fmt.Errorf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
	
	fmt.Println("body = ", string(body))
	return nil
}

func New(url, token string) {
	internalUrl = url
	internalToken = token
	timeout = 5 * time.Second
}

func Info(fields map[string]string) error {
	return write(fields, "info")
}

func Warn(fields map[string]string) error {
	return write(fields, "warn")
}

func Error(fields map[string]string) error {
	return write(fields, "error")
}
