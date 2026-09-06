package main

import (
	"crypto/tls"
	"net"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

func requestHandler(ctx *fasthttp.RequestCtx) { _ = "STUB: not implemented"; return }

func main() {
	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.com"),
		Cache:      autocert.DirCache("./certs"),
	}

	cfg := &tls.Config{
		GetCertificate: m.GetCertificate,
		NextProtos: []string{
			"http/1.1", acme.ALPNProto,
		},
	}

	ln, err := net.Listen("tcp4", "0.0.0.0:443")
	if err != nil {
		panic(err)
	}

	lnTLS := tls.NewListener(ln, cfg)

	if err := fasthttp.Serve(lnTLS, requestHandler); err != nil {
		panic(err)
	}
}
