FROM dongfeng

LABEL maintainer="ilovelili<route666@live.cn>"

ENV SRC_DIR=/go/src/github.com/ilovelili/dongfeng/core-proxy

WORKDIR $SRC_DIR

# Installing Go package manager
RUN go get github.com/Masterminds/glide

# Running glide
RUN glide up

WORKDIR $SRC_DIR/services/proxy

RUN go build