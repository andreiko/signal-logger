FROM ubuntu:xenial
ENV DEBIAN_FRONTEND=noninteractive
ENV APT_MIRROR=us-west-2.ec2.archive.ubuntu.com

RUN cat /etc/apt/sources.list | \
    sed "s/archive\.ubuntu\.com/${APT_MIRROR}/g" | grep -v '^deb-src' > /tmp/sources.list && \
    mv /tmp/sources.list /etc/apt/sources.list

RUN apt-get update -y && apt-get install -y ca-certificates

ADD ./signal-logger /usr/bin/signal-logger
ENTRYPOINT [ "/usr/bin/signal-logger" ]
