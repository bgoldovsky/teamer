FROM golang

ADD . /go/src/github.com/bgoldovsky/teamer/gateway-api

RUN go install /go/src/github.com/bgoldovsky/teamer/gateway-api/cmd/service

ENTRYPOINT /go/bin/service