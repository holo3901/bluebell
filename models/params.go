package models

//定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 处理注册请求参数
type ParamSignUp struct { //blinding用于校验，该字段不能为空,eqfield使得rePassword必须等于password，否则报错
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

//ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	//UserID 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`              //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票1还是反对票-1,这个地方不能使用required，应为会默认把0当作空值
}

// ParamPostList 获取帖子子列表query string参数

// bluebell/models/params.go

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	Page        int64  `json:"page" form:"page"`                 //页码
	Size        int64  `json:"size" form:"size"`                 //每页数据量
	Order       string `json:"order" form:"order"`               //排序
	CommunityID int64  `json:"community_id" form:"community_id"` //可为空
}
