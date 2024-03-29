package api

import (
	"exam/models"
	"exam/pkg/app"
	"exam/pkg/setting"
	"exam/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetQuestionList(c *gin.Context) {
	appG := app.Gin{C: c}

	questionType := c.DefaultPostForm("questionType", "")
	pageSize, _ := com.StrTo(c.DefaultPostForm("pageSize", "0")).Int()

	params := make(map[string]interface{})
	data := make(map[string]interface{})

	// 参数
	pageNum := util.GetPage(c)
	if pageSize == 0 {
		pageSize = setting.AppSetting.PageSize
	}
	userId := c.GetInt("userId")
	params["user_id"] = userId

	if questionType != "" {
		params["question_type"] = questionType
	}

	// 分组列表数据
	questionList, err := models.GetQuestionList(pageNum, pageSize, params)
	if err != nil {
		appG.ResponseErrMsg("数据查询出错")
		return
	}

	data["list"] = questionList
	data["total"] = models.GetQuestionTotal(params)

	appG.ResponseSuccess("ok", data)
}

func GetQuestionDetail(c *gin.Context) {
	appG := app.Gin{C: c}
	questionId := com.StrTo(c.PostForm("questionId")).MustInt()

	// 参数
	userId := c.GetInt("userId")

	// 分组列表数据
	questionDetail, err := models.GetQuestionDetail(questionId)
	if err != nil {
		appG.ResponseErrMsg("数据查询出错")
		return
	}
	if userId != questionDetail.UserId {
		appG.ResponseErrMsg("参数有误")
		return
	}
	appG.ResponseSuccess("ok", questionDetail)
}

func AddQuestion(c *gin.Context) {
	appG := app.Gin{C: c}
	params := make(map[string]interface{})

	questionType, _ := com.StrTo(c.PostForm("questionType")).Int()
	title := c.PostForm("title")
	remark := c.PostForm("remark")
	option := c.PostForm("option")

	if questionType == 0 || title == "" {
		appG.ResponseErrMsg("参数不能为空")
		return
	}
	userId := c.GetInt("userId")

	params["userId"] = userId
	params["questionType"] = questionType
	params["title"] = title
	params["remark"] = remark
	params["option"] = option

	err := models.AddQuestion(params)
	if err != nil {
		appG.ResponseErrMsg("新增失败")
		return
	}
	appG.ResponseSuccess("ok", nil)
}

func EditQuestion(c *gin.Context) {
	appG := app.Gin{C: c}

	questionId := com.StrTo(c.PostForm("questionId")).MustInt()
	title := c.PostForm("title")
	remark := c.PostForm("remark")
	option := c.PostForm("option")

	data := make(map[string]interface{})
	data["Title"] = title
	data["Remark"] = remark
	data["Option"] = option

	err := models.EditQuestion(questionId, data)
	if err != nil {
		appG.ResponseErrMsg("修改失败")
		return
	}
	appG.ResponseSuccess("ok", nil)
}

func DeleteQuestion(c *gin.Context) {
	appG := app.Gin{C: c}
	questionId := com.StrTo(c.PostForm("questionId")).MustInt()

	userId := c.GetInt("userId")
	question, _ := models.GetQuestionDetail(questionId)

	if question.UserId != userId {
		appG.ResponseErrMsg("不能删除")
		return
	}

	err := models.DeleteQuestion(questionId)
	if err != nil {
		appG.ResponseErrMsg("删除失败")
		return
	}
	appG.ResponseSuccess("ok", nil)
}
