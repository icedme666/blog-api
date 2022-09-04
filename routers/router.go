package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"xiamei.guo/blog-api/docs"
	"xiamei.guo/blog-api/middleware/jwt"
	"xiamei.guo/blog-api/pkg/setting"
	"xiamei.guo/blog-api/pkg/upload"
	"xiamei.guo/blog-api/routers/api"
	"xiamei.guo/blog-api/routers/api/v1"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func InitRouter() *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	// use ginSwagger middleware to serve the API docs
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath())) //读取静态文件
	r.GET("/test")
	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取多个文章标签
		apiv1.GET("/tags", v1.GetTags)

		//新增文章标签
		apiv1.POST("/tags", v1.AddTags)

		//修改文章标签
		apiv1.PUT("/tags/:id", v1.EditTag)

		//删除文章标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)

		//获取文章
		apiv1.GET("/articles/:id", v1.GetArticle)

		//新增文章
		apiv1.POST("/articles", v1.AddArticle)

		//修改文章
		apiv1.PUT("/articles/:id", v1.EditArticle)

		//删除文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
