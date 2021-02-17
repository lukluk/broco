package validation

import "github.com/Jeffail/gabs/v2"

func (v *validation) ValidateSuccessResponse(upstreamId string, body []byte, respStatus int) bool {
	backend, err := v.cfg.GetUpstreamById(upstreamId)
	if err != nil {
		return false
	}
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return false
	}
	cond := backend.SuccessCondition
	httpStatusMatch := true
	httpPayloadMatch := true
	if cond.HttpStatus != 0 && cond.HttpStatus != respStatus {
		httpStatusMatch = false
	}

	jsonValue, ok := jsonParsed.Path(cond.ResponseBody.Path).Data().(string)
	if  !ok {
		jsonValue = ""
	}
	if cond.ResponseBody.Path != "" && cond.ResponseBody.Value != "" &&
		jsonValue != cond.ResponseBody.Value {
		httpPayloadMatch = false
	}
	return httpPayloadMatch && httpStatusMatch
}
