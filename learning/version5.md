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
# 优化你的应用结构

# 实现Redis缓存

# 实现导出、导入 Excel

# 生成二维码、合并海报

# 在图片上绘制文字

# 用Nginx部署Go应用