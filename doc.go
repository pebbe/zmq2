/*
A Go interface to ZeroMQ (zmq, 0mq) version 2.

For ZeroMQ version 4, see: http://github.com/pebbe/zmq4

For ZeroMQ version 3, see: http://github.com/pebbe/zmq3

Requires ZeroMQ version 2.1 or 2.2

The following functions return ErrorNotImplemented in 0MQ version 2.1:

(*Socket)GetRcvtimeo, (*Socket)GetSndtimeo, (*Socket)SetRcvtimeo, (*Socket)SetSndtimeo

http://www.zeromq.org/

See also the wiki (for zmq4): https://github.com/pebbe/zmq4/wiki

A note on the use of a context:

This package provides a default context. This is what will be used by
the functions without a context receiver, that create a socket or
manipulate the context. Package developers that import this package
should probably not use the default context with its associated
functions, but create their own context(s). See: type Context.
*/
package zmq2
