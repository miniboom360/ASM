package handler

import (
	"backend/app"
	"backend/app/module"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(g *gin.Context) {

	req := app.LoginReq{}
	g.BindJSON(&req)
	if req.Username == "" || req.Password == "" {
		g.String(http.StatusOK, "用户名或者密码不能为空")
		return
	}
	user, _ := handlerLoginLogic(req)
	if user == nil {
		// 返回账号或者密码错误
		g.String(http.StatusBadRequest, "{err_code:'账号或者密码为空'}")
		return
	}

	lr := new(app.LoginResp)
	lr.Type = "success"
	lr.Message = "ok"
	lr.Code = 0
	lr.Result = user.Result
	// lr.Result.Token = user.Token
	// lr.Result.Roles

	g.JSON(http.StatusOK, lr)
}

func handlerLoginLogic(req app.LoginReq) (*app.User, error) {
	users, err := module.GetUsersInfo()
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]*app.User, 0)

	for _, user := range users {
		userMap[user.Username] = user
	}

	if user, ok := userMap[req.Username]; ok {
		if user.Password == req.Password {
			// 登录成功，将user的token内容返回给前端
			return user, nil
		} else {
			// 密码错误
			return nil, err
		}
	} else {
		// 用户名不存在
		return nil, err
	}

	return nil, err
}

func GetUserInfo(g *gin.Context) {
	// 获取这里的值，判断是不是正确的token，Authorization
	// g.Header("Authorization")
	token := g.GetHeader("Authorization")
	if token == "" {
		g.String(http.StatusOK, "传入参数错误")
		return
	}
	user, _ := handlerUserInfoLogic(token)
	if user == nil {
		// 返回账号或者密码错误
		g.String(http.StatusBadRequest, "{err_code:'账号或者密码为空'}")
		return
	}

	guir := new(app.GetUserInfoResp)
	guir.Result = user
	guir.Code = 0
	guir.Message = "ok"
	guir.Type = "success"
	g.JSON(http.StatusOK, guir)
}

func handlerUserInfoLogic(token string) (*app.User, error) {
	users, err := module.GetUsersInfo()
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]*app.User, 0)

	for _, user := range users {
		userMap[user.Token] = user
	}

	if user, ok := userMap[token]; ok {
		return user, nil
	} else {
		// 用户名不存在
		return nil, err
	}

	return nil, err
}
