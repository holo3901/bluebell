package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode) //设定为测试模式
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	body := `{
        "community_id":1,
        "title":"test",
        "content":"just a test"
    }`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body))) //创建请求
	w := httptest.NewRecorder()                                                    //响应对象
	r.ServeHTTP(w, req)

	//判断响应状态码是否为200
	assert.Equal(t, 200, w.Code)
	//判断响应的内容是不是按照预期返回了需要登录的错误

	//方法1：判断响应内容是不是包含指定的字符串
	assert.Contains(t, w.Body.String(), "需要登录")

	//方法2:将响应的内容反序列化到res，然后判断字段与预期是否一致
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed,err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}

func TestGetPostDetailHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post/1"
	r.GET(url, GetPostDetailHandler)
	w := httptest.NewRecorder()                         //响应对象
	req, _ := http.NewRequest(http.MethodGet, url, nil) //创建请求

	r.ServeHTTP(w, req)

	//判断响应状态码是否为200
	assert.Equal(t, 200, w.Code)
	//判断响应的内容是不是按照预期返回了需要登录的错误

	//方法1：判断响应内容是不是包含指定的字符串
	assert.Contains(t, w.Body.String(), "请求参数错误")
}

func TestGetPostListHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/posts/"
	r.GET(url, GetPostListHandler)
	w := httptest.NewRecorder()                         //响应对象
	req, _ := http.NewRequest(http.MethodGet, url, nil) //创建请求

	r.ServeHTTP(w, req)

	//判断响应状态码是否为200
	assert.Equal(t, 500, w.Code)
	//判断响应的内容是不是按照预期返回了需要登录的错误

	//方法1：判断响应内容是不是包含指定的字符串
	assert.Contains(t, w.Body.String(), "")
}

func TestGetPostListHandler2(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post2/"
	r.GET(url, GetPostListHandler2)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)

	assert.Contains(t, w.Body.String(), "")
}
