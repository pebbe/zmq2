#if ZMQ_VERSION_MAJOR != 2
#error "You need ZeroMQ version 2 to build this"
#endif
#ifndef ZMQ_RCVTIMEO
#define ZMQ_RCVTIMEO 0
#endif
#ifndef ZMQ_SNDTIMEO
#define ZMQ_SNDTIMEO 0
#endif
