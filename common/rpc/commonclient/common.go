// Code generated by goctl. DO NOT EDIT!
// Source: common.proto

//go:generate mockgen -destination ./common_mock.go -package commonclient -source $GOFILE

package commonclient

import (
	"context"

	"datacenter/common/rpc/common"

	"github.com/zeromicro/go-zero/zrpc"
)

type (
	BaseAppReq    = common.BaseAppReq
	BaseAppResp   = common.BaseAppResp
	AppConfigReq  = common.AppConfigReq
	AppConfigResp = common.AppConfigResp

	Common interface {
		GetAppConfig(ctx context.Context, in *AppConfigReq) (*AppConfigResp, error)
		GetBaseApp(ctx context.Context, in *BaseAppReq) (*BaseAppResp, error)
	}

	defaultCommon struct {
		cli zrpc.Client
	}
)

func NewCommon(cli zrpc.Client) Common {
	return &defaultCommon{
		cli: cli,
	}
}

func (m *defaultCommon) GetAppConfig(ctx context.Context, in *AppConfigReq) (*AppConfigResp, error) {
	client := common.NewCommonClient(m.cli.Conn())
	return client.GetAppConfig(ctx, in)
}

func (m *defaultCommon) GetBaseApp(ctx context.Context, in *BaseAppReq) (*BaseAppResp, error) {
	client := common.NewCommonClient(m.cli.Conn())
	return client.GetBaseApp(ctx, in)
}
