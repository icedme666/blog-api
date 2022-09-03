package routers

import (
	"github.com/gin-gonic/gin"
	"xiamei.guo/blog-api/middleware/jwt"
	"xiamei.guo/blog-api/pkg/setting"
	"xiamei.guo/blog-api/routers/api"
	"xiamei.guo/blog-api/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取多个文章标签
		apiv1.GET("/tags", v1.GetTags)

		//新增文章标签
		apiv1.POST("/tags", v1.AddTags)

		//修改文章标签
		apiv1.PUT("/tags", v1.EditTag)

		//删除文章标签
		apiv1.DELETE("/articles", v1.DeleteTag)

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
