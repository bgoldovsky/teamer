FROM golang

ADD . /go/src/github.com/bgoldovsky/teamer/auth-api

RUN go install /go/src/github.com/bgoldovsky/teamer/auth-api/cmd/service

ENTRYPOINT /go/bin/service