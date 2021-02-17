package config

import (
	"errors"
	"fmt"
	"github.com/lukluk/link-proxy/config/upstream"
	"strings"
)

func (c *Config) FindUpstreamsByPathURL(path string) (upstream.Upstream, string, error) {
	for key, val := range c.Upstreams {
		if strings.Contains(path, key) {
			return val, key, nil
		}
	}
	return upstream.Upstream{}, "", errors.New("upstream not found")
}

func (c *Config) GetUpstreamById(id string) (upstream.Upstream, error) {
	if val, ok := c.Upstreams[id]; ok {
		return val, nil
	}
	return upstream.Upstream{}, errors.New(fmt.Sprintf("upstream id: %s, not found", id))
}
