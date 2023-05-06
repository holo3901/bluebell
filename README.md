# gin-mall

**基于 gin+sqlx+mysql读写分离 的一个社区系统**
# 项目的主要功能介绍

- 用户注册登录(JWT-Go鉴权)
- 用户基本信息修改，修改密码
- 帖子的发布，删除，浏览和投票等
- 社区的增加，删除，修改等
- 支持事务，对帖子创建时间和人帖子加入相应社区错误进行回退处理
- 跨域处理和使用令牌桶对请求进行限流处理
# 配置文件
`conf/config.yaml` 文件配置

```yaml
name: "bluebell"
mode: "release"
port: 8084
version: "v0.0.1"
start_time: "2020-07-01"
machine_id: 1

auth:
  jwt_expire: 8760

log:
  level: "info"
  filename: "web_app.log"
  max_size: 200
  max_age: 30
  max_backups: 7
mysql:
  host: 127.0.0.1
  port: 3306
  user: "root"
  password: "root1234"
  dbname: "bluebell"
  max_open_conns: 200
  max_idle_conns: 50
redis:
  host: 127.0.0.1
  port: 6379
  password: ""
  db: 0
  pool_size: 100
```
