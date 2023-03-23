package controller

// 服务的入口，负责处理路由，参数校验，请求转发
import (
	"errors"
	"fmt"
	"xxx/dao/mysql"
	"xxx/logic"
	"xxx/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//SignUpHandler 处理注册请求函数
func SignUpHandler(ctx *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamSignUp)                   //分配了一个零初始化的T值，并返回指向它的指针
	if err := ctx.ShouldBindJSON(&p); err != nil { //让gin框架自动从请求中把想要的数据自动把他绑定到结构体里
		//请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err)) //记录日志
		//判断err是否为validator类型(无法进行翻译的类型)
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			/*1.ctx.JSON(http.StatusOK, gin.H{
				"msg": "请求参数有误" + err.Error(),
			})
			return*/
			//2.
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		/*1.ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})*/
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//手动对请求参数进行详细的业务规则校验
	/*if len(p.Username) == 0 || len(p.Password) == 0 || p.RePassword != p.Password {
		zap.L().Error("SignUp with invalid param,密码或用户为0，或者两次输入的密码不同")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}*/
	fmt.Println(p)
	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))
		/*ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})*/
		if errors.Is(err, mysql.ErrorUserExist) { //errors.Is判断两个err是否相同
			ResponseError(ctx, CodeUserExist)
			return
		}
		ResponseError(ctx, CodeServerBusy)
		return
	}
	//3.返回响应
	/*ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})*/
	ResponseSuccess(ctx, nil)
}
func LoginHandler(ctx *gin.Context) {
	//1.获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := ctx.ShouldBindJSON(&p); err != nil { //让gin框架自动从请求中把想要的数据自动把他绑定到结构体里
		//请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err)) //记录日志
		//判断err是否为validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			/*ctx.JSON(http.StatusOK, gin.H{
				"msg": "请求参数有误" + err.Error(),
			})*/
			ResponseError(ctx, CodeInvalidParam)
			return
		}
		/*ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})*/
		ResponseErrorWithMsg(ctx, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		/*ctx.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})*/
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(ctx, CodeUserNotExist)
		}
		ResponseError(ctx, CodeInvalidPassword)
		return
	}
	//3.返回响应
	/*ctx.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})*/
	ResponseSuccess(ctx, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), //id值大于1<<53-1 int64最大值是1<<64-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}
