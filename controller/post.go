package controller

import (
	"strconv"
	"xxx/logic"
	"xxx/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	Page = 1
	Size = 10
)

func CreatePostHandler(ctx *gin.Context) {
	//1.获取参数及参数的校验
	//ctx.ShouldBindJSON() //validator --> binding tag
	p := new(models.Post)

	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Debug("Ctx.shouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	//从c取到当前发请求的用户的id
	userID, err := GetCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(ctx, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(ctx *gin.Context) {
	//1.获取参数(从URL中获取帖子的id)
	pidStr := ctx.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
	}
	//2.根据id取出帖子数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPost(pid) failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(ctx, data)
}

// GetPostListHandler 获取帖子列表接口
func GetPostListHandler(ctx *gin.Context) {
	page, size := getPageInfo(ctx)
	//1.获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//2.返回响应
	ResponseSuccess(ctx, data)
}

// GetPostListHandler2 升级版帖子列表接口
//根据前端传入的参数动态的获取帖子列表
//按创建时间排序，或者按照分数排序
//1.获取请求的query string参数
//2.去redis查询id列表
//3.根据id去数据库查询帖子详细信息

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /post2 [get]
func GetPostListHandler2(ctx *gin.Context) {
	//GET请求参数(query string):/api/v1/post2?page=_&size=_&order=time(score)
	//1.获取参数
	//p := new(models.ParamPostList)
	p := &models.ParamPostList{ //初始化结构体时，指定初始参数
		Page:  Page,
		Size:  Size,
		Order: models.OrderTime,
	}

	if err := ctx.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostListNew(p) // 更新:合二为一
	//ctx.shouldBind() 根据请求的数据类型选择相应的方法获取数据
	//ctx.shouldBindJSON() 如果请求中携带的是json格式的数据，才能用此法获取到数据
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//2.返回响应
	ResponseSuccess(ctx, data)
}

//根据社区去查询帖子列表
/*
func GetCommunityPostListHandler(ctx *gin.Context) {
	//初始化结构体时指定初始参数
	p := &models.ParamPostList{
		ParamPostList:*models.ParamPostList{
			Page:  1,
			Size:  10,
			Order: models.OrderTime,
		}
	}

	if err := ctx.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	if p.CommunityID ==0{

	}

	//获取参数
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//2.返回响应
	ResponseSuccess(ctx, data)

}*/
