# DJ-Blog
创新实践个人博客任务

# 项目简介
- 必做
- [ ] 首页
- [ ] 文章列表
- [ ] 文章详情
- [ ] 文章编辑

- 选做(画饼)
- [ ] 用户
- [ ] 评论
- [ ] 搜索
- [ ] 音乐播放器
- [ ] 关于

idea:
- [ ] 按钮导入markdown
- [ ] 按钮导出markdown

# 部署说明
## 常规部署
> 1. 配置文件在config/config.yaml，配置好相关参数
> 2. 日志文件如果在文件夹下，需要先创建好文件夹
> 3. 运行`go mod download`下载依赖
> 4. 运行`go run main.go`启动项目

## Docker部署
> 1. 修改docker-compose.yaml中配置文件的挂载位置
> 2. 运行`docker-compose up -d`启动项目