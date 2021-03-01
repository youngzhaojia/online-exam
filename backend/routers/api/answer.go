package api

import (
	"exam/models"
	"exam/pkg/app"
	"exam/pkg/setting"
	"exam/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetAnswerList(c *gin.Context) {
	appG := app.Gin{C: c}
	params := make(map[string]interface{})
	data := make(map[string]interface{})

	// 参数
	pageNum := util.GetPage(c)
	examId := c.DefaultPostForm("examId", "")
	userId := c.GetInt("userId")
	params["teacher_id"] = userId
	if examId != "" {
		params["exam_id"] = examId
	}

	// 分组列表数据
	answerList, err := models.GetAnswerList(pageNum, setting.AppSetting.PageSize, params)
	if err != nil {
		appG.ResponseErrMsg("数据查询出错")
		return
	}

	data["list"] = answerList
	data["total"] = models.GetAnswerTotal(params)

	appG.ResponseSuccess("ok", data)
}

func GetAnswerDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	answerId := com.StrTo(c.PostForm("answerId")).MustInt()

	// 参数
	userId := c.GetInt("userId")

	// 分组列表数据
	answerDetail, err := models.GetAnswerDetail(answerId)
	if err != nil {
		appG.ResponseErrMsg("试卷数据查询出错")
		return
	}
	if userId != answerDetail.TeacherId {
		appG.ResponseErrMsg("参数有误")
		return
	}

	questionList, err := models.GetQuestionListByIds(answerDetail.QuestionList)
	if err != nil {
		appG.ResponseErrMsg("试题数据查询出错")
		return
	}

	data := make(map[string]interface{})
	data["answer"] = answerDetail
	data["questionList"] = questionList

	appG.ResponseSuccess("ok", data)
}

func AddAnswer(c *gin.Context) {
	appG := app.Gin{C: c}
	params := make(map[string]interface{})

	answerList := c.PostForm("answerList")

	if answerList == "" {
		appG.ResponseErrMsg("参数不能为空")
		return
	}
	userId := c.GetInt("userId")

	params["userId"] = userId
	params["answerList"] = answerList

	err := models.AddAnswer(params)
	if err != nil {
		appG.ResponseErrMsg("新增失败")
		return
	}
	appG.ResponseSuccess("ok", nil)
}

func EditAnswer(c *gin.Context) {
	appG := app.Gin{C: c}

	answerId := com.StrTo(c.PostForm("answerId")).MustInt()
	answerList := c.PostForm("answerList")

	data := make(map[string]interface{})
	data["AnswerList"] = answerList

	err := models.EditAnswer(answerId, data)
	if err != nil {
		appG.ResponseErrMsg("修改失败")
		return
	}
	appG.ResponseSuccess("ok", nil)
}

func DeleteAnswer(c *gin.Context) {
	appG := app.Gin{C: c}
	answerId := com.StrTo(c.PostForm("answerId")).MustInt()

	userId := c.GetInt("userId")
	answer, _ := models.GetAnswerDetail(answerId)

	if answer.UserId != userId {
		appG.ResponseErrMsg("不能删除")
		return
	}

	err := models.DeleteAnswer(answerId)
	if err != nil {
		appG.ResponseErrMsg("删除失败")
		return
	}
	appG.ResponseSuccess("ok", nil)
}
