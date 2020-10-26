package user

import (
	"xs.bbs/internal/app/user/dao"
	"xs.bbs/internal/app/user/service"

	"github.com/google/wire"
)

var Set = wire.NewSet(dao.UserDaoSet, service.UserServiceSet)
