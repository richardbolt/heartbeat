Heartbeat
=========

Heartbeat is a small module to use with a Goa app to provide a heartbeat to ping for health check purposes.

By adding Heartbeat your app will respond with a 200 OK and a JSON object like so: `{"ENV":"production"}` on the url you specify or /heartbeat if no url is specified.

Heartbeat is not a middleware, instead it adds itself to service.Mux as an http handler.

Usage
-----

```
import (
  "github.com/goadesign/goa"
  "github.com/richardbolt/heartbeat"
)

service := goa.New("API")
heartbeat.Heartbeat(service, "")

service.ListenAndServe(":8080")
```

The usage signature of Heartbeat is:

```
Heartbeat(*goa.Service, string)
```
