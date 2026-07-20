package main

import (
	"time"

	"github.com/valyala/fasthttp"
)

var headerContentTypeJSON = []byte("application/json")

var client *fasthttp.Client

type Entity struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func main() {

	readTimeout, _ := time.ParseDuration("500ms")
	writeTimeout, _ := time.ParseDuration("500ms")
	maxIdleConnDuration, _ := time.ParseDuration("1h")
	client = &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,

		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}
	sendGetRequest()
	sendPostRequest()
}

func sendGetRequest() { _ = "STUB: not implemented"; return }

func sendPostRequest() { _ = "STUB: not implemented"; return }

//nolint:errchkjson

func httpConnError(err error) (string, bool) { _ = "STUB: not implemented"; return "", false }
