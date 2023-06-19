## Envoy lua E2E test

## Build and run

```bash

//start envoyProxy configuration can be reloaded without rebuild images with only start/stop
$ make
$ make run
$ curl localhost:8000 | jq '.headers.errorheaders' 
retun a line
$ curl -H 'X-auth-request-access-token: works-for-me' localhost:8000 | jq '.headers.errorheaders' 
return empty
```
