package post

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"xs.bbs/internal/app/post/controller"

	commuDao "xs.bbs/internal/app/community/dao"
	postDao "xs.bbs/internal/app/post/dao"
	"xs.bbs/internal/app/post/model"
	"xs.bbs/internal/app/post/service"
	userDao "xs.bbs/internal/app/user/dao"
)

var (
	Model = &model.Post{}
	//Set   = wire.NewSet(
	//	dao.PostDaoSet,
	//	service.PostServiceSet,
	//	controller.PostControllerSet,
	//)
)

/*
	笔记：
	曾经想 service.NewPostService()修改为service.NewPostService(store)，其中store是一个接口，包含了全部数据访问接口
	然后service层可以链式调用：
		c.repository.Post().Find(args1)....
		c.repository.User().Create(args2)....
		c.repository.Community().Update(args3)....

	但是后来否决了这个想法，原因如下：
		1.我们应该给函数传入它关心的最小集合作为参数，而不是，我有一个 struct，当某个函数需要这个 struct 的成员的时候，
	    我们把整个 struct 都作为参数传递进去。应该仅仅传递函数关心的最小集合。
	    2.不要链式调用方法，传进去的一整条调用链对函数来说，都是无关的耦合，只会让代码更 hard to change，让工程师惧怕去修改
		3.只管命令不要询问，直接做具体的事，不要去找是哪个接口，链式调用就出现了先找到Post接口，然后再调用其实现
*/
func Init(engine *gin.Engine, db *gorm.DB, rdb *redis.Client) *controller.PostController {
	postDao := postDao.NewPostDao(db, rdb)
	userDao := userDao.NewUserDao(db)
	commuDao := commuDao.NewCommunityDao(db)
	postService := service.NewPostService(postDao, userDao, commuDao)
	return controller.NewPostController(engine, postService)
}
