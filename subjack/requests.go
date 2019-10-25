package subjack

import (
	"crypto/tls"
	"time"

	"github.com/valyala/fasthttp"
)

func get(url string, ssl bool, followRedirects bool, timeout int) (body []byte) {
	client := &fasthttp.Client{TLSConfig: &tls.Config{InsecureSkipVerify: true}}

	if followRedirects {
		_, body, _ := client.GetTimeout(nil, site(url, ssl), time.Duration(timeout)*time.Second)
		return body
	} else {
		req := fasthttp.AcquireRequest()
		req.SetRequestURI(site(url, ssl))
		req.Header.Add("Connection", "close")
		resp := fasthttp.AcquireResponse()
		client.DoTimeout(req, resp, time.Duration(timeout)*time.Second)
		return resp.Body()
	}
}

func site(url string, ssl bool) (site string) {
	site = "http://" + url
	if ssl {
		site = "https://" + url
	}

	return site
}
