# gin-mall

**基于 gin+sqlx+mysql读写分离 的一个社区系统**
# 项目的主要功能介绍

- 用户注册登录(JWT-Go鉴权)
- 用户基本信息修改，修改密码
- 帖子的发布，删除，浏览和投票等
- 社区的增加，删除，修改等
- 支持事务，对帖子创建时间和人帖子加入相应社区错误进行回退处理
- 跨域处理和使用令牌桶对请求进行限流处理
