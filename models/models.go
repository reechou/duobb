package models

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/reechou/duobb/config"
)

var x *xorm.Engine

func InitDB(cfg *config.Config) {
	var err error
	x, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		cfg.AccountDBInfo.User,
		cfg.AccountDBInfo.Pass,
		cfg.AccountDBInfo.Host,
		cfg.AccountDBInfo.DBName))
	if err != nil {
		logrus.Fatalf("Fail to init new engine: %v", err)
	}
	//x.SetLogger(nil)
	x.SetMapper(core.GonicMapper{})
	x.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	x.ShowSQL(true)

	if err = x.Sync2(new(DuobbAccount), new(SpPlan), new(DuobbAccountCommission), new(DuobbAccountCookie)); err != nil {
		logrus.Fatalf("Fail to sync database: %v", err)
	}
}
