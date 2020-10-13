FROM golang

ADD . /go/src//github.com/bgoldovsky/dutyer/service-slack-bot

RUN go install /go/src/github.com/bgoldovsky/dutyer/service-slack-bot/cmd/service

ENTRYPOINT /go/bin/service