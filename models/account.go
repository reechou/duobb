package models

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
)

type DuobbAccount struct {
	ID        int64  `xorm:"pk autoincr"`
	UserName  string `xorm:"not null default '' varchar(128) unique"`
	Password  string `xorm:"not null default '' varchar(128)" json:"-"`
	Phone     string `xorm:"not null default '' varchar(64) unique"`
	PicUrl    string `xorm:"not null default '' varchar(128)"`
	Level     int64  `xorm:"not null default 0 int"`
	Status    int64  `xorm:"not null default 0 int"`
	CreatedAt int64  `xorm:"not null default 0 int"`
	UpdatedAt int64  `xorm:"not null default 0 int"`
}

func GetDuobbAccount(info *DuobbAccount) error {
	has, err := x.Where("user_name = ?", info.UserName).Get(info)
	if err != nil {
		logrus.Errorf("get duobb account[%s] error: %v", info.UserName, err)
		return err
	}
	if !has {
		logrus.Errorf("cannot found account[%s]", info.UserName)
		return fmt.Errorf("cannot found account[%s]", info.UserName)
	}
	return nil
}

func CreateDuobbAccount(info *DuobbAccount) error {
	if info.UserName == "" || info.Password == "" || info.Phone == "" {
		return CREATE_ACCOUNT_ERROR_ARGV
	}
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now
	_, err := x.Insert(info)
	if err != nil {
		logrus.Errorf("create duobb account error: %v", err)
		return DB_ERROR
	}
	logrus.Infof("create duobb account[%v] success.", info)

	return nil
}

func UpdateDuobbAccountPassword(info *DuobbAccount) error {
	if info.UserName == "" || info.Password == "" {
		return UPDATE_ACCOUNT_ERROR_ARGV
	}
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Cols("password", "updated_at").Update(info, &DuobbAccount{UserName: info.UserName})
	if err != nil {
		logrus.Errorf("update duobb account error: %v", err)
		return DB_ERROR
	}

	return nil
}

func UpdateDuobbAccountPhone(info *DuobbAccount) error {
	if info.UserName == "" || info.Phone == "" {
		return UPDATE_ACCOUNT_ERROR_ARGV
	}
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Cols("phone", "updated_at").Update(info, &DuobbAccount{UserName: info.UserName})
	if err != nil {
		logrus.Errorf("update duobb account error: %v", err)
		return DB_ERROR
	}

	return nil
}

func UpdateDuobbAccountPicUrl(info *DuobbAccount) error {
	if info.UserName == "" {
		return UPDATE_ACCOUNT_ERROR_ARGV
	}
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Cols("pic_url", "updated_at").Update(info, &DuobbAccount{UserName: info.UserName})
	if err != nil {
		logrus.Errorf("update duobb account error: %v", err)
		return DB_ERROR
	}

	return nil
}

func UpdateDuobbAccountLevel(info *DuobbAccount) error {
	if info.UserName == "" || info.Password == "" {
		return UPDATE_ACCOUNT_ERROR_ARGV
	}
	now := time.Now().Unix()
	info.UpdatedAt = now
	_, err := x.Cols("level", "updated_at").Update(info, &DuobbAccount{UserName: info.UserName})
	if err != nil {
		logrus.Errorf("update duobb account error: %v", err)
		return DB_ERROR
	}

	return nil
}
