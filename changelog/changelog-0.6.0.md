# 实现导出、导入 Excel
  + 配置
    - conf/app.ini
    - pkg/setting/setting.go
  + 导出
    - pkg/export/excel.go
    - service/tag.go
    - routers/api/tag.go
    - routers/router.go
  + 访问：POST http://127.0.0.1:8000/tags/export

# 生成带二维码和文字的海报
* 二维码
  + 工具包：pkg/qrcode/qrcode.go
  + 路由
    1. 接口开发：routers/api/v1/article.go
    2. 新增路由：routers/router.go
    3. 编写对应的生成逻辑
  + 访问：http://127.0.0.1:8000/api/v1/articles/poster/generate?token=$token，查看runtime/qrcode目录下生成二维码
* 合成海报
  + service/article_service/article_poster.go
  + 路由方法：routers/api/v1/article.go
  + StaticFS：routers/router.go
* 在图片上绘制文字
  + 绘制文字的业务逻辑：service/article_service/article_poster.go
  + generate：service/article_service/article_poster.go

# 用Nginx部署Go应用
  * 启动nginx
  * 配置hosts：127.0.0.1       api.blog.com
  * 配置nginx.conf
    ```bash
    worker_processes  1;

    events {
    worker_connections  1024;
    }
    
    
    http {
    include       mime.types;
    default_type  application/octet-stream;
    
        sendfile        on;
        keepalive_timeout  65;
    
        server {
            listen       8081;
            server_name  api.blog.com;
    
            location / {
                proxy_pass http://127.0.0.1:8000/;
            }
        }
    }
    ```
  * 打包并启动：make
  * 访问：GET http://api.blog.com:8000/auth?username=test&password=test123456