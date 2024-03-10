package main

import (
	"log"
	"net/url"
	"strings"

	utls "github.com/refraction-networking/utls"
)

type CLIFlags struct {
	addr          string
	port          string
	cert          string
	key           string
	upstreamProxy string
	client        string
	ja3           string
	verbose       bool
}

var (
	Flags                  CLIFlags
	DefaultClientHelloID   utls.ClientHelloID
	DefaultClientHelloSpec *utls.ClientHelloSpec
	DefaultUpstreamProxy   *url.URL
)

func getClientHelloID(client string) (utls.ClientHelloID, bool) {
	clientArr := strings.Split(client, "-")
	if len(clientArr) != 2 {
		if clientArr[0] == "Custom" {
			return utls.HelloCustom, true
		}
		return utls.ClientHelloID{}, false
	}

	return utls.ClientHelloID{
		Client:  clientArr[0],
		Version: clientArr[1],
		Seed:    nil,
		Weights: nil,
	}, true
}

func setDefaultClientHelloID(client string) {
	clientHelloId, ok := getClientHelloID(client)
	if !ok {
		log.Fatalf("Invalid client format: %s", client)
	}

	DefaultClientHelloID = clientHelloId
}

func setDefaultClientHelloSpec(ja3 string) {
	if len(ja3) > 0 {
		spec, err := StringToSpec(ja3, "", false)
		if err != nil {
			log.Fatalf("Invalid ja3 format: %s", err)
		}

		DefaultClientHelloSpec = spec
	}
}

func setDefaultUpstreamProxy(upstreamProxy string) {
	proxyUrl, err := url.Parse(upstreamProxy)
	if err != nil {
		log.Fatalf("Invalid upstream proxy: %s", upstreamProxy)
	}

	DefaultUpstreamProxy = proxyUrl
}

func loadDefaultProxyConfig() {
	setDefaultClientHelloID(Flags.client)

	if len(Flags.ja3) > 0 {
		setDefaultClientHelloID("Custom-0")
		setDefaultClientHelloSpec(Flags.ja3)
	}

	if len(Flags.upstreamProxy) > 0 {
		setDefaultUpstreamProxy(Flags.upstreamProxy)
	}
}
