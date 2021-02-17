package handler

import (
	"crypto/md5"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/lukluk/link-proxy/config/upstream"
	"io"
	"io/ioutil"
	"net/http"
)

func (e *entryPoint) buildRequestKey(backend upstream.Upstream, r *http.Request) (string, error) {
	var payloadValues string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return hashMd5(r.URL.String()), nil
	}
	for _, path := range backend.CircuitBreakerKey.RequestBodyJsonPaths {
		payloadValues = payloadValues + jsonParsed.Path(path).Data().(string)
	}
	return hashMd5(fmt.Sprintf("%s%s",  r.URL.String(), payloadValues)), nil
}

func hashMd5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return string(h.Sum(nil))
}
