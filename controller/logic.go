package controller

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/reechou/duobb/config"
	"github.com/reechou/duobb/models"
	"github.com/reechou/duobb/pkg/signal"
)

type Logic struct {
	cfg     *config.Config
	httpSrv *DuobbHttpServer
	tcpSrv  *DuobbTcpServer
}

func NewLogic(cfg *config.Config) *Logic {
	if cfg.Debug {
		EnableDebug()
	}

	l := &Logic{
		cfg:     cfg,
		httpSrv: NewDuobbHttpServer(cfg),
		tcpSrv:  NewDuobbTcpServer(cfg),
	}
	models.InitDB(cfg)

	return l
}

func (self *Logic) Run() {
	serveWait := make(chan error)

	go self.httpSrv.Start(serveWait)
	//self.tcpSrv.Start()

	signal.Trap(func() {
		//self.tcpSrv.Stop()
		self.httpSrv.Stop()
		<-serveWait
	})

	errServer := <-serveWait
	if errServer != nil {
		logrus.Fatalf("shutting down due to serve error: %v", errServer)
	}
	logrus.Infof("shutdown server success.")
}

func EnableDebug() {
	os.Setenv("DEBUG", "1")
	logrus.SetLevel(logrus.DebugLevel)
}
