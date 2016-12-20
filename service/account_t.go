package service

import (
	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/reechou/duobb/models"
	"github.com/reechou/duobb_proto"
)

type TAccountService struct{}

func (self *TAccountService) CreateDuobbAccount(req *duobb_proto.CreateDuobbAccountReq) error {
	account := &models.DuobbAccount{
		UserName: req.User,
		Password: req.Password,
		Phone:    req.Phone,
	}
	err := models.CreateDuobbAccount(account)
	if err != nil {
		logrus.Errorf("create duobb account error: %v", err)
		return errors.New(duobb_proto.MSG_DUOBB_DB_ERROR)
	}

	return nil
}

func (self *TAccountService) GetDuobbAccount(req *duobb_proto.GetDuobbAccountReq) (*models.DuobbAccount, error) {
	account := &models.DuobbAccount{
		UserName: req.User,
	}
	err := models.GetDuobbAccount(account)
	if err != nil {
		logrus.Errorf("create duobb account error: %v", err)
		return nil, errors.New(duobb_proto.MSG_DUOBB_DB_ERROR)
	}

	return account, nil
}
