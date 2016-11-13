FROM ubuntu:xenial
ADD ./signal-logger /usr/bin/signal-logger
ENTRYPOINT [ "/usr/bin/signal-logger" ]
