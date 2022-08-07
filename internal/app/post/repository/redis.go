package repository

import (
	"math"
	"time"

	"xs.bbs/internal/pkg/constant"

	"github.com/go-redis/redis/v9"
	"xs.bbs/internal/pkg/constant/e"
)

// 推荐阅读
// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm/

// 本项目使用简化版的投票分数
// 投一票就加432分   86400/200  --> 200张赞成票可以给你的帖子续一天

/* 投票的几种情况：
   direction=1时，有两种情况：
   	1. 之前没有投过票，现在投赞成票    --> 更新分数和投票记录  差值的绝对值：1  +432
   	2. 之前投反对票，现在改投赞成票    --> 更新分数和投票记录  差值的绝对值：2  +432*2
   direction=0时，有两种情况：
   	1. 之前投过反对票，现在要取消投票  --> 更新分数和投票记录  差值的绝对值：1  +432
	2. 之前投过赞成票，现在要取消投票  --> 更新分数和投票记录  差值的绝对值：1  -432
   direction=-1时，有两种情况：
   	1. 之前没有投过票，现在投反对票    --> 更新分数和投票记录  差值的绝对值：1  -432
   	2. 之前投赞成票，现在改投反对票    --> 更新分数和投票记录  差值的绝对值：2  -432*2

   投票的限制：
   每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
   	1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
   	2. 到期之后删除那个 KeyPostVotedZSetPF
*/
const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值多少分
)

func (p *postRepo) Vote(userID, postID string, value float64) (err error) {
	// 1. 判断投票限制
	// 去redis取帖子发布时间
	postTime := p.rdb.ZScore(constant.RedisKey(constant.KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds { // 发帖超过一周不允许参与投票
		return e.ErrVoteTimeExpire
	}
	// 2和3需要放到一个pipeline事务中操作

	// 2. 更新贴子的分数
	// 先查当前用户给当前帖子的投票记录  ?? todo
	oldVal := p.rdb.ZScore(constant.RedisKey(constant.KeyPostVotedZSetPre+postID), userID).Val()
	var op float64
	// 如果当前投票值大于查询出oldVal
	if value > oldVal {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(oldVal - value) // 计算两次投票的差值
	pipeline := p.rdb.Pipeline()
	// 更新贴子的分数
	pipeline.ZIncrBy(constant.RedisKey(constant.KeyPostScoreZSet), op*diff*scorePerVote, postID)
	// 3. 记录用户为该贴子投票的数据
	if value == 0 { // 如果未投票，删除
		pipeline.ZRem(constant.RedisKey(constant.KeyPostVotedZSetPre+postID), userID)
	} else {
		pipeline.ZAdd(constant.RedisKey(constant.KeyPostVotedZSetPre+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	return
}
