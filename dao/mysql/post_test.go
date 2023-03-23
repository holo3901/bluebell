package mysql

import (
	"testing"
	"xxx/models"
	"xxx/settings"
)

func init() { //用此方法db初始化,init函数会优先自动
	dbcfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "miku1314",
		Db:           "sql_demo",
		Port:         3306,
		MaxOpenConns: 200,
		MaxIdleConns: 50,
	}
	err := Init(&dbcfg)
	if err != nil {
		panic(err)
	}
}
func TestCreatePost(t *testing.T) { //不能单用这个函数来进行测试，因为db默认为空指针，需要初始化
	post := models.Post{
		ID:          4, //因为是添加进入库，每次都要改
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("createPost insert record into mysql failed,err:%v\n", err)
	}
	t.Logf("createPost insert record into mysql success")
}

func TestGetPostById(t *testing.T) {
	post := models.Post{
		ID:          1,
		AuthorID:    123,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	_, err := GetPostById(post.ID)
	if err != nil {
		t.Fatalf("TestGetPostById find record failed,err:%v\n", err)
	}
	t.Logf("TestGetPostById find record success")
}

func TestGetPostList(t *testing.T) {

	_, err := GetPostList(1, 1)
	if err != nil {
		t.Fatalf("TestGetPostList find record failed,err:%v\n", err)
	}
	t.Logf("TestGetPostList find record success")
}

func TestGetPostListByIDs(t *testing.T) {
	_, err := GetPostListByIDs([]string{"123", "24"})
	if err != nil {
		t.Fatalf("TestGetPostList find record failed,err:%v\n", err)
	}
	t.Logf("TestGetPostList find record success")
}
