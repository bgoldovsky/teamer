FROM golang

ADD . /go/src/github.com/bgoldovsky/dutyer/auth-api

RUN go install /go/src/github.com/bgoldovsky/dutyer/auth-api/cmd/service

ENTRYPOINT /go/bin/service