package community

import (
	"gorm.io/gorm"
	"xs.bbs/internal/app/community/controller"
	"xs.bbs/internal/app/community/model"
	"xs.bbs/internal/app/community/repository"
	"xs.bbs/internal/app/community/service"
)

var (
	Entity = &model.Community{}
	//Set   = wire.NewSet(
	//	repo.CommunityDaoSet,
	//	service.CommunityServiceSet,
	//	controller.CommunityControllerSet,
	//)
)

func Build(db *gorm.DB) *controller.CommunityController {
	repo := repository.NewCommunityRepo(db)
	cs := service.NewCommunityService(repo)
	return controller.NewCommunityController(cs)
}
