# 通过docker部署应用
  + 编写 Dockerfile
  + mysql：
    - 拉取镜像：```docker pull mysql```
    - 创建并运行mysql容器：```docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootroot -d mysql```
    - 修改配置文件：conf/app.ini
  + 应用：
    - 构建镜像：```docker build -t gin-blog-docker .```
    - 创建并运行一个新容器：```docker run --link mysql:mysql -p 8000:8000 gin-blog-docker```