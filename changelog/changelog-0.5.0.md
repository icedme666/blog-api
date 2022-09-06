# 实现上传图片
  + 图片上传接口
    - 流程：
      1. 图片名md5加密：pkg/util/md5.go
      2. 封装image的处理逻辑：
         - pkg/upload/image.go
         - 错误码：pkg/e/code.go、msg.go
      3. 上传图片的业务逻辑：routers/api/upload.go
      4. 添加路由：routers/router.go
    - 验证：POST http://127.0.0.1:8000/upload，选择form-data，参数为image，格式为图片
  + 实现http.FileServer
    - 增加路由：routers/router.go的 AddArticle、EditArticle 两个接口
    - 验证：GET http://127.0.0.1:8000/upload/images/59b514174bffe4ae402b3d63aad79fe0.jpeg
  + 修改文章接口：routers/api/v1/article.go
    - 新增、更新文章接口：支持入参 cover_image_url
    - 新增、更新文章接口：增加对 cover_image_url 的非空、最长长度校验

# 优化应用结构
  + 重构箭头型代码
  + 方法：
    - 错误提前返回：
      - pkg/app/request.go
      - pkg/app/response.go
    - 统一返回方法：routers/api/v1/article.go
    - 抽离 Service，减轻 routers/api 的逻辑，进行分层：service/article_service
    - 增加 gorm 错误判断，让错误提示更明确（增加内部错误码）： models/article.go

# 实现Redis缓存
  1. 配置：conf/app.ini
  2. 缓存Prefix：pkg/e/cache.go
  3. 缓存Key：编写获取缓存 KEY 的方法，直接参考传送门即可
     - cache_service/article.go
     - cache_service/tag.go
  4. Redis工具包：pkg/gredis/redis.go
     - 设置 RedisConn 为 redis.Pool（连接池）并配置了它的一些参数
     - 封装基础方法