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
2019/09/24 17:50:22 Response: &{...}
```
