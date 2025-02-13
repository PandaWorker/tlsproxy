package main

import (
	"net/http"

	"github.com/elazarl/goproxy"
)

func invalidClientResponse(req *http.Request, ctx *goproxy.ProxyCtx, client string) *http.Response {
	ctx.Logf("Client specified invalid client: %s", client)
	return goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusBadRequest, "Invalid client: "+client)
}

func invalidJA3StringResponse(req *http.Request, ctx *goproxy.ProxyCtx, client string) *http.Response {
	ctx.Logf("JA3String specified invalid: %s", client)
	return goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusBadRequest, "Invalid client: "+client)
}

func invalidClientHelloSpecResponse(req *http.Request, ctx *goproxy.ProxyCtx, client string) *http.Response {
	ctx.Logf("ClientHelloSpec specified invalid: %s", client)
	return goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusBadRequest, "Invalid client: "+client)
}

func invalidUpstreamProxyResponse(req *http.Request, ctx *goproxy.ProxyCtx, upstreamProxy string) *http.Response {
	ctx.Logf("Client specified invalid upstream proxy: %s", upstreamProxy)
	return goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusBadRequest, "Invalid upstream proxy: "+upstreamProxy)
}
