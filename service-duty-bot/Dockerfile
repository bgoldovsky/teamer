FROM golang

ADD . /go/src//github.com/bgoldovsky/dutyer/service-duty-bot

RUN go install /go/src/github.com/bgoldovsky/dutyer/service-duty-bot/cmd/service

ENTRYPOINT /go/bin/service