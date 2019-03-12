package service

import (
	"context"
	"fmt"
	"github.com/godcong/role-manager-server/config"
	"github.com/godcong/role-manager-server/proto"
	"github.com/json-iterator/go"
	"github.com/micro/go-micro"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"syscall"
	"time"
)

// GRPCServer ...
type GRPCServer struct {
	config  *config.Configure
	service micro.Service
	Type    string
	Port    string
	Path    string
}

func (s *GRPCServer) NodeBack(ctx context.Context, req *proto.ManagerNodeRequest, res *proto.ManagerReply) error {
	var nc NodeCallback
	var err error
	log.Infof("%+v", req.Detail)
	err = jsoniter.UnmarshalFromString(req.Detail, &nc)
	if err != nil {
		return err
	}
	log.Infof("%+v", nc)
	err = NodeCallbackProcess(nc.ID, &nc)
	if err != nil {
		return err
	}
	res = Result(&proto.ManagerReplyDetail{
		ID:   req.ID,
		Json: "",
	})
	return nil
}

func (s *GRPCServer) CensorBack(ctx context.Context, req *proto.ManagerCensorRequest, res *proto.ManagerReply) error {
	var cc CensorCallback
	var err error

	err = jsoniter.UnmarshalFromString(req.Detail, &cc)
	if err != nil {
		return err
	}

	err = CensorCallbackProcess(cc.ID, cc.Detail)
	if err != nil {
		return err
	}
	res = Result(&proto.ManagerReplyDetail{
		ID:   req.ID,
		Json: "",
	})
	return nil
}

// NodeBack ...
//func (s *GRPCServer) NodeBack(ctx context.Context, req *proto.ManagerNodeRequest) (*proto.ManagerReply, error) {
//	var nc NodeCallback
//	var err error
//	log.Infof("%+v", req.Detail)
//	err = jsoniter.UnmarshalFromString(req.Detail, &nc)
//	if err != nil {
//		return nil, err
//	}
//	log.Infof("%+v", nc)
//	err = NodeCallbackProcess(nc.ID, &nc)
//	if err != nil {
//		return nil, err
//	}
//	return Result(&proto.ManagerReplyDetail{
//		ID:   req.ID,
//		Json: "",
//	}), nil
//}

// CensorBack ...
//func (s *GRPCServer) CensorBack(ctx context.Context, req *proto.ManagerCensorRequest) (*proto.ManagerReply, error) {
//	var cc CensorCallback
//	var err error
//
//	err = jsoniter.UnmarshalFromString(req.Detail, &cc)
//	if err != nil {
//		return nil, err
//	}
//
//	err = CensorCallbackProcess(cc.ID, cc.Detail)
//	if err != nil {
//		return nil, err
//	}
//	return Result(&proto.ManagerReplyDetail{
//		ID:   req.ID,
//		Json: "",
//	}), nil
//}

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
		Type:   config.DefaultString("", Type),
		Port:   config.DefaultString("", ":7781"),
		Path:   config.DefaultString("", "/tmp/manager.sock"),
	}
}

// GRPCClient ...
type GRPCClient struct {
	config  *config.Configure
	service micro.Service
	Type    string
	Port    string
	Addr    string
}

// Conn ...
func (c *GRPCClient) Conn() (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error

	if c.Type == "unix" {
		conn, err = grpc.Dial("passthrough:///unix://"+c.Addr, grpc.WithInsecure())
	} else {
		conn, err = grpc.Dial(c.Addr+c.Port, grpc.WithInsecure())
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
func NewNodeGRPC(cfg *config.Configure) proto.NodeService {
	client := &GRPCClient{
		service: micro.NewService(),
		config:  cfg,
		Type:    config.DefaultString(cfg.Node.Type, Type),
		Port:    config.DefaultString(cfg.Node.Port, ":7787"),
		Addr:    config.DefaultString(cfg.Node.Addr, "/tmp/node.sock"),
	}
	client.service.Init()
	return proto.NewNodeService("go.micro.grpc.node", client.service.Client())
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
	var lis net.Listener
	var port string
	var err error
	s.service = micro.NewService(
		micro.Name("manager"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Version("latest"),
	)
	s.service.Init()
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

		_ = proto.RegisterManagerServiceHandler(s.service.Server(), s)

		log.Printf("Listening and serving TCP on %s\n", port)
		if err := s.service.Run(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}

// Stop ...
func (s *GRPCServer) Stop() {
	//s.service.Stop()
}
