package user

import (
	"context"
	"errors"

	"go-zore/user-api/internal/svc"
	"go-zore/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {

	// m := map[int64]string{
	// 	1: "张三",
	// 	2: "李四",
	// }

	// nickname := "unknow"
	// if name, ok := m[req.UserId]; ok {
	// 	nickname = name
	// }

	// return &types.UserInfoResponse{
	// 	UserId:   req.UserId,
	// 	Nickname: nickname,
	// }, nil
	logx.Error("user info ing")
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		return nil, errors.New("查询数据失败")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	return &types.UserInfoResponse{
		UserId:   user.Id,
		Nickname: user.Nickname,
	}, nil
}
