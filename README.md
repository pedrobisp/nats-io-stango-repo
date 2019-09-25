# nats-io-stango-repo

This repository provides minimalist examples on how to reproduce an issue concerning https requests with the NATS streaming go API. This issue occurs only on **Linux** and **macOS**, not on Windows.

First, make sure that you have a nats-streaming-server running locally with the default parameters (port ```4222```, cluster id ```test-cluster```).

First, run the subscriber:

```
user@hostname:~/go/src/github.com/pedrobisp/nats-io-stango-repo/subscriber$ go run subscriber/main.go
```

Then, in another terminal run the publisher:

```
user@hostname:~/go/src/github.com/pedrobisp/nats-io-stango-repo/subscriber$ go run publisher/main.go
```

In the terminal running the subscriber, you are going to see the logging hanging after the request in sent:

```
user@hostname:~/go/src/github.com/pedrobisp/nats-io-stango-repo/subscriber$ go run main.go
2019/09/24 17:50:21 Sending request
``` 

If you change the protocol to http in subscriber/main.go (line 13), the request in the handler is not going to hang anymore:

```
user@hostname:~/go/src/github.com/pedrobisp/nats-io-stango-repo/subscriber$ go run main.go
2019/09/24 17:50:21 Sending request
2019/09/24 17:50:22 Response: &{200 OK 200 HTTP/2.0 2 0 map[Alt-Svc:[quic=":443"; ma=2592000; v="46,43"] Cache-Control:[private, max-age=0] Content-Type:[text/html; charset=ISO-8859-1] Date:[Wed, 25 Sep 2019 12:39:53 GMT] Expires:[-1] P3p:[CP="This is not a P3P policy! See g.co/p3phelp for more info."] Server:[gws] Set-Cookie:[1P_JAR=2019-09-25-12; expires=Fri, 25-Oct-2019 12:39:53 GMT; path=/; domain=.google.com; SameSite=none NID=188=nTIv7bPsnkmHVZgVYmzJ9N9l6m09G8B5J79Y9HyiznGX8WfzvA34mpX--YKPBrC7bJoWHFMVYp3pvnYjE_GoOeE6w-Xiv995AGjG4qn3oKaJ6NHwPfgrACrJ_Ra3PyG2X1tmoMWAL4FySGJ6K60-bgE-ZBni2iXvxZ2DYcjGvMo; expires=Thu, 26-Mar-2020 12:39:53 GMT; path=/; domain=.google.com; HttpOnly] X-Frame-Options:[SAMEORIGIN] X-Xss-Protection:[0]] 0xc00018e4b0 -1 [] false true map[] 0xc000124100 0xc0001f6210}
```
