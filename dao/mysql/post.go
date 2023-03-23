package mysql

import (
	"strings"
	"xxx/models"

	"github.com/jmoiron/sqlx"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
              post_id,title,content,author_id,community_id)
              values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

//GetPostById 根据id查询单个帖子数据
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time 
             from post
             where post_id=?
             `
	err = db.Get(post, sqlStr, pid)
	return
}

//GetPostList 查询帖子列表函数
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select
              post_id,title,content,author_id,community_id,create_time
              from post
              ORDER BY create_time
              DESC 
              limit ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

//根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
    from post
    where post_id in (?)
    order by FIND_IN_SET(post_id,?) ` //find_in_set(str,strlist),str为要查询的字符串,strlist字段名 参数以","分隔，查询字段中包含str的结果，
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ",")) //strings.Join(ids,",")将ids中的子串连接成一个单独的字符串，子串之间用“,”隔开
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query) // sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	err = db.Select(&postList, query, args...)
	return
}
