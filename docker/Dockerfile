FROM golang:1.14.15 AS builder

ADD . /go/src/github.com/irita-api/

RUN export GO111MODULE=on && export GOPROXY="https://goproxy.cn" && cd /go/src/github.com/irita-api && make irita-api

FROM alpine:3.15.0

COPY --from=builder /lib/x86_64-linux-gnu/libc.so.6 /lib/x86_64-linux-gnu/libc.so.6
COPY --from=builder /lib/x86_64-linux-gnu/libpthread.so.0 /lib/x86_64-linux-gnu/libpthread.so.0
COPY --from=builder /lib64/ld-linux-x86-64.so.2 /lib64/ld-linux-x86-64.so.2

COPY --from=builder /go/src/github.com/irita-api/bin/irita-api /usr/local/bin/
COPY --from=builder /go/src/github.com/irita-api/config/irita_config.yaml /irita-api/config/

WORKDIR /irita-api

CMD ["irita-api"]