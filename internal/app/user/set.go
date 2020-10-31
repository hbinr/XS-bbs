package user

import (
	"xs.bbs/internal/app/user/controller"
	"xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/app/user/model"
	"xs.bbs/internal/app/user/service"

	"github.com/google/wire"
)

var (
	Model = &model.User{}
	Set   = wire.NewSet(
		dao.UserDaoSet,
		service.UserServiceSet,
		controller.UserControllerSet,
	)
)
