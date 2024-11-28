package controllers

import (
	account "api-login-proto/account"
	"api-login-proto/common"
	"context"
	"service-account/models"
	"standard-library/consts"
	"standard-library/models/dto"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/v2/server/web"
)

type AccountController struct {
	account.UnimplementedUserAccountServiceServer
	BaseController
}

// Info
//
//	@Title			Account Info
//	@Description	Account Info Detail
//	@Success		200			{object}	web.M
//	@Param			AccountId	query		int64	false	"AccountID"
//	@Param			Username	query		string	false	"Account Username"
//	@Param			Email		query		string	false	"Account Email"
//	@Param			CreateTime	query		int64	false	"Account create time"
//	@Param			Page		query		int64	false	"Page"
//	@Param			PageSize	query		int64	false	"Page Size"
//	@router			/info [get]
func (ctl *AccountController) Info(ctx context.Context, request *common.Request) (*common.Response, error) {
	req := dto.ReqAccountDetail{}
	err := ctl.ParseJson(request, req)
	if err != nil {
		logs.Error("[AccountController][Info] Parse Json error", err)
		return ctl.Error(consts.PARAM_ERROR)
	}

	acc := models.Account{}
	data, errCode, err := acc.Info(req)
	if errCode != 0 || err != nil {
		logs.Error("[AccountController][Info] Get account info fail", err)
		if errCode == 0 {
			errCode = consts.OPERATION_FAILED
		}
		return ctl.Error(errCode)
	}

	return ctl.Success(web.M{
		"Item": data,
	})
}

// Info
//
//	@Title			Account Info
//	@Description	Account Info Detail
//	@Success		200			{object}	web.M
//	@Param			AccountId	query		int64	false	"AccountID"
//	@Param			Username	query		string	false	"Account Username"
//	@Param			Email		query		string	false	"Account Email"
//	@Param			CreateTime	query		int64	false	"Account create time"
//	@Param			Page		query		int64	false	"Page"
//	@Param			PageSize	query		int64	false	"Page Size"
//	@router			/info [get]
func (ctl *AccountController) List(ctx context.Context, request *common.Request) (*common.Response, error) {
	req := dto.ReqAccountList{}
	err := ctl.ParseJson(request, req)
	if err != nil {
		logs.Error("[AccountController][List] Parse Json error", err)
		return ctl.Error(consts.PARAM_ERROR)
	}

	acc := models.Account{}
	data, pagination, errCode, err := acc.List(req)
	if errCode != 0 || err != nil {
		logs.Error("[AccountController][List] Get account info fail", err)
		if errCode == 0 {
			errCode = consts.OPERATION_FAILED
		}
		return ctl.Error(errCode)
	}

	return ctl.Success(web.M{
		"Item":       data,
		"Pagination": pagination,
	})
}
