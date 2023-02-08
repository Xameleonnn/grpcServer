FROM golang:1.19

ADD . /go/src/github.com/Xameleonnn/grpcServer

RUN go install github.com/Xameleonnn/grpcServer

ENTRYPOINT ["/go/bin/server"]

EXPOSE 5300