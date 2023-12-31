package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zhanghongliang12/lebron/apps/user/rpc/internal/svc"
	"github.com/zhanghongliang12/lebron/apps/user/rpc/model"
	"github.com/zhanghongliang12/lebron/apps/user/rpc/user"
	"github.com/zhanghongliang12/lebron/pkg/tool"
	"github.com/zhanghongliang12/lebron/pkg/xerr"

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

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line

	userDB, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "根据username查询用户信息失败，username:%s,err:%v", in.Username, err)
		}
		return nil, err
	}
	//verify user password
	md5ByString, err := tool.Md5ByString(in.Password)
	fmt.Println(md5ByString)
	fmt.Println(userDB.Password)
	if !(md5ByString == userDB.Password) {
		return nil, errors.Wrap(xerr.NewErrMsg("账号或密码错误"), "密码错误")
	}

	//return sql
	var resp user.LoginResponse
	_ = copier.Copy(&resp, userDB)

	return &resp, nil
}
