package api

import (
	"exam/models"
	"exam/pkg/app"
	"exam/pkg/setting"
	"exam/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetExamList(c *gin.Context) {
	appG := app.Gin{C: c}
	params := make(map[string]interface{})
	data := make(map[string]interface{})

	// 参数
	pageNum := util.GetPage(c)
	userId := c.GetInt("userId")
	params["user_id"] = userId

	// 分组列表数据
	examList, err := models.GetExamList(pageNum, setting.AppSetting.PageSize, params)
	if err != nil {
		appG.ResponseErrMsg("数据查询出错")
		return
	}

	data["list"] = examList
	data["total"] = models.GetExamTotal(params)

	appG.ResponseSuccess("ok", data)
}

func AddExam(c *gin.Context) {
	appG := app.Gin{C: c}
	params := make(map[string]interface{})

	name := c.PostForm("name")
	time, _ := com.StrTo(c.PostForm("time")).Int()
	questionList := c.PostForm("questionList")

	if name == "" || time == 0 || questionList == "" {
		appG.ResponseErrMsg("参数不能为空")
		return
	}
	userId := c.GetInt("userId")

	params["userId"] = userId
	params["name"] = name
	params["time"] = time
	params["questionList"] = questionList

	err := models.AddExam(params)
	if err != nil {
		appG.ResponseErrMsg("新增失败")
		return
	}
	appG.ResponseSuccess("ok", nil)
}

func EditExam(c *gin.Context) {
	appG := app.Gin{C: c}

	examId := com.StrTo(c.PostForm("examId")).MustInt()
	name := c.PostForm("name")
	time, _ := com.StrTo(c.PostForm("time")).Int()
	questionList := c.PostForm("questionList")

	data := make(map[string]interface{})
	data["Name"] = name
	data["Time"] = time
	data["QuestionList"] = questionList

	err := models.EditExam(examId, data)
	if err != nil {
		appG.ResponseErrMsg("修改失败")
		return
	}
	appG.ResponseSuccess("ok", nil)
}

func DeleteExam(c *gin.Context) {
	appG := app.Gin{C: c}
	examId := com.StrTo(c.PostForm("examId")).MustInt()

	userId := c.GetInt("userId")
	exam, _ := models.GetExamDetail(examId)

	if exam.UserId != userId {
		appG.ResponseErrMsg("不能删除")
		return
	}

	err := models.DeleteExam(examId)
	if err != nil {
		appG.ResponseErrMsg("删除失败")
		return
	}
	appG.ResponseSuccess("ok", nil)
}
