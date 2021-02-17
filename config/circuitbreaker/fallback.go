package circuitbreaker

type Fallback struct {
	HttpStatus int `yaml:"http_status" env-default:"503"`
	Response string `yaml:"response"`
}