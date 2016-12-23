package controller

import (
	"net"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/reechou/duobb/config"
	"github.com/reechou/duobb/service"
	"github.com/reechou/duobb_proto"
)

type DuobbHttpServer struct {
	cfg    *config.Config
	server *rpc.Server
	l      net.Listener
}

func NewDuobbHttpServer(cfg *config.Config) *DuobbHttpServer {
	dhs := &DuobbHttpServer{
		cfg: cfg,
	}

	return dhs
}

func (self *DuobbHttpServer) Start(waitChan chan error) {
	self.server = rpc.NewServer()
	self.server.RegisterCodec(json.NewCodec(), "application/json")
	self.server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	self.server.RegisterService(new(service.AccountService), duobb_proto.AccountService)
	self.server.RegisterService(new(service.SpPlanService), duobb_proto.SPPlanService)

	r := mux.NewRouter()
	r.Handle("/rpc", self.server)

	l, err := net.Listen("tcp", self.cfg.HttpHost)
	if err != nil {
		logrus.Fatal(err)
	}
	self.l = l

	logrus.Infof("HTTP server listen on %s", self.cfg.HttpHost)
	err = http.Serve(self.l, r)
	if err != nil && strings.Contains(err.Error(), "use of closed network connection") {
		err = nil
	}
	waitChan <- err
}

func (self *DuobbHttpServer) Stop() {
	err := self.l.Close()
	if err != nil {
		logrus.Error(err)
	}
}
