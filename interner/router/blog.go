package router

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/interner/api"
)

type blogRouter struct{}

func InitBlogRouter(router *gin.RouterGroup) {
	blogRouter := router.Group("/blog")
	blogApi := api.BlogApi{}
	{
		blogRouter.GET("/list", blogApi.GetBlogList)
	}
}