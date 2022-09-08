# blog-api

# Feature
* Auth(jwt)
  + GET http://127.0.0.1:8000/auth?username=test&password=test123456
* Tag
  + GET http://127.0.0.1:8000/api/v1/tags
  + POST http://127.0.0.1:8000/api/v1/tags?name=2&state=1&created_by=test
  + PUT http://127.0.0.1:8000/api/v1/tags/1?name=edit1&state=0&modified_by=edit1
  + DELETE http://127.0.0.1:8000/api/v1/tags/1
* Article
  + POST http://127.0.0.1:8000/api/v1/articles?tag_id=1&title=test1&desc=test-desc&content=test-content&created_by=test-created&state=1
  + GET http://127.0.0.1:8000/api/v1/articles
  + GET http://127.0.0.1:8000/api/v1/articles/1
  + PUT http://127.0.0.1:8000/api/v1/articles/1?tag_id=1&title=test-edit1&desc=test-desc-edit&content=test-content-edit&modified_by=test-created-edit&state=0
  + DELETE：http://127.0.0.1:8000/api/v1/articles/1
* 图片上传
  + POST http://127.0.0.1:8000/upload，选择form-data，参数为image，格式为图片
  + GET http://127.0.0.1:8000/upload/images/59b514174bffe4ae402b3d63aad79fe0.jpeg
* excel导出、导入
  + POST http://127.0.0.1:8000/tags/export
  + POST http://127.0.0.1:8000/tags/import，选择form-data，参数为file，格式为excel


# 应用优化
* log
* 优雅重启
  + 目的
    - 不关闭现有连接（正在运行中的程序）
    - 新的进程启动并替代旧进程
    - 新的进程接管新的连接
    - 连接要随时响应用户的请求，当用户仍在请求旧进程时要保持连接，新用户应请求新进程，不可以出现拒绝请求的情况
  + 流程
    1. 替换可执行文件或修改配置文件
    2. 发送信号量 SIGHUP
    3. 拒绝新连接请求旧进程，但要保证已有连接正常
    4. 启动新的子进程
    5. 新的子进程开始 Accet
    6. 系统将新的请求转交新的子进程
    7. 旧进程处理完所有旧连接后正常结束
* swagger
  + http://127.0.0.1:8000/swagger/index.html#/
* 通过gorm实现callbacks
  + 文档：https://gorm.io/docs/write_plugins.html
  + 通过回调实现更新created_on、modified_on字段和软删除
* corn
  + 通过corn实现定时任务清理无效数据
  + 运行
    ```bash
    go run cron.go
    ```
* 由程序把控初始化的先后顺序
  + 在 Go 中，当存在多个 init 函数时，执行顺序为： 
    - 相同包下的 init 函数：按照源文件编译顺序决定执行顺序（默认按文件名排序）
    - 不同包下的 init 函数：按照包导入的依赖关系决定先后顺序
* redis

# deploy
* docker
  ```bash
  docker build -t blog-api .
  cd deploy/mysql
  docker build -t blog-mysql .
  docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -v /Users/xiamei.guo/data/docker-mysql:/var/lib/mysql -d blog-mysql
  docker run --link mysql:mysql -p 8000:8000 blog-api
  ```
* 构建 Scratch 镜像
  ```bash
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o blog-api .
  docker build -t blog-api-scratch -f ScratchDockerfile .
  docker run --link mysql:mysql -p 8000:8000 blog-api-scratch
  ```