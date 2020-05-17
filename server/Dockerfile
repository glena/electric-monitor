FROM golang:1.14

ENV GO111MODULE=on
ENV PORT=3000
ENV WATCHER_VERSION 0.2.4

SHELL ["/bin/bash", "-c"]

ENV GOPATH /go
VOLUME /go/src
WORKDIR /go/src

ADD https://github.com/canthefason/go-watcher/releases/download/v${WATCHER_VERSION}/watcher-${WATCHER_VERSION}-linux-amd64 /usr/local/bin/watcher
RUN chmod +x /usr/local/bin/watcher

ADD entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]
CMD ["watcher"]
