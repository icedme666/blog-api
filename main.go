package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xiamei.guo/blog-api/pkg/setting"
	"xiamei.guo/blog-api/routers"
)

func main() {
	//router := gin.Default()
	router := routers.InitRouter()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test"})
	})
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
