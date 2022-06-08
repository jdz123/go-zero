package user

import (
	"context"

	"go-zore/user-api/internal/svc"
	"go-zore/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateRequest) (resp *types.UserUpdateResponse, err error) {
	// todo: add your logic here and delete this line
	if ern := 1.sveCtx.UserModel.TransCtx(l.ctx, func(etx context.Context, session sqlx.Session) Error { 
		user := &model.User{}
		user.Mobile = req.Mobile
		user.Nickname = req.Nickname
		//添加user
		dbResult,err := l.svcCtx.UserModel.Insert(ctx,session,user)
		if err !=nil{
			return err
		}
		userId,_:=dbResult.LastInsertId()
		//添加userData
		userData := &model.UserData{}
		userData.UserId = userId
		userData.Data = "xxx"
		if _,err := l.svcCtx.UserDataModel.Insert(ctx,session,userData);err!=nil{
			return err
		}
		return nil
	});err !=nil{
		return nil,errors.New("创建用户失败")
	}

	return
}
