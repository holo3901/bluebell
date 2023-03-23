package controller

import (
	"xxx/logic"
	"xxx/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//投票

func PostVoteController(ctx *gin.Context) {
	//参数校验
	p := new(models.ParamVoteData)
	if err := ctx.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言,获取错误自定义返回响应错误
		if !ok {
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) //transl翻译,removeTopStruct去掉翻译信息(错误提示)里结构体字段标识
		ResponseErrorWithMsg(ctx, CodeInvalidParam, errData)
		return
	}

	//获取当前请求的用户的id
	userID, err := GetCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeNeedLogin)
		return
	}
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, nil)
}
