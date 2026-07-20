package main

import (
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

func main() {

	url := fasthttp.AcquireURI()
	url.Parse(nil, []byte("http://localhost:8080/")) //nolint:errcheck
	url.SetUsername("Aladdin")
	url.SetPassword("Open Sesame")

	hc := &fasthttp.HostClient{
		Addr: "localhost:8080",
	}

	req := fasthttp.AcquireRequest()
	req.SetURI(url)
	fasthttp.ReleaseURI(url)

	req.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	err := hc.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err == nil {
		fmt.Printf("Response: %s\n", resp.Body())
	} else {
		fmt.Fprintf(os.Stderr, "Connection error: %v\n", err)
	}
	fasthttp.ReleaseResponse(resp)
}
