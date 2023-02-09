package main

import (
	"context"
	"flag"
	"fmt"
	tester "github.com/Xameleonnn/grpctester"
	"google.golang.org/grpc"
	"log"
	"net"
)

type S struct {
	tester.UnimplementedHandshakerServer
}

type Server struct {
	gs  *grpc.Server
	lis net.Listener
}

func NewHandler() *S {
	return &S{}
}

// comment just for testing shit
func (s *S) Handshake(_ context.Context, req *tester.HandshakeReq) (*tester.HandshakeResp, error) {
	fmt.Printf("Incoming request, message - %s\n", req.GetHelloOut())
	resp := &tester.HandshakeResp{
		HelloBack: "from server",
	}
	return resp, nil
}

func newServer(addr string) (server *Server, err error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("listening on %s\n", addr)

	grpcServer := grpc.NewServer()
	h := NewHandler()
	tester.RegisterHandshakerServer(grpcServer, h)
	return &Server{
		gs:  grpcServer,
		lis: lis,
	}, nil
}

func (srv *Server) Start() error {
	return srv.gs.Serve(srv.lis)
}

func main() {
	var addr = flag.String("serveraddr", ":5300", "where to bang to")
	flag.Parse()
	s, err := newServer(*addr)
	if err != nil {
		log.Fatalf("Couldnt make new grpc server, error - %v\n", err)
	}

	err = s.Start()
	if err != nil {
		log.Fatalf("Couldnt start grpc server, error - %v\n", err)
	}
}
