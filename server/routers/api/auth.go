package api

import (
	"gitee.com/favefan/doconsole/models"
	"net/http"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"

	"gitee.com/favefan/doconsole/pkg/app"
	"gitee.com/favefan/doconsole/pkg/e"
	"gitee.com/favefan/doconsole/pkg/util"
)

type loginForm struct {
	Username string `json:"username" valid:"Required; MaxSize(100)"`
	Password string `json:"password" valid:"Required; MaxSize(100)"`
}

// @Summary Get User
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /loginForm [get]
func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	// 表单数据结构体绑定与验证
	var a loginForm
	valid := validation.Validation{}
	err := c.BindJSON(&a)
	ok, _ := valid.Valid(&a)

	// 登陆表单数据结构体绑定错误处理 & 登陆表单数据验证错误处理 (！需要分离)
	if err != nil || !ok {
		// app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	isAuthenticated, err := models.CheckAuth(a.Username, a.Password)
	// 登录凭证错误的错误处理
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAuthCheckTokenFail, nil)
		return
	}
	if !isAuthenticated {
		appG.Response(http.StatusUnauthorized, e.ErrorAuth, nil)
		return
	}

	tokenString, err := util.GenerateToken(a.Username, a.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAuth, nil)
		return
	}
	appG.C.Header("Access-Token", tokenString)

	appG.Response(http.StatusOK, e.Success,
		map[string]string{
			"id":            "1",
			"name":          "1",
			"username":      "admin",
			"password":      "",
			"avatar":        "https://gw.alipayobjects.com/zos/rmsportal/jZUIxmJycoymBprLOUbT.png",
			"status":        "1",
			"telephone":     "",
			"lastLoginIp":   "27.154.74.117",
			"lastLoginTime": "1534837621348",
			"creatorId":     "admin",
			"createTime":    "1497160610259",
			"deleted":       "0",
			"roleId":        "admin",
			"lang":          "zh-CN",
			"token":         tokenString,
		},
	)
}

func Logout(c *gin.Context) {

}
