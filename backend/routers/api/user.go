package api

import (
	"exam/models"
	"exam/pkg/app"
	"exam/pkg/e"
	"exam/pkg/logging"
	"exam/pkg/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 用户登录
func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	mobile := c.PostForm("mobile")
	password := c.PostForm("password")

	if mobile == "" || password == "" {
		logging.Info("参数错误")
		appG.ResponseErr(e.ERROR, "参数错误")
		return
	}

	data := make(map[string]interface{})

	user := models.GetUser(mobile, password)

	if user.UserId <= 0 {
		logging.Info("参数错误, mobile:" + mobile + "|password:" + password)
		appG.ResponseErrMsg("用户名或密码错误")
		return
	} else {
		token, err := util.GenerateToken(user.UserId, user.Mobile)

		logging.Info(token)
		// token生成失败
		if err != nil {
			logging.Error("token生成失败:" + fmt.Sprintf("%s", err))
			appG.ResponseErrMsg(e.GetMsgLabel(e.ERROR_AUTH_TOKEN))
			return
		} else {
			data["token"] = token
		}
	}
	appG.ResponseSuccess("ok", data)
}

// 用户注册
func Register(c *gin.Context) {
	appG := app.Gin{C: c}

	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	name := c.PostForm("name")

	if mobile == "" || password == "" {
		logging.Info("参数错误")
		appG.ResponseErr(e.ERROR, "参数错误")
		return
	}

	isMobileExist, _ := models.IsMobileExist(mobile)
	if isMobileExist {
		logging.Info("手机号已存在:" + mobile)
		appG.ResponseErr(e.ERROR, "手机号已存在")
		return
	}

	err := models.AddUser(mobile, password, name)
	if err != nil {
		appG.ResponseErrMsg("注册失败:" + fmt.Sprintf("%s", err))
		return
	}
	appG.ResponseSuccess("ok", nil)
}

func GetUserInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	userId := c.GetInt("userId")
	user := models.GetUserById(userId)
	data := make(map[string]interface{})

	data["mobile"] = user.Mobile
	data["name"] = user.Name
	data["userId"] = user.UserId
	appG.ResponseSuccess("ok", user)
}
