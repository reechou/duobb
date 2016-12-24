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

type DuobbAccountCommission struct {
	ID                int64   `xorm:"pk autoincr"`
	UserName          string  `xorm:"not null default '' varchar(128) unique(uni_user_day)"`
	Day               string  `xorm:"not null default '' varchar(16) unique(uni_user_day) index"`
	TodaySendItemsNum int64   `xorm:"not null default 0 int"`
	TodayBuyItemsNum  int64   `xorm:"not null default 0 int"`
	TodayCommission   float32 `xorm:"not null default 0.00 float(9,2)"`
	CreatedAt         int64   `xorm:"not null default 0 int"`
	UpdatedAt         int64   `xorm:"not null default 0 int"`
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

func GetDuobbAccountFromPhone(info *DuobbAccount) error {
	has, err := x.Where("phone = ?", info.Phone).Get(info)
	if err != nil {
		logrus.Errorf("get duobb account from phone[%s] error: %v", info.Phone, err)
		return err
	}
	if !has {
		logrus.Errorf("cannot found account from phone[%s]", info.Phone)
		return fmt.Errorf("cannot found account from phone[%s]", info.Phone)
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
	if info.UserName == "" {
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

func CreateDuobbAccountCommission(info *DuobbAccountCommission) error {
	if info.UserName == "" || info.Day == "" {
		return CREATE_ACCOUNT_COMMISSION_ERROR_ARGV
	}
	now := time.Now().Unix()
	info.CreatedAt = now
	info.UpdatedAt = now
	_, err := x.Insert(info)
	if err != nil {
		logrus.Errorf("create duobb account commission error: %v", err)
		return DB_ERROR
	}
	logrus.Infof("create duobb account commission[%v] success.", info)

	return nil
}

func UpdateDuobbAccountCommission(info *DuobbAccountCommission) (int64, error) {
	if info.UserName == "" || info.Day == "" {
		return 0, UPDATE_ACCOUNT_COMMISSION_ERROR_ARGV
	}
	now := time.Now().Unix()
	info.UpdatedAt = now
	affected, err := x.Cols("today_send_items_num", "today_buy_items_num", "today_commission", "updated_at").Update(info, &DuobbAccountCommission{UserName: info.UserName, Day: info.Day})
	if err != nil {
		logrus.Errorf("update duobb account today commission error: %v", err)
		return 0, DB_ERROR
	}

	return affected, nil
}

func GetDuobbAllCommission() (float64, error) {
	ss := new(DuobbAccountCommission)
	total, err := x.Where("id > ?", 0).Sum(ss, "today_commission")
	if err != nil {
		logrus.Errorf("duobb all commission sum error: %v", err)
		return 0.0, err
	}

	return total, nil
}

func GetDuobbAllCommissionByDay(day string) (float64, error) {
	ss := new(DuobbAccountCommission)
	total, err := x.Where("day = ?", day).Sum(ss, "today_commission")
	if err != nil {
		logrus.Errorf("duobb day commission sum error: %v", err)
		return 0.0, err
	}

	return total, nil
}