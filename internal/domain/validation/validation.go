package validation

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/lukluk/link-proxy/config"
	"github.com/lukluk/link-proxy/config/upstream"
)

type IValidation interface {
	ValidateErrorResponse(upstreamId string, body []byte) bool
}

type validation struct {
	cfg config.Config
}

func NewValidation(cfg config.Config) *validation {
	return &validation{
		cfg,
	}
}

func (v *validation) ValidateAdditionalErrorsFromResponse(upstreamId string, body []byte) bool {
	backend := v.getUpstreamById(upstreamId)
	if backend.Host == "" {
		return false
	}
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return false
	}
	for _, cond := range backend.AddErrorConditions {
		httpPayloadMatch := false

		jsonValue, ok := jsonParsed.Path(cond.ResponseBody.Path).Data().(string)
		if  !ok {
			jsonValue = ""
		}
		if cond.ResponseBody.Path != "" && cond.ResponseBody.Value != "" &&
			jsonValue == cond.ResponseBody.Value {
			httpPayloadMatch = true
		}
		return httpPayloadMatch
	}
	return false

}

func (v *validation) getUpstreamById(id string) upstream.Upstream {
	if val, ok := v.cfg.Upstreams[id]; ok {
		return val
	}
	return upstream.Upstream{}
}