# HTTP Mock Server

Runs a HTTP mock server based on endpoints specifications specified in file or standard input. A server runs on specified port or default one if not specified.

Default port is `8888`.

## Usage

To start a mock server on specific port while mocked endpoints are specified in a file, just follow the examples below.

```shell
mhttp serve -p 4444 -f mock_spec.yaml

# to run on default port
mhttp serve -f mock_spec.yaml
```
```shell
mhttp serve -p 4444 -f - <<EOF
spec:
  - type: RestRequestResponse
    id: "standard error with json response"
    request:
      verb: GET
      path: /error
    response:
      code: 500
      headers:
        - key: "Content-Type"
          value: "application/json"
      body:
        type: string
        ## a json string response
        value: '{"error": "standard error response"}'
EOF
```

## Mock specification

A mock specification should be expressed as YAML file wich pre-defined sections of type `RestRequestResponse`. These sections specify a request match and a specific response pairs.  

See an example of simple mock server below.

```yaml

spec:
  ## an array with a request-response specifications 
  - type: RestRequestResponse 
    ## id is mandatory and must be unique across the server instance
    id: 'test-endpoint-id'
    ## defines a request section
    request:
      verb: GET
      path: /hello-world
    ## defines a matched response 
    response:
      code: 200
      headers:
        - key: "Content-Type"
          value: "text/plain; charset=utf-8"
      body:
        # type of the body value specified in the configuration file
        # type string puts a specified string directly to the response 
        type: string
        value: 'hello-wold'

```
