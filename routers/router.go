package routers

import (
	"github.com/gin-gonic/gin"
	"xiamei.guo/blog-api/pkg/setting"
	v1 "xiamei.guo/blog-api/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		//获取多个文章标签
		apiv1.GET("/tags", v1.GetTags)

		//新增文章标签
		apiv1.POST("/tags", v1.AddTags)

		//修改文章标签
		apiv1.PUT("/tags", v1.EditTag)

		//删除文章标签
		apiv1.DELETE("/tags", v1.DeleteTag)
	}
	return r
}
