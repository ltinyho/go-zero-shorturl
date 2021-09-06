package logic

import (
	"context"
	"database/sql"
	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/model"
	"shorturl/rpc/transform/transform"

	"github.com/tal-tech/go-zero/core/hash"
	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenReq) (*transform.ShortenResp, error) {
	// 手动代码开始，生成短链接
	key := hash.Md5Hex([]byte(in.Url))[:6]
	_, err := l.svcCtx.Model.Insert(model.Shorturl{
		Shorten: sql.NullString{
			String: key,
			Valid:  true,
		},
		Url: sql.NullString{
			String: in.Url,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &transform.ShortenResp{
		Shorten: key,
	}, nil
	// 手动代码结束
}
