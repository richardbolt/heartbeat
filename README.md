Heartbeat
=========

Heartbeat is a small module to use with a [Goa](http://goa.design/) app to provide a heartbeat to ping for health check purposes.

By adding Heartbeat your app will respond with a 200 OK and a JSON object like so: `{"ENV":"production"}` on the url you specify or /health if no url is specified. Heartbeat returns a simple JSON object with one key, ENV, having the value of the ENV environment variable or "dev" if ENV has no value.

Heartbeat is not a middleware, instead it adds itself to service.Mux as an http handler.

Usage
-----

```
import (
  "github.com/goadesign/goa"
  "github.com/richardbolt/heartbeat"
)

service := goa.New("API")
heartbeat.Heartbeat(service, "/health")

service.ListenAndServe(":8080")
```

The usage signature of Heartbeat is:

```
Heartbeat(*goa.Service, string)
```
