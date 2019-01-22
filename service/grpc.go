package service

import (
	"context"
	"fmt"
	"github.com/godcong/role-manager-server/config"
	"github.com/godcong/role-manager-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"syscall"
)

// GRPCServer ...
type GRPCServer struct {
	config *config.Configure
	server *grpc.Server
	Type   string
	Port   string
	Path   string
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
func NewGRPCServer(cfg *config.Configure) *GRPCServer {
	return &GRPCServer{
		config: cfg,
		Type:   config.DefaultString("unix", Type),
		Port:   config.DefaultString("", ":7781"),
		Path:   config.DefaultString("", "/tmp/manager.sock"),
	}
}

// GRPCClient ...
type GRPCClient struct {
	config *config.Configure
	Type   string
	Port   string
	Addr   string
}

// Conn ...
func (c *GRPCClient) Conn() (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error

	if c.Type == "unix" {
		conn, err = grpc.Dial("passthrough:///unix://"+c.Addr, grpc.WithInsecure())
	} else {
		conn, err = grpc.Dial(c.Addr, grpc.WithInsecure())
	}

	return conn, err
}

// NodeClient ...
func NodeClient(g *GRPCClient) proto.NodeServiceClient {
	clientConn, err := g.Conn()
	if err != nil {
		log.Println(err)
		return nil
	}
	return proto.NewNodeServiceClient(clientConn)
}

// NewNodeGRPC ...
func NewNodeGRPC(cfg *config.Configure) *GRPCClient {
	return &GRPCClient{
		config: cfg,
		Type:   config.DefaultString(cfg.Node.Type, Type),
		Port:   config.DefaultString(cfg.Node.Port, ":7787"),
		Addr:   config.DefaultString(cfg.Node.Addr, "/tmp/node.sock"),
	}
}

// ManagerClient ...
func ManagerClient(g *GRPCClient) proto.ManagerServiceClient {
	clientConn, err := g.Conn()
	if err != nil {
		log.Println(err)
		return nil
	}
	return proto.NewManagerServiceClient(clientConn)
}

// NewManagerGRPC ...
func NewManagerGRPC(cfg *config.Configure) *GRPCClient {
	return &GRPCClient{
		config: cfg,
		Type:   config.DefaultString("unix", Type),
		Port:   config.DefaultString("", ":7781"),
		Addr:   config.DefaultString("", "/tmp/manager.sock"),
	}
}

// CensorClient ...
func CensorClient(g *GRPCClient) proto.CensorServiceClient {
	clientConn, err := g.Conn()
	if err != nil {
		log.Println(err)
		return nil
	}
	return proto.NewCensorServiceClient(clientConn)
}

// NewCensorGRPC ...
func NewCensorGRPC(cfg *config.Configure) *GRPCClient {
	return &GRPCClient{
		config: cfg,
		Type:   config.DefaultString(cfg.Censor.Type, Type),
		Port:   config.DefaultString(cfg.Censor.Port, ":7785"),
		Addr:   config.DefaultString(cfg.Censor.Addr, "/tmp/censor.sock"),
	}
}

// Start ...
func (s *GRPCServer) Start() {

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
