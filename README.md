# Broco Proxy

![](https://i.imgur.com/bwoP01W.jpg)
# How to run
```
broco-proxy --config=config.yaml
```

# Config
## port
Proxy server port (default 8080)
## circuit_breaker
configuration for circuit breaker logic
### thresholds
max_error_percentage (default 70)

min_error_percentage (default 10)

#### durations
state_interval (default 1m)

> timer for circuit breaker take action based on thresholds

closed_duration (default 15m)

#### fallback
http_status (default 503)

### upstreams
#### host
upstream host
#### add_error_conditions

response_body : path, value

> example: path: data.status, value: 00

#### circuit_breaker_key
request_body_json_paths

> circuit_breaker_key is used when the endpoint too generic or multiple purpose, and the differentiate is on payload

# Example Config

```yaml
port: 8080
circuit_breaker:
  thresholds:
    max_error_percentage: 70
    min_success_percentage: 90
  durations:
    state_interval: 1m
    closed_duration: 1m
  fallback:
    http_status: 503    
statsd_host: http://localhost:8125
upstreams:
  /v1/foo:
    host: http://foo:3000
    add_error_conditions:      
      - response_body:
          path: error.code
          value: 111
      - response_body:
          path: error.code
          value: 222
  /v1/bar:
    host: http://bar:3000
    circuit_breaker_key:
      request_body_json_paths:
        - bank

```

