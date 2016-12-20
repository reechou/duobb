package service

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/reechou/duobb/models"
	"github.com/reechou/duobb_proto"
)

type AccountService struct{}

func (self *AccountService) LoginDuobbAccount(r *http.Request, req *duobb_proto.DuobbLogin, rsp *duobb_proto.Response) error {
	logrus.Debugf("LoginDuobbAccount req: %v", req)
	account := &models.DuobbAccount{
		UserName: req.User,
	}
	err := models.GetDuobbAccount(account)
	if err != nil {
		logrus.Errorf("create duobb account error: %v", err)
		return err
	}
	if account.Password == req.Password {
		rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	} else {
		rsp.Code = duobb_proto.DUOBB_MSG_LOGIN_ERROR
	}

	return nil
}

func (self *AccountService) CreateDuobbAccount(r *http.Request, req *duobb_proto.CreateDuobbAccountReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("CreateDuobbAccount req: %v", req)
	account := &models.DuobbAccount{
		UserName: req.User,
		Password: req.Password,
		Phone:    req.Phone,
		PicUrl:   req.PicUrl,
	}
	err := models.CreateDuobbAccount(account)
	if err != nil {
		logrus.Errorf("create duobb account error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}

func (self *AccountService) UpdateDuobbAccountPassword(r *http.Request, req *duobb_proto.UpdateDuobbAccountPasswordReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("UpdateDuobbAccountPassword req: %v", req)
	account := &models.DuobbAccount{
		UserName: req.User,
		Password: req.Password,
	}
	err := models.UpdateDuobbAccountPassword(account)
	if err != nil {
		logrus.Errorf("update duobb account password error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}

func (self *AccountService) UpdateDuobbAccountPhone(r *http.Request, req *duobb_proto.UpdateDuobbAccountPhoneReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("UpdateDuobbAccountPhone req: %v", req)
	account := &models.DuobbAccount{
		UserName: req.User,
		Phone:    req.Phone,
	}
	err := models.UpdateDuobbAccountPhone(account)
	if err != nil {
		logrus.Errorf("update duobb account phone error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}

func (self *AccountService) UpdateDuobbAccountPicUrl(r *http.Request, req *duobb_proto.UpdateDuobbAccountPicUrlReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("UpdateDuobbAccountPicUrl req: %v", req)
	account := &models.DuobbAccount{
		UserName: req.User,
		PicUrl:   req.PicUrl,
	}
	err := models.UpdateDuobbAccountPicUrl(account)
	if err != nil {
		logrus.Errorf("update duobb account pic url error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}

func (self *AccountService) GetDuobbAccount(r *http.Request, req *duobb_proto.GetDuobbAccountReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("GetDuobbAccount req: %v", req)
	account := &models.DuobbAccount{
		UserName: req.User,
	}
	err := models.GetDuobbAccount(account)
	if err != nil {
		logrus.Errorf("get duobb account error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	rsp.Data = account

	return nil
}
