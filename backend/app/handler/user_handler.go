package handler

import (
	"backend/app"
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

	// workflowID := req.ScanName + "-" + fmt.Sprintf("%d", time.Now().Unix())
	// go handlerScanLogic(workflowID, req)
	// g.String(http.StatusOK, fmt.Sprintf(`{"taskid":"%s"}`, workflowID))
}

func handlerLoginLogic(req app.LoginReq) {

}
