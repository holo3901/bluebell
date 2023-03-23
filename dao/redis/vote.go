package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

//投1票就加432分 86400、200 ->需要200张赞成票可以给帖子续一天，时间戳时间+1
/*
这些都需要更新分数和投票记录
有票的几种情况:
direction=1时，有两种情况
1.之前没有投票，现在投赞成票            ==》差值绝对值为1
2.之前投反对票，现在改投赞成票           ==》差值绝对值为2
direction=0时，有两种情况:
1.之前投过赞成票，现在要取消投票         ==》差值绝对值为1
2.之前投过反对票,现在要取消投票          ==》差值绝对值为1
direction=-1时，有两种情况：
1.之前没有投票，现在投反对票            ==》差值绝对值为1
2.之前投赞成票，现在改投反对票           ==》差值绝对值为2

投票的限制:
每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了
   1.到期之后将redis中保存的赞成票数及反对票数存储到mysql表
   2.到期之后删除那个keyPostVotedSetPF
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 //每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepested   = errors.New("不允许重复投票")
)

func CreatePost(postID int64, communityID int64) error {
	pipeline := client.TxPipeline() //pipeline事务,两者必须同时成功才能运行
	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//把帖子id加到社区的set
	cKey := getRedisKey(keyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err := pipeline.Exec()
	return err
	/*

	 */
	//帖子分数
	/*pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})*、
	/*_, err = client.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	*/
}
func VoteForPost(userID, postID string, value float64) error {
	//1.判断投票限制
	//去redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//2.更新帖子分数
	//先查之前用户给当前帖子的投票记录
	ov := client.ZScore(getRedisKey(keyPostVotedZSetPF+postID), userID).Val()
	//更新:如果这一次有票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepested
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算两次投票的差值

	//2和3需要放到一个pipeline事务中操作
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	/*
		_, err := client.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID).Result()
		if err != nil {
			return err
		}*/
	//3.记录用户为该帖子投过票
	if value == 0 {
		pipeline.ZRem(getRedisKey(keyPostVotedZSetPF+postID), userID)
		/*
			_, err = client.ZRem(getRedisKey(keyPostVotedZSetPF+postID), userID).Result()
			/
		*/
	}
	pipeline.ZAdd(getRedisKey(keyPostVotedZSetPF+postID), redis.Z{
		Score:  value, //赞成还是反对票
		Member: userID,
	})
	/*
		_, err = client.ZAdd(getRedisKey(keyPostVotedZSetPF+postID), redis.Z{
			Score:  value, //赞成还是反对票
			Member: userID,
		}).Result()
	*/
	_, err := pipeline.Exec() //这个时候运行pipeline
	return err
}
