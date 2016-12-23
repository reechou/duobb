package models

import (
	"errors"
)

var (
	DB_ERROR                  = errors.New("数据库错误")
	CREATE_ACCOUNT_ERROR_ARGV = errors.New("创建账户参数错误")
	UPDATE_ACCOUNT_ERROR_ARGV = errors.New("创建账户参数错误")
	CREATE_SP_PLAN_ERROR_ARGV = errors.New("创建选品计划参数错误")
)
