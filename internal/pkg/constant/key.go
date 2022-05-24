package constant

const (
	KeyMySecret  = "xiangshouduan.xs.bbs"
	KeyCtxUserID = "userID"
)
const (
	KeyPrefix           = "xs.bbs:"
	KeyPostTimeZSet     = "post:time"   // zset;帖子及发帖时间
	KeyPostScoreZSet    = "post:score"  // zset;帖子及投票的分数
	KeyPostVotedZSetPre = "post:voted:" // zset;记录用户及投票类型;参数是post_id
)

// RedisKey redis key
func RedisKey(key string) string {
	return KeyPrefix + key
}
