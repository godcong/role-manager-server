package service

import (
	"context"
	"fmt"
	"github.com/godcong/role-manager-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"syscall"
)

// GRPCServer ...
type GRPCServer struct {
	Type   string
	Port   string
	Path   string
	server *grpc.Server
}

// Back ...
func (s *GRPCServer) Back(ctx context.Context, r *proto.ManagerCallbackRequest) (*proto.ManagerReply, error) {
	log.Printf("%+v", r)
	return Result(&proto.ManagerReplyDetail{
		ID:   "",
		Json: "",
	}), nil
}

// Result ...
func Result(detail *proto.ManagerReplyDetail) *proto.ManagerReply {
	return &proto.ManagerReply{
		Code:    0,
		Message: "success",
		Detail:  detail,
	}
}

// NewGRPCServer ...
func NewGRPCServer() *GRPCServer {
	return &GRPCServer{
		Type: DefaultString("", Type),
		Port: DefaultString("", ":7784"),
		Path: DefaultString("", "/tmp/manager.sock"),
	}
}

// DefaultString ...
func DefaultString(v, s string) string {
	if v == "" {
		return s
	}
	return v
}

// Start ...
func (s *GRPCServer) Start() {
	//if !config.GRPC.Enable {
	//	return
	//}
	s.server = grpc.NewServer()
	var lis net.Listener
	var port string
	var err error
	go func() {
		if s.Type == "unix" {
			_ = syscall.Unlink(s.Path)
			lis, err = net.Listen(s.Type, s.Path)
			port = s.Path
		} else {
			lis, err = net.Listen("tcp", s.Port)
			port = s.Port
		}

		if err != nil {
			panic(fmt.Sprintf("failed to listen: %v", err))
		}

		proto.RegisterManagerServiceServer(s.server, s)
		// Register reflection service on gRPC server.
		reflection.Register(s.server)
		log.Printf("Listening and serving TCP on %s\n", port)
		if err := s.server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}

// Stop ...
func (s *GRPCServer) Stop() {
	s.server.Stop()
}
