package routes

//路由
import (
	"net/http"
	"xxx/controller"
	"xxx/logger"
	"xxx/middlewares"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "xxx/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置成发布模式
	}
	r := gin.New()
	//r.Use(middlewares.Cors())                           //
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) //, middlewares.RateLimitMiddleware(2*time.Second, 1)) //两秒钟添加一个，总量为1

	r.LoadHTMLFiles("./template/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/it", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "打怪欧",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	v1 := r.Group("/api/v1")
	//注册
	v1.POST("/signup", controller.SignUpHandler)
	//登录
	v1.POST("/login", controller.LoginHandler)

	v1.GET("/community", controller.CommunityHandler)           //通过token查看社区所有的id和name
	v1.GET("/community/:id", controller.CommunityDetailHandler) //通过id查看所有内容
	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.GET("/posts/", controller.GetPostListHandler)
	//更具时间或分数获取帖子列表
	v1.GET("/post2/", controller.GetPostListHandler2)
	//应用JWT认证中间件
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.POST("/post", controller.CreatePostHandler) //写帖子
		//投票
		v1.POST("/vote", controller.PostVoteController)
	}
	pprof.Register(r) //注册pprof相关路由

	r.NoRoute(func(ctx *gin.Context) { //无路由输出
		ctx.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})

	return r
}
