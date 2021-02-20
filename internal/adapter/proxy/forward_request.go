package proxy

import (
	"github.com/gojektech/heimdall/v6/httpclient"
	"github.com/lukluk/link-proxy/internal/errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func ForwardRequest (host string, req *http.Request, timeout time.Duration) (*http.Response, error) {
	url := req.URL
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		)
	proxyReq, err := http.NewRequest(req.Method, host + url.RequestURI(), req.Body)
	if err != nil {
		log.Error().Msgf("cannot create request to upstream, error:%v", err)
		return nil, errors.NewRequestError()
	}

	proxyReq.Header.Set("Host", req.Host)
	proxyReq.Header.Set("X-Forwarded-For", req.RemoteAddr)

	for header, values := range req.Header {
		for _, value := range values {
			proxyReq.Header.Add(header, value)
		}
	}

	return client.Do(proxyReq)

}