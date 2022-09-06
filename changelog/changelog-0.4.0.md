# gorm callback

# cron定时任务

# 优化配置结构
  + 修改配置文件，将配置文件conf/app.ini修改为大驼峰命名
  + 优化配置读取及设置初始化顺序，将多模块的初始化函数放到启动流程中
    1. 将其他文件中的配置都删掉，统一在 setting 中处理以及修改 init 函数为 Setup 方法
        - pkg/setting/setting.go: init 函数修改为 Setup 方法
        - pkg/logging/file.go: 独立的 LOG 配置项删除，改为统一读取 setting
        - pkg/util/pagination.go
        - models/models.go: 独立读取的 DB 配置项删除，改为统一读取 setting
        - middleware/jwt/jwt.go
        - routers/api/v1/article.go
        - routers/api/v1/tag.go
        - routers/router.go
    2. 设置初始化的流程:
        - main.go

# 抽离 File
  + pkg/logging/file.go