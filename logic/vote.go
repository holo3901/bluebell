package logic

import (
	"strconv"
	"xxx/dao/redis"
	"xxx/models"

	"go.uber.org/zap"
)

//投票功能
//1.用户投票的数据 基于用户投票的相关算法:http://www.ruanyifeng.com/blog/algorithm/ ,此项目用简化版的投票分数

func VoteForPost(userID int64, p *models.ParamVoteData) error {
	//1.判断投票限制
	zap.L().Debug("voteForPost", zap.Int64("userID", userID), zap.String("postID", p.PostID), zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
	//2.更新帖子分数
	//3.记录用户为该帖子投过票
}
