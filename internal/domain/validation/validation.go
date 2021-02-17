package validation

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/lukluk/link-proxy/config"
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

func (v *validation) ValidateExtraErrorByResponse(upstreamId string, body []byte) bool {
	backend, err := v.cfg.GetUpstreamById(upstreamId)
	if err != nil {
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