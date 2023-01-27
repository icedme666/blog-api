# 初始化目录
  + conf：用于存储配置文件
  + middleware：应用中间件
  + models：应用数据库模型
  + pkg：第三方包
  + routers 路由逻辑处理
  + runtime：应用运行时数据

# 编辑go.mod

# 初始化项目数据库
  + 数据库名称：blog
  + 表：
    - blog_tag
    - blog_article
    - blog_auth

# 配置
  + 依赖包：
    ```go get -u github.com/go-ini/ini```
  + conf/app.ini
  + setting/setting.go

# API 错误码包
  + e/code.go
  + e/msg.go

# 工具包
  + 依赖包：
    ```go get -u github.com/unknwon/com```
  + pkg/util/pagination.go

# models init
  + 依赖包：
    - gorm：```go get -u github.com/jinzhu/gorm```
    - mysql：```go get -u github.com/go-sql-driver/mysql```
  + models/models.go

# 项目启动
  + main.go

# 运行
  ```go run main.go```

# 验证
  ```curl 127.0.0.1:8000/test```

# 路由
  + 路由空壳：
    - routers/api/v1/tag.go
    - routers/api/v1/article.go
  + 注册路由：routers/router.go
  + 检测路由是否注册成功：
    ```go run main.go```  查看输出

# 接口
  1. 定义接口
     + 标签：
       - 获取标签列表：GET("/tags”)
       - 新建标签：POST("/tags”)
       - 更新指定标签：PUT("/tags/:id”)
       - 删除指定标签：DELETE("/tags/:id”)
     + 文章
       - 获取文章列表：GET("/articles”)
       - 获取指定文章：POST("/articles/:id”)
       - 新建文章：POST("/articles”)
       - 更新指定文章：PUT("/articles/:id”)
       - 删除指定文章：DELETE("/articles/:id”)
  2. 依赖包：
     ```go get -u github.com/astaxie/beego/validation```
  3. 实现
     + models逻辑：
       - models/tag.go
       - models/article.go
     + 路由逻辑：
       - routers/api/v1/tag.go
       - routers/api/v1/article.go

# 访问
  ```
  POST http://127.0.0.1:8000/api/v1/tags?name=2&state=1&created_by=test
  PUT http://127.0.0.1:8000/api/v1/tags/1?name=edit1&state=0&modified_by=edit1
  DELETE http://127.0.0.1:8000/api/v1/tags/1

  POST：http://127.0.0.1:8000/api/v1/articles?tag_id=1&title=test1&desc=test-desc&content=test-content&created_by=test-created&state=1
  GET：http://127.0.0.1:8000/api/v1/articles
  GET：http://127.0.0.1:8000/api/v1/articles/1
  PUT：http://127.0.0.1:8000/api/v1/articles/1?tag_id=1&title=test-edit1&desc=test-desc-edit&content=test-content-edit&modified_by=test-created-edit&state=0
  DELETE：http://127.0.0.1:8000/api/v1/articles/1
  ```

# 身份验证
  + 依赖包
    ```go get -u github.com/dgrijalva/jwt-go```
  + 实现jwt工具包：pkg/util/jwt.go
  + 实现用于Gin的中间件：middleware/jwt/jwt.go
  + 新增获取Token的API：
    - models/auth.go
    - routers/api/auth.go
    - 将中间件接入Gin：routers/router.go
  + 验证：
    ```
    GET http://127.0.0.1:8000/auth?username=test&password=test123456
    GET http://127.0.0.1:8000/api/v1/articles?token=...
    ```