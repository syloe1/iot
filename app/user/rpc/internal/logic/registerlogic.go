package logic

import (
	"context"
	"errors"
	"strings"

	"iot-platform/app/user/rpc/internal/svc"
	"iot-platform/app/user/rpc/model"
	"iot-platform/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	username := strings.TrimSpace(in.Username)
	password := strings.TrimSpace(in.Password)
	if username == "" {
		return nil, errors.New("username is required")
	}
	if password == "" {
		return nil, errors.New("password is required")
	}

	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, username)
	if err == nil {
		return nil, errors.New("username already exists")
	}
	if !errors.Is(err, model.ErrNotFound) {
		return nil, err
	}

	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	userId, err := ret.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{
		Success: true,
		UserId:  userId,
		Msg:     "register success",
	}, nil
}
