FROM golang:1.19

ADD . /go/src/github.com/Xameleonnn/grpcServer

RUN go install github.com/Xameleonnn/grpcServer@master

ENTRYPOINT ["/go/bin/grpcServer"]

EXPOSE 5300