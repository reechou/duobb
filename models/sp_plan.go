package models

import (
	"fmt"
	"time"
	
	"github.com/Sirupsen/logrus"
)

type SpPlan struct {
	Id            int64   `xorm:"pk autoincr"`
	Name          string  `xorm:"not null default '' varchar(128)"`
	CreateUser    string  `xorm:"not null default '' varchar(128) index"`
	Password      string  `xorm:"not null default '' varchar(64) index"`
	SourceFromId  int64   `xorm:"not null default 0 int"`
	ItemsNum      int64   `xorm:"not null default 0 int"`
	ItemsAvgPrice float32 `xorm:"not null default 0.00 float(9,2)"`
	AvgCommission float32 `xorm:"not null default 0.00 float(9,2)"`
	ItemsList     string  `xorm:"not null mediumtext" json:"ItemsList,omitempty"`
	Remark        string  `xorm:"not null default '' varchar(128)"`
	CreatedAt     int64   `xorm:"not null default 0 int" json:"CreatedAt,omitempty"`
	UpdatedAt     int64   `xorm:"not null default 0 int" json:"UpdatedAt,omitempty"`
}

func CreateSpPlan(info *SpPlan) error {
	if info.Name == "" {
		return CREATE_SP_PLAN_ERROR_ARGV
	}
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now
	_, err := x.Insert(info)
	if err != nil {
		logrus.Errorf("create duobb sp plan error: %v", err)
		return DB_ERROR
	}
	logrus.Infof("create duobb sp plan[%s] success.", info.Name)
	
	return nil
}

func GetSpPlanListFromUser(user string, offset, num int64) ([]SpPlan, error) {
	var spPlanList []SpPlan
	err := x.Cols("id", "name", "create_user", "password", "items_num", "items_avg_price", "avg_commission", "remark", "created_at").Where("create_user = ?", user).Limit(int(num), int(offset)).Find(&spPlanList)
	if err != nil {
		logrus.Errorf("get sp plan list error: %v", err)
		return nil, err
	}
	return spPlanList, nil
}

func GetSpPlanListPublic(offset, num int64) ([]SpPlan, error) {
	var spPlanList []SpPlan
	err := x.Cols("id", "name", "create_user", "password", "items_num", "items_avg_price", "avg_commission", "remark", "created_at").Where("password = ''").Limit(int(num), int(offset)).Find(&spPlanList)
	if err != nil {
		logrus.Errorf("get sp plan list public error: %v", err)
		return nil, err
	}
	return spPlanList, nil
}

func GetSpPlanInfoFromUser(info *SpPlan) error {
	has, err := x.Where("id = ?", info.Id).And("create_user = ?", info.CreateUser).Get(info)
	if err != nil {
		logrus.Errorf("get duobb sp plan[%d][%s] error: %v", info.Id, info.CreateUser, err)
		return err
	}
	if !has {
		logrus.Errorf("cannot found sp paln[%d] with user[%s]", info.Id, info.CreateUser)
		return fmt.Errorf("cannot found sp paln[%d] with user[%s]", info.Id, info.CreateUser)
	}
	return nil
}

func GetSpPlanInfoFromPassword(info *SpPlan) error {
	has, err := x.Where("id = ?", info.Id).And("password = ?", info.Password).Get(info)
	if err != nil {
		logrus.Errorf("get duobb sp plan[%d][%s] error: %v", info.Id, info.Password, err)
		return err
	}
	if !has {
		logrus.Errorf("cannot found sp paln[%d] with password[%s]", info.Id, info.Password)
		return fmt.Errorf("cannot found sp paln[%d] with password[%s]", info.Id, info.Password)
	}
	return nil
}

func UpdateSpPlanItems(info *SpPlan) error {
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Id(info.Id).Cols("items_num", "items_avg_price", "avg_commission", "items_list", "updated_at").Update(info, &SpPlan{CreateUser: info.CreateUser})
	if err != nil {
		logrus.Errorf("update duobb sp plan items error: %v", err)
		return DB_ERROR
	}
	
	return nil
}

func UpdateSpPlanPassword(info *SpPlan) error {
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Id(info.Id).Cols("password", "updated_at").Update(info, &SpPlan{CreateUser: info.CreateUser})
	if err != nil {
		logrus.Errorf("update duobb sp plan password error: %v", err)
		return DB_ERROR
	}
	
	return nil
}

func UpdateSpPlanRemark(info *SpPlan) error {
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Id(info.Id).Cols("remark", "updated_at").Update(info, &SpPlan{CreateUser: info.CreateUser})
	if err != nil {
		logrus.Errorf("update duobb sp plan remark error: %v", err)
		return DB_ERROR
	}
	
	return nil
}

func UpdateSpPlanSourceFrom(info *SpPlan) error {
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Id(info.Id).Cols("source_from_id", "updated_at").Update(info, &SpPlan{CreateUser: info.CreateUser})
	if err != nil {
		logrus.Errorf("update duobb sp plan source from id error: %v", err)
		return DB_ERROR
	}
	
	return nil
}
