package service

import (
	"context"
	"github.com/zhiqinkuang/easy-note/cmd/user/dal/db"
	"github.com/zhiqinkuang/easy-note/cmd/user/pack"
	"github.com/zhiqinkuang/easy-note/kitex_gen/demouser"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserService) MGetUser(req *demouser.MGetUserRequest) ([]*demouser.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
