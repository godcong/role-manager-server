package service

import (
	"context"
	"github.com/godcong/role-manager-server/config"
	"github.com/godcong/role-manager-server/proto"
	"github.com/json-iterator/go"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	log "github.com/sirupsen/logrus"
	"time"
)

// GRPCServer ...
type GRPCServer struct {
	config  *config.Configure
	service micro.Service
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
	*res = Result(&proto.ManagerReplyDetail{
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
	*res = Result(&proto.ManagerReplyDetail{
		ID:   req.ID,
		Json: "",
	})
	return nil
}

// Result ...
func Result(detail *proto.ManagerReplyDetail) proto.ManagerReply {
	return proto.ManagerReply{
		Code:    0,
		Message: "success",
		Detail:  detail,
	}
}

// NewGRPCServer ...
func NewGRPCServer(cfg *config.Configure) *GRPCServer {
	return &GRPCServer{
		config: cfg,
	}
}

// GRPCClient ...
type GRPCClient struct {
	config  *config.Configure
	service micro.Service
}

// NewGRPCClient ...
func NewGRPCClient(cfg *config.Configure) *GRPCClient {
	reg := consul.NewRegistry()
	client := &GRPCClient{
		service: micro.NewService(
			micro.Registry(reg)),
		config: cfg,
	}
	client.service.Init()
	return client
}

// NodeClient ...
func NodeClient(g *GRPCClient) proto.NodeService {
	return proto.NewNodeService(g.config.Manager.NodeName, g.service.Client())
}

// ManagerClient ...
func ManagerClient(g *GRPCClient) proto.ManagerService {
	return proto.NewManagerService(g.config.Manager.ManagerName, g.service.Client())
}

// CensorClient ...
func CensorClient(g *GRPCClient) proto.CensorService {
	return proto.NewCensorService(g.config.Manager.CensorName, g.service.Client())
}

// Start ...
func (s *GRPCServer) Start() {
	if !s.config.Manager.EnableGRPC {
		return
	}
	reg := consul.NewRegistry()
	s.service = micro.NewService(
		micro.Name(s.config.Manager.ManagerName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(reg),
	)
	s.service.Init()
	go func() {
		_ = proto.RegisterManagerServiceHandler(s.service.Server(), s)
		if err := s.service.Run(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}

// Stop ...
func (s *GRPCServer) Stop() {
	if s.service != nil {
		s.service.Server().Stop()
	}
}
