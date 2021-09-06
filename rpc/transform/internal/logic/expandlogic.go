package logic

import (
	"context"
	"database/sql"

	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/transform"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandReq) (*transform.ExpandResp, error) {
	// 手动代码开始
	res, err := l.svcCtx.Model.FindOne(sql.NullString{
		String: in.Shorten,
		Valid:  true,
	})
	if err != nil {
		return nil, err
	}

	return &transform.ExpandResp{
		Url: res.Url.String,
	}, nil
	// 手动代码结束
}
