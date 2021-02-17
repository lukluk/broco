package upstream

type condition struct {
	HttpStatus   int           `yaml:"http_status"`
	ResponseBody JsonPathValue `yaml:"response_body"`
}