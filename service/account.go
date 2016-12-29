package service

import (
	"math/rand"
	"net/http"
	"time"

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
		logrus.Errorf("get account error: %v", err)
		return err
	}
	if account.Password == req.Password {
		rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	} else {
		rsp.Code = duobb_proto.DUOBB_MSG_LOGIN_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_LOGIN_ERROR
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
	rsp.Data = account.ID

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

func (self *AccountService) GetDuobbAccountFromPhone(r *http.Request, req *duobb_proto.GetDuobbAccountFromPhoneReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("GetDuobbAccountFromPhone req: %v", req)
	account := &models.DuobbAccount{
		Phone: req.Phone,
	}
	err := models.GetDuobbAccountFromPhone(account)
	if err != nil {
		logrus.Errorf("get duobb account from phone error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	rsp.Data = account

	return nil
}

func (self *AccountService) AccountUploadData(r *http.Request, req *duobb_proto.DuobbAccountUploadDataReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("AccountUploadData req: %v", req)
	info := &models.DuobbAccountCommission{
		UserName:          req.User,
		AlimamaName:       req.AlimamaName,
		Day:               req.Day,
		TodaySendItemsNum: req.TodaySendItemsNum,
		TodayBuyItemsNum:  req.TodayBuyItemsNum,
		TodayCommission:   req.TodayCommission,
	}
	affected, err := models.UpdateDuobbAccountCommission(info)
	if err != nil {
		logrus.Errorf("update duobb account commission error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	if affected == 0 {
		err := models.CreateDuobbAccountCommission(info)
		if err != nil {
			logrus.Errorf("create duobb account commission error: %v", err)
			rsp.Code = duobb_proto.DUOBB_DB_ERROR
			rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
			return err
		}
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}

func (self *AccountService) GetAllDuobbData(r *http.Request, req *duobb_proto.GetAllDuobbDataReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("GetAllDuobbData req: %v", req)
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	type AllDuobbData struct {
		LoginUser       int   `json:"loginUser"`
		TodayCommission int64 `json:"todayCommission"`
		AllCommission   int64 `json:"allCommission"`
	}
	data := &AllDuobbData{}
	data.LoginUser = 1975 + ra.Intn(1357)

	startTime := 1475251200
	now := time.Now()
	hour := now.Hour()
	data.TodayCommission = int64(hour*1253) + (now.Unix() % 60 * 7)
	data.AllCommission = (now.Unix()-int64(startTime))/86400*23157 + data.TodayCommission

	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	rsp.Data = data

	return nil
}

func (self *AccountService) AccountUploadAC(r *http.Request, req *duobb_proto.DuobbAccountUploadAlimamaCookieReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("AccountUploadAC req: %v", req)
	info := &models.DuobbAccountCookie{
		UserName:    req.User,
		AlimamaName: req.AlimamaName,
		Cookie:      req.Cookie,
	}
	affected, err := models.UpdateDuobbAccountCookie(info)
	if err != nil {
		logrus.Errorf("update duobb account cookie error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	if affected == 0 {
		err := models.CreateDuobbAccountCookie(info)
		if err != nil {
			logrus.Errorf("create duobb account cookie error: %v", err)
			rsp.Code = duobb_proto.DUOBB_DB_ERROR
			rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
			return err
		}
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	
	return nil
}

func (self *AccountService) GetAccountAC(r *http.Request, req *duobb_proto.DuobbAccountGetAlimamaCookieReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("GetAccountAC req: %v", req)
	info := &models.DuobbAccountCookie{
		UserName:    req.User,
		AlimamaName: req.AlimamaName,
	}
	err := models.GetDuobbAccountCookie(info)
	if err != nil {
		logrus.Errorf("get duobb account cookie error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	rsp.Data = info
	
	return nil
}
