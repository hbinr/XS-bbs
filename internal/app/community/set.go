package community

import (
	"github.com/google/wire"
	"xs.bbs/internal/app/community/controller"
	"xs.bbs/internal/app/community/dao"
	"xs.bbs/internal/app/community/model"
	"xs.bbs/internal/app/community/service"
)

var (
	Model = &model.Community{}
	Set   = wire.NewSet(
		dao.CommunityDaoSet,
		service.CommunityServiceSet,
		controller.CommunityControllerSet,
	)
)
