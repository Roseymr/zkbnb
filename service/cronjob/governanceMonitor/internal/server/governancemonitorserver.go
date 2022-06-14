// Code generated by goctl. DO NOT EDIT!
// Source: governanceMonitor.proto

package server

import (
	"context"

	"github.com/zecrey-labs/zecrey-legend/service/cronjob/governanceMonitor/governanceMonitor"
	"github.com/zecrey-labs/zecrey-legend/service/cronjob/governanceMonitor/internal/logic"
	"github.com/zecrey-labs/zecrey-legend/service/cronjob/governanceMonitor/internal/svc"
)

type GovernanceMonitorServer struct {
	svcCtx *svc.ServiceContext
	governanceMonitor.UnimplementedGovernanceMonitorServer
}

func NewGovernanceMonitorServer(svcCtx *svc.ServiceContext) *GovernanceMonitorServer {
	return &GovernanceMonitorServer{
		svcCtx: svcCtx,
	}
}

func (s *GovernanceMonitorServer) Ping(ctx context.Context, in *governanceMonitor.Request) (*governanceMonitor.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}