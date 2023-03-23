package controller

import (
	"strconv"
	"xxx/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(ctx *gin.Context) {
	//查询到所有的社区(community_id,community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy) //不轻易把服务端报错暴露在外面
		return
	}
	ResponseSuccess(ctx, data)
}

//CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(ctx *gin.Context) {
	//1.获取社区id
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64) //"10"表示是十进制,bitSize表示是Int64
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy) //不轻易把服务端报错暴露在外面
		return
	}
	ResponseSuccess(ctx, data)

}
