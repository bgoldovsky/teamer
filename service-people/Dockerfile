FROM golang

ADD . /go/src//github.com/bgoldovsky/teamer/service-people

#ENV PORT=50051
#ENV CONNECTION_STRING=user=postgres password=postgres dbname=service_rus_id sslmode=disable
#EXPOSE 50051

RUN go install /go/src/github.com/bgoldovsky/teamer/service-people/cmd/service

ENTRYPOINT /go/bin/service