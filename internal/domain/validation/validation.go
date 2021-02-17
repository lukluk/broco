package validation

import "github.com/lukluk/link-proxy/config"

type IValidation interface {
	ValidateSuccessResponse(upstreamId string, body []byte, respStatus int) bool
}

type validation struct {
	cfg config.Config
}

func NewValidation(cfg config.Config) *validation {
	return &validation{
		cfg,
	}
}
