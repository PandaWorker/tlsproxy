package main

import "net/http"

const (
	UpStreamProxyHeader  = "x-tlsproxy-upstream"
	ClientProfileHeader  = "x-tlsproxy-client"
	JA3StringHeader      = "x-tlsproxy-ja3"
	RawClientHelloHeader = "x-tlsproxy-raw"
)

var CustomHeaders = []string{UpStreamProxyHeader, ClientProfileHeader, JA3StringHeader, RawClientHelloHeader}

type ProxyConfig struct {
	client          string
	ja3             string
	clientHelloSpec string
	upstreamProxy   string
}

func parseCustomHeaders(headers *http.Header) ProxyConfig {
	return ProxyConfig{
		upstreamProxy:   headers.Get(UpStreamProxyHeader),
		client:          headers.Get(ClientProfileHeader),
		ja3:             headers.Get(JA3StringHeader),
		clientHelloSpec: headers.Get(RawClientHelloHeader),
	}
}

func removeCustomHeaders(headers *http.Header) {
	for _, header := range CustomHeaders {
		headers.Del(header)
	}
}
