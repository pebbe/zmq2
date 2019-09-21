A Go interface to [ZeroMQ](http://www.zeromq.org/) version 2.

[![Go Report Card](https://goreportcard.com/badge/github.com/pebbe/zmq2)](https://goreportcard.com/report/github.com/pebbe/zmq2)
[![GoDoc](https://godoc.org/github.com/pebbe/zmq2?status.svg)](https://godoc.org/github.com/pebbe/zmq2)

Requires ZeroMQ version 2.1 or 2.2

For ZeroMQ version 4, see: http://github.com/pebbe/zmq4

For ZeroMQ version 3, see: http://github.com/pebbe/zmq3

Including all examples of [ØMQ - The Guide](http://zguide.zeromq.org/page:all).

Keywords: zmq, zeromq, 0mq, networks, distributed computing, message passing, fanout, pubsub, pipeline, request-reply

### See also

 * [Mangos](https://github.com/go-mangos/mangos) — An implementation in pure Go of the SP ("Scalable Protocols") protocols
 * [go-nanomsg](https://github.com/op/go-nanomsg) — Language bindings for nanomsg in Go
 * [goczmq](https://github.com/zeromq/goczmq) — A Go interface to CZMQ

## Requirements

zmq2 is just a wrapper for the ZeroMQ library. It doesn't include the
library itself. So you need to have ZeroMQ installed, including its
development files. On Linux and Darwin you can check this with (`$` is
the command prompt):

```
$ pkg-config --modversion libzmq
2.2.0
```

The Go compiler must be able to compile C code. You can check this
with:
```
$ go env CGO_ENABLED
1
```

You can't do cross-compilation. That would disable C.

## Install

    go get github.com/pebbe/zmq2

## Docs

 * [package help](http://godoc.org/github.com/pebbe/zmq2)
 * [wiki](https://github.com/pebbe/zmq4/wiki) (for zmq4)

## Support for ZeroMQ version 2.1

 * The following functions are not supported in ZeroMQ version 2.1, and will return an error:
  * (*Socket) GetRcvtimeo
  * (*Socket) GetSndtimeo
  * (*Socket) SetRcvtimeo
  * (*Socket) SetSndtimeo
