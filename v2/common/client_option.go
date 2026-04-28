package common

import (
	"net/http"
	"net/url"
)

type ProxyFunc func(*http.Request) (*url.URL, error)

type ClientConfig struct {
	UseTestnet   bool
	RoundTripper http.RoundTripper
	ProxyFunc    ProxyFunc
}

func ParseClientConfig(opts ...ClientOptionFunc) ClientConfig {
	cfg := ClientConfig{UseTestnet: false}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

type ClientOptionFunc func(*ClientConfig)

func UseTestnet(useTestnet bool) ClientOptionFunc {
	return func(cfg *ClientConfig) {
		cfg.UseTestnet = useTestnet
	}
}

func WithRoundTripper(rt http.RoundTripper) ClientOptionFunc {
	return func(cfg *ClientConfig) {
		cfg.RoundTripper = rt
	}
}

func WithProxyFunc(proxyFunc ProxyFunc) ClientOptionFunc {
	return func(cfg *ClientConfig) {
		cfg.ProxyFunc = proxyFunc
	}
}
