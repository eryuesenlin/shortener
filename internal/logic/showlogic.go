package logic

import (
	"context"
	"database/sql"
	"errors"

	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// 输入短链接，重定向到真实的链接
	// 1. 根据短链接查找原始的长链接
	u,err:=l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx,sql.NullString{Valid:true,String:req.ShortUrl})
	if err!=nil{
		if err==sql.ErrNoRows{
			return nil,errors.New("404")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed",logx.LogField{Value:err.Error(),Key:"err"})
		return nil,err
	}
	// 2. 返回重定向响应/返回查找到的长链接，在调用handler层返回重定向响应
	
	return &types.ShowResponse{LongUrl: u.Lurl.String},nil
}
