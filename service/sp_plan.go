package service

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/reechou/duobb/models"
	"github.com/reechou/duobb_proto"
)

type SpPlanService struct{}

func (self *SpPlanService) CreateSpPlan(r *http.Request, req *duobb_proto.CreateSpPlanReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("CreateSpPlan req: %v", req)
	count, err := models.GetSpPlanCountFromUser(req.User)
	if err != nil {
		logrus.Errorf("get duobb sp plan count error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	if count >= SP_PLAN_MAX_COUNT {
		rsp.Code = duobb_proto.DUOBB_CREATE_PLAN_OVER_LIMIT_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_CREATE_PLAN_OVER_LIMIT
		return nil
	}
	plan := &models.SpPlan{
		Name:       req.Name,
		CreateUser: req.User,
		Password:   req.Password,
		Remark:     req.Remark,
	}
	err = models.CreateSpPlan(plan)
	if err != nil {
		logrus.Errorf("create duobb sp plan error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	rsp.Data = plan

	return nil
}

func (self *SpPlanService) DeleteSpPlan(r *http.Request, req *duobb_proto.DeleteSpPlanReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("CreateSpPlan req: %v", req)
	plan := &models.SpPlan{
		Id:         req.PlanId,
		CreateUser: req.User,
	}
	err := models.DeleteSpPlan(plan)
	if err != nil {
		logrus.Errorf("delete duobb sp plan error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}

func (self *SpPlanService) GetSpPlanListFromUser(r *http.Request, req *duobb_proto.GetSpPlanListFromUserReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("GetSpPlanListFromUser req: %v", req)
	list, err := models.GetSpPlanListFromUser(req.User, req.Offset, req.Num)
	if err != nil {
		logrus.Errorf("get duobb sp plan list from user error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	rsp.Data = list

	return nil
}

func (self *SpPlanService) GetSpPlanListPublic(r *http.Request, req *duobb_proto.GetSpPlanListPublicReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("GetSpPlanListPublic req: %v", req)
	list, err := models.GetSpPlanListPublic(req.QueryPriceStart, req.QueryPriceEnd, req.QueryCommissionStart, req.QueryCommissionEnd, req.QueryNumStart, req.QueryNumEnd, req.Offset, req.Num)
	if err != nil {
		logrus.Errorf("get duobb sp plan list public error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	rsp.Data = list

	return nil
}

func (self *SpPlanService) GetSpPlanInfoFromUser(r *http.Request, req *duobb_proto.GetSpPlanInfoFromUserReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("GetSpPlanInfoFromUser req: %v", req)
	plan := &models.SpPlan{
		Id:         req.PlanId,
		CreateUser: req.User,
	}
	err := models.GetSpPlanInfoFromUser(plan)
	if err != nil {
		logrus.Errorf("get duobb sp plan info from user error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	rsp.Data = plan
	//logrus.Debug(plan)

	return nil
}

func (self *SpPlanService) GetSpPlanInfoFromPassword(r *http.Request, req *duobb_proto.GetSpPlanFromPasswordReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("GetSpPlanInfoFromPassword req: %v", req)
	plan := &models.SpPlan{
		Id:       req.PlanId,
		Password: req.Password,
	}
	err := models.GetSpPlanInfoFromPassword(plan)
	if err != nil {
		logrus.Errorf("get duobb sp plan info from password error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS
	rsp.Data = plan

	return nil
}

func (self *SpPlanService) UpdateSpPlanInfo(r *http.Request, req *duobb_proto.UpdateSpPlanInfoReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("UpdateSpPlanInfo req: %v", req)
	plan := &models.SpPlan{
		Id:         req.PlanId,
		CreateUser: req.User,
		Name:       req.Name,
		Password:   req.Password,
		Remark:     req.Remark,
	}
	err := models.UpdateSpPlanInfo(plan)
	if err != nil {
		logrus.Errorf("update duobb sp plan info error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}

func (self *SpPlanService) UpdateSpPlanItems(r *http.Request, req *duobb_proto.UpdateSpPlanItemsReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("UpdateSpPlanItems req planid[%d] itemsnum: %d", req.PlanId, req.ItemsNum)
	plan := &models.SpPlan{
		Id:            req.PlanId,
		CreateUser:    req.User,
		ItemsNum:      req.ItemsNum,
		ItemsAvgPrice: req.ItemsAvgPrice,
		AvgCommission: req.AvgCommission,
		ItemsList:     req.ItemsList,
	}
	err := models.UpdateSpPlanItems(plan)
	if err != nil {
		logrus.Errorf("update duobb sp plan items error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}

func (self *SpPlanService) UpdateSpPlanPassword(r *http.Request, req *duobb_proto.UpdateSpPlanPasswordReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("UpdateSpPlanPassword req: %v", req)
	plan := &models.SpPlan{
		Id:         req.PlanId,
		CreateUser: req.User,
		Password:   req.Password,
	}
	err := models.UpdateSpPlanPassword(plan)
	if err != nil {
		logrus.Errorf("update duobb sp plan password error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}

func (self *SpPlanService) UpdateSpPlanRemark(r *http.Request, req *duobb_proto.UpdateSpPlanRemarkReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("UpdateSpPlanRemark req: %v", req)
	plan := &models.SpPlan{
		Id:         req.PlanId,
		CreateUser: req.User,
		Remark:     req.Remark,
	}
	err := models.UpdateSpPlanRemark(plan)
	if err != nil {
		logrus.Errorf("update duobb sp plan remark error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}

func (self *SpPlanService) SyncSpPlanSource(r *http.Request, req *duobb_proto.SyncSpPlanSourceReq, rsp *duobb_proto.Response) error {
	logrus.Debugf("SyncSpPlanSource req: %v", req)
	plan := &models.SpPlan{
		Id:       req.SourceFromId,
		Password: req.SourceIdPassword,
	}
	err := models.GetSpPlanInfoFromPassword(plan)
	if err != nil {
		logrus.Errorf("get duobb sp plan from password error: %v", err)
		rsp.Code = duobb_proto.DUOBB_GET_PLAN_FROM_PW_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_PLAN_FROM_PW_ERROR
		return err
	}
	plan = &models.SpPlan{
		Id:           req.PlanId,
		CreateUser:   req.User,
		SourceFromId: req.SourceFromId,
	}
	err = models.UpdateSpPlanSourceFrom(plan)
	if err != nil {
		logrus.Errorf("update duobb sp plan source error: %v", err)
		rsp.Code = duobb_proto.DUOBB_DB_ERROR
		rsp.Msg = duobb_proto.MSG_DUOBB_DB_ERROR
		return err
	}
	rsp.Code = duobb_proto.DUOBB_RSP_SUCCESS

	return nil
}
