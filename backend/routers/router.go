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

	// 考试详情
	apiGroup.POST("/exam/detail", api.GetExamDetail)
	apiGroup.POST("/answer/add", api.AddAnswer)
	// jwt验证
	apiGroup.Use(jwt.JWT())
	// 用户
	{
		apiGroup.POST("/user/info", api.GetUserInfo)
	}
	// 问题
	{
		apiGroup.POST("/question/list", api.GetQuestionList)
		apiGroup.POST("/question/detail", api.GetQuestionDetail)
		apiGroup.POST("/question/add", api.AddQuestion)
		apiGroup.POST("/question/edit", api.EditQuestion)
		apiGroup.POST("/question/delete", api.DeleteQuestion)
	}
	// 试卷
	{
		apiGroup.POST("/exam/list", api.GetExamList)
		apiGroup.POST("/exam/add", api.AddExam)
		apiGroup.POST("/exam/edit", api.EditExam)
		apiGroup.POST("/exam/delete", api.DeleteExam)
	}
	// 答题
	{
		apiGroup.POST("/answer/list", api.GetAnswerList)
		apiGroup.POST("/answer/detail", api.GetAnswerDetail)
	}
	return r
}
