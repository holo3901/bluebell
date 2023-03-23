package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"xxx/controller"
	"xxx/dao/mysql"
	"xxx/dao/redis"
	"xxx/logger"
	"xxx/pkg/snowflake"
	"xxx/routes"
	"xxx/settings"

	"go.uber.org/zap"
)

//Go Web开发较通用否认脚手架模板

// @title 请求社区及投票数
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name zh
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	/*var x string
	flag.StringVar(&x, "name", "xxx", "姓名")
	flag.Parse()*/
	/*	if len(os.Args) < 2 { //如果不使用这个os,这需要修改setting里的Init，注销os获取文件目录
		fmt.Println("need config file")
		return
	}*/
	/*area := fmt.Sprintf(x)*/
	//1.加载配置
	if err := settings.Init( /*os.Args[1] area*/); err != nil {
		fmt.Printf("init setings failed,err:%v\n", err)
		return
	}

	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed failed,err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//3.初始化mysql连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
	}
	defer mysql.Close()
	//4.初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
	}
	defer redis.Close()
	if err := snowflake.Init(settings.Conf.AppConfig.StartTime, settings.Conf.AppConfig.MachineID); err != nil {
		fmt.Printf("init snowflake failed,err:%v\n", err)
		return
	}
	//初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator Trans failed,err:%v\n", err)
	}
	//5.注册路由
	r := routes.SetUp(settings.Conf.AppConfig.Mode)
	//6.启动服务(优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.AppConfig.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务,不然listenAndServe()会不断请求接收,循环.保证可以执行后面的代码
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
