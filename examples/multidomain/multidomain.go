package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

var domains = make(map[string]fasthttp.RequestHandler)

func main() {
	server := &fasthttp.Server{

		Handler: func(ctx *fasthttp.RequestCtx) {
			h, ok := domains[string(ctx.Host())]
			if !ok {
				ctx.NotFound()
				return
			}
			h(ctx)
		},
	}

	cert, priv, err := fasthttp.GenerateTestCertificate("localhost:8080")
	if err != nil {
		panic(err)
	}
	domains["localhost:8080"] = func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("You are accessing to localhost:8080\n") //nolint:errcheck
	}

	err = server.AppendCertEmbed(cert, priv)
	if err != nil {
		panic(err)
	}

	cert, priv, err = fasthttp.GenerateTestCertificate("127.0.0.1")
	if err != nil {
		panic(err)
	}
	domains["127.0.0.1:8080"] = func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("You are accessing to 127.0.0.1:8080\n") //nolint:errcheck
	}

	err = server.AppendCertEmbed(cert, priv)
	if err != nil {
		panic(err)
	}

	fmt.Println(server.ListenAndServeTLS(":8080", "", ""))
}
