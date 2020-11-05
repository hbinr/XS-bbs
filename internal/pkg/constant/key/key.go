package key

const (
	KeyMySecret  = "xiangshouduan.xs.bbs"
	KeyCtxUserID = "userID"

	// redis key
	KeyPrefix           = "xs.bbs:"
	KeyPostTimeZset     = "post:time"   // zset;帖子及发帖时间
	KeyPostScoreZset    = "post:score"  // zset;帖子及投票的分数
	KeyPostVotedZsetPre = "post:voted:" // zset;记录用户及投票类型;参数是post_id
)

// Redis 凭借redis key
func Redis(key string) string {
	return KeyPrefix + key
}
