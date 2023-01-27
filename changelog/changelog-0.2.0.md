# log
  + 新建logging包
    - logging/file.go
    - logging/log.go
  + 验证：访问接口，并到$GOPATH/gin-blog/runtime/logs下查看日志
# 优雅重启
  + 目的
    - 不关闭现有连接（正在运行中的程序）
    - 新的进程启动并替代旧进程
    - 新的进程接管新的连接
    - 连接要随时响应用户的请求，当用户仍在请求旧进程时要保持连接，新用户应请求新进程，不可以出现拒绝请求的情况
  + 流程
    - 替换可执行文件或修改配置文件
    - 发送信号量 SIGHUP
    - 拒绝新连接请求旧进程，但要保证已有连接正常
    - 启动新的子进程
    - 新的子进程开始 Accet
    - 系统将新的请求转交新的子进程
    - 旧进程处理完所有旧连接后正常结束
  + 实现
    1. 依赖包
       ```go get -u github.com/fvbock/endless```
    2. 修改main.go
  + 验证
    1. 编译：```go build main.go```
    2. 执行：```./main```
    3. 查看：
       - 启动成功后输出了pid为 48601；在另外一个终端执行 kill -1 48601 ，检验先前服务的终端效果。
       - 可以看到该命令已经挂起，并且 fork 了新的子进程 pid 为 48755
    4. 唤醒：在 postman 上再次访问接口，可以发现子进程再次启动
  + http.Server的Shutdown()：如果你的Golang >= 1.8，也可以考虑使用 http.Server 的 Shutdown 方法
# swagger
  + 安装swag：
    - ```go get -u github.com/swaggo/swag/cmd/swag@v1.6.5```
    - 验证：```swag -v```
  + 安装gin-swagger
    ```
     go get -u github.com/swaggo/gin-swagger@v1.2.0 
     go get -u github.com/swaggo/files
     go get -u github.com/alecthomas/template
    ```
  + 初始化
    - 编写 API 注释
    - 路由：routers/router.go
    - 执行生成命令(在项目根目录，完毕后会在项目根目录下生成docs)：
      ```swag init```
  + 验证：```GET http://127.0.0.1:8000/swagger/index.html```