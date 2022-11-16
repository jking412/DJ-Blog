# DJ-Blog
创新实践个人博客任务

# 项目简介
- 首页（可选）
- 文章列表
  - 登录提示
  - 搜索栏
  - tag汇总
  - 导航栏
    - 首页
    - 收藏
    - 点赞
    - 个人
    - 钱包
  - 文章
    - title
    - avatar
    - tag
    - author
    - excerpt
    - createdtime
    - 点赞
    - 收藏
    - 评论
  - 分页

- 文章详情
  - title
  - avatar
  - tag
  - author
  - 正文
  - createdtime
  - 点赞
  - 收藏
  - 评论
  - 上/下一篇文章

- 登录（模态窗口实现）
  - 邮箱/用户名+密码
  - 邮箱+验证码
  - github三方登录






# 部署说明
## 常规部署(不推荐)
> 1. 配置文件在config/config.yaml，配置好相关参数
> 2. 下载Zinc，如果使用mysql要先配置好mysql
> 3. 运行`go mod download`下载依赖
> 4. 运行`go run main.go`启动项目

## Docker部署(推荐)
> 1. 修改docker-compose.yaml中配置文件的挂载位置
> 2. 运行`docker-compose up -d`启动项目