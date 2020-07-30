FROM golang

ADD . /go/src//github.com/bgoldovsky/teamer/service-duty-bot

RUN go install /go/src/github.com/bgoldovsky/teamer/service-duty-bot/cmd/service

ENTRYPOINT /go/bin/service