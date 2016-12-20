package controller

import (
	"github.com/Sirupsen/logrus"
	"github.com/absolute8511/gorpc"
	"github.com/reechou/duobb/config"
	"github.com/reechou/duobb/service"
	"github.com/reechou/duobb_proto"
)

type DuobbTcpServer struct {
	cfg    *config.Config
	server *gorpc.Server
}

func NewDuobbTcpServer(cfg *config.Config) *DuobbTcpServer {
	dts := &DuobbTcpServer{
		cfg: cfg,
	}

	return dts
}

func (self *DuobbTcpServer) Start() {
	d := gorpc.NewDispatcher()
	d.AddService(duobb_proto.AccountService, new(service.TAccountService))

	self.server = gorpc.NewTCPServer(self.cfg.TcpHost, d.NewHandlerFunc())
	if err := self.server.Start(); err != nil {
		logrus.Fatalf("cannot start tcp rpc server: [%v]", err)
	}
	logrus.Infof("TCP server listen on %s", self.cfg.TcpHost)
}

func (self *DuobbTcpServer) Stop() {
	self.server.Stop()
}
