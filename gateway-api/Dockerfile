FROM golang

ADD . /go/src/github.com/bgoldovsky/dutyer/gateway-api

RUN go install /go/src/github.com/bgoldovsky/dutyer/gateway-api/cmd/service

ENTRYPOINT /go/bin/service