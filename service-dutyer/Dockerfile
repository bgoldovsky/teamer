FROM golang

ADD . /go/src//github.com/bgoldovsky/dutyer/service-dutyer

#ENV PORT=50051
#ENV CONNECTION_STRING=user=postgres password=postgres dbname=service_dutyer sslmode=disable
#EXPOSE 50051

RUN go install /go/src/github.com/bgoldovsky/dutyer/service-dutyer/cmd/service

ENTRYPOINT /go/bin/service