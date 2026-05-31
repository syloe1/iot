package logic

import (
	"context"
	"errors"
	"time"

	"iot-platform/app/user/rpc/internal/svc"
	"iot-platform/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if u.Password != in.Password {
		return nil, errors.New("wrong password")
	}

	now := time.Now().Unix()
	token, err := l.svcCtx.JwtAuth.GenerateToken(now, u.Id)
	if err != nil {
		return nil, err
	}

	return &user.LoginResp{
		Token:  token,
		UserId: u.Id,
	}, nil
}
