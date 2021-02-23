package routers

import (
	"exam/jwt"
	"exam/pkg/setting"
	"exam/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	apiGroup := r.Group("/api")
	// 用户登录
	apiGroup.POST("/user/login", api.Login)
	apiGroup.POST("/user/register", api.Register)
	// jwt验证
	apiGroup.Use(jwt.JWT())
	// 问题
	{
		apiGroup.POST("/question/list", api.GetQuestionList)
		apiGroup.POST("/question/add", api.AddQuestion)
		apiGroup.POST("/question/edit", api.EditQuestion)
		apiGroup.POST("/question/delete", api.DeleteQuestion)
	}
	// 试卷
	// {
	// 	apiGroup.GET("/word/list", api.GetWordList)
	// 	apiGroup.POST("/word/add", api.AddWord)
	// 	apiGroup.POST("/word/delete", api.DeleteWord)
	// }
	return r
}
